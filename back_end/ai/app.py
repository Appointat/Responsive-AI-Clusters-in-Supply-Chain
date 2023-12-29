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
        "olive_oil": {
            "current_storage_amount": 1000,
        },
        "baguette": {
            "current_storage_amount": 1000,
        },
        "manchego_cheese": {
            "current_storage_amount": 1000,
        },
        "black_tea": {
            "current_storage_amount": 1000,
        },
    }
}

response_json = {
    "outlet_inventory": {
        "olive_oil": {
            "changed_replenishment_amount_from_central_hub": 1000,
        },
        "baguette": {
            "changed_replenishment_amount_from_central_hub": 1000,
        },
        "manchego_cheese": {
            "changed_replenishment_amount_from_central_hub": 1000,
        },
        "black_tea": {
            "changed_replenishment_amount_from_central_hub": 1000,
        }
    },
    "central_hub_inventory": {
        "olive_oil": {
            "current_storage_amount": 1000,
        },
        "baguette": {
            "current_storage_amount": 1000,
        },
        "manchego_cheese": {
            "current_storage_amount": 1000,
        },
        "black_tea": {
            "current_storage_amount": 1000,
        },
    },
    "transportation_duration": 1
}


app = Flask(__name__)

async def get_message_from_queue(messages_queue):
    return await asyncio.to_thread(messages_queue.get)

async def send_streaming_message(websocket, path):
    import copy
    queue_copy = copy.deepcopy(messages_queue.queue)  # Avoid modifying the original queue
    print(f"messages_queue, send_streaming_message:\n{list(queue_copy)}")

    while True:
        message = await get_message_from_queue(messages_queue)  # Retrieve a message from the queue
        print(f"The message from the message queue:\n{message}")
        if message is None:
            break

        sender_id = message["sender_id"]

        user_message = message["user_message"] + "\n\n"
        for char in user_message:  # user
            msg_to_send = {
                "SpeakerID": sender_id,
                "ReceiverID": "0",
                "text": char,
            }
            time.sleep(0.005)
            await websocket.send(json.dumps(msg_to_send))

        assistant_message = message["assistant_message"] + "\n\n"
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

    # Perform some AI-related processing with role_playing
    global central_hub_json
    response_json, updated_central_hub_json = role_playing(request_json=request_data, central_hub_json=central_hub_json)
    central_hub_json = updated_central_hub_json

    # Return the response from role_playing
    return jsonify(response_json)


@app.route('/send/<message>')
def send_message(message):
    messages_queue.put(message)
    return 'Message enqueued'

def run_flask_app():
    # Running on http://0.0.0.0:5000/ without threading even in debug mode
    app.run(debug=False, host='0.0.0.0', port=5000)

def main():
    flask_thread = threading.Thread(target=run_flask_app)
    flask_thread.start()

    run_websocket_server()

if __name__ == "__main__":
    main()
