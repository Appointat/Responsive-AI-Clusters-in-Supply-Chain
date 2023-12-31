from flask import Flask, request, jsonify
import asyncio
import functools
import json
import threading
import time
import websockets

from multi_agent_communication_supply_chain import role_playing, messages_queue


global central_hub_json

central_hub_json = {
    "central_hub_inventory": {
        "baguette": {
            "current_storage_amount": 1000,
        },
        "black_tea": {
            "current_storage_amount": 1000,
        },
        "manchego_cheese": {
            "current_storage_amount": 1000,
        },
        "olive_oil": {
            "current_storage_amount": 1000,
        },
    }
}

# Example of request JSON
# response_json = {
#     "outlet_inventory": {
#         "olive_oil": {
#             "changed_replenishment_amount_from_central_hub": 1000,
#         },
#         "baguette": {
#             "changed_replenishment_amount_from_central_hub": 1000,
#         },
#         "manchego_cheese": {
#             "changed_replenishment_amount_from_central_hub": 1000,
#         },
#         "black_tea": {
#             "changed_replenishment_amount_from_central_hub": 1000,
#         }
#     },
#     "central_hub_inventory": {
#         "olive_oil": {
#             "current_storage_amount": 1000,
#         },
#         "baguette": {
#             "current_storage_amount": 1000,
#         },
#         "manchego_cheese": {
#             "current_storage_amount": 1000,
#         },
#         "black_tea": {
#             "current_storage_amount": 1000,
#         },
#     },
#     "transportation_duration": 1
# }


app = Flask(__name__)

async def get_message_from_queue(messages_queue):
    return await asyncio.to_thread(messages_queue.get)

async def send_streaming_message(websocket, path):
    while True:
        message = await get_message_from_queue(messages_queue)  # Retrieve a message from the queue
        print(f"The message from the message queue:\n{message}")
        if message is None:
            break

        sender_id = message["sender_id"]
        user_message = message["user_message"] + "\n\n"
        assistant_message = message["assistant_message"] + "\n\n"

        for char in user_message:  # user
            msg_to_send = {
                "SpeakerID": sender_id,
                "ReceiverID": "0",
                "text": char,
            }
            time.sleep(0.005)
            await websocket.send(json.dumps(msg_to_send))

        for char in assistant_message:  # assistant
            msg_to_send = {
                "SpeakerID": "0",
                "ReceiverID": sender_id,
                "text": char,
            }
            time.sleep(0.005)
            await websocket.send(json.dumps(msg_to_send))

        messages_queue.task_done()  # Mark the task as done

def run_websocket_server():
    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)
    start_server = websockets.serve(functools.partial(send_streaming_message), 'localhost', 8000)
    loop.run_until_complete(start_server)
    loop.run_forever()

websocket_server_thread = None
current_messages_queue = None

# Define a route for the AI request
@app.route('/ai', methods=['POST'])
def handle_ai_request():
    # Get JSON data from the request
    request_data = request.get_json()

    def format_product_names(json_data):
        if "outlet_inventory" in json_data:
            formatted_inventory = {}
            for product_name, details in json_data["outlet_inventory"].items():
                # Convert the product name to lowercase and replace spaces with underscores
                formatted_name = product_name.lower().replace(" ", "_")
                formatted_inventory[formatted_name] = details

            json_data["outlet_inventory"] = formatted_inventory
        return json_data

    request_data = format_product_names(request_data)

    # Perform some AI-related processing with role_playing
    global central_hub_json
    try:
        response_json, updated_central_hub_json = role_playing(request_json=request_data, central_hub_json=central_hub_json)
    except:
        # If the role_playing function fails, return a default response
        response_json = {
            "outlet_inventory": {
                "baguette": {
                    "future_storage_amount": 50,
                    "specific_reason_of_replenishment": "to meet the moderate demand as per the client\"s preferences"
                },
                "black_tea": {
                    "future_storage_amount": 20,
                    "specific_reason_of_replenishment": "to maintain a minimal stock level due to the client\"s minimal interest"
                },
                "manchego_cheese": {
                    "future_storage_amount": 40,
                    "specific_reason_of_replenishment": "to meet the strong demand as per the client\"s preferences"
                },
                "olive_oil": {
                    "future_storage_amount": 30,
                    "specific_reason_of_replenishment": "to meet the strong demand as per the client\"s preferences"
                }
            },
            "central_hub_inventory": {
                "baguette": {
                    "current_storage_amount": 530
                },
                "black_tea": {
                    "current_storage_amount": 364
                },
                "manchego_cheese": {
                    "current_storage_amount": 530
                },
                "olive_oil": {
                    "current_storage_amount": 180
                }
            },
            "transportation_duration": 1
        }
        updated_central_hub_json = {
            "central_hub_inventory": {
                "baguette": {
                    "current_storage_amount": 530
                },
                "black_tea": {
                    "current_storage_amount": 364
                },
                "manchego_cheese": {
                    "current_storage_amount": 530
                },
                "olive_oil": {
                    "current_storage_amount": 180
                }
            }
        }

    for product in updated_central_hub_json["central_hub_inventory"]:
        product_account = int(updated_central_hub_json["central_hub_inventory"][product]["current_storage_amount"])
        if product_account <= 0:
            updated_central_hub_json["central_hub_inventory"][product]["current_storage_amount"] = 1000
            print(f"Product {product} is out of stock, replenished to 1000.")
    central_hub_json = updated_central_hub_json

    # Return the response from role_playing
    print(f"Response JSON:\n{response_json}")
    return jsonify(response_json)

def run_flask_app():
    # Running on http://0.0.0.0:5000/ without threading even in debug mode
    app.run(debug=False, host='0.0.0.0', port=5000)

def main():
    flask_thread = threading.Thread(target=run_flask_app)
    flask_thread.start()

    run_websocket_server()

if __name__ == "__main__":
    main()
