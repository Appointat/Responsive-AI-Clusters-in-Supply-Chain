from flask import Flask, request, jsonify

from multi_agent_communication_supply_chain import role_playing

import logging, queue

global messages_queue, central_hub_json

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

import asyncio
import websockets
import json
import logging
import functools
import threading

# async def send_streaming_message(websocket, path, message):
#     try:
#         while True:
#             msg = message
#             await websocket.send(json.dumps({
#                 "SpeakerID": "0",
#                 "ReceiverID": "1",
#                 "text": msg
#             }))
#             await asyncio.sleep(1)

#     except websockets.exceptions.ConnectionClosedError as e:
#         logging.error(f"WebSocket connection closed: {e}")

async def get_message_from_queue(q):
    loop = asyncio.get_running_loop()
    import concurrent.futures
    with concurrent.futures.ThreadPoolExecutor() as pool:
        # Run a blocking function in the background
        message = await loop.run_in_executor(pool, q.get)
        return message

async def send_streaming_message(websocket, path, messages_queue):
    print(f"messages_queue, send_streaming_message:\n{messages_queue}")

    while True:
        message = await get_message_from_queue(messages_queue)  # Retrieve a message from the queue
        print(f"message:\n{message}")
        if message is None:  # You can use a sentinel value like None to stop the loop
            break

        sender_id = message["sender_id"]
        user_message = message["user_message"]
        for char in user_message:  # user
            msg_to_send = {
                "SpeakerID": sender_id,
                "ReceiverID": "0",
                "text": char,
            }
            await asyncio.sleep(0.05)
            await websocket.send(json.dumps(msg_to_send))

        assistant_message = message["assistant_message"]
        for char in assistant_message:  # assistant
            msg_to_send = {
                "SpeakerID": "0",
                "ReceiverID": sender_id,
                "text": char,
            }
            await asyncio.sleep(0.05)
            await websocket.send(json.dumps(msg_to_send))

        messages_queue.task_done()  # Mark the task as done
 
def run_websocket_server(messages_queue):
    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)
    start_server = websockets.serve(functools.partial(send_streaming_message, messages_queue=messages_queue), 'localhost', 8000)
    loop.run_until_complete(start_server)
    loop.run_forever()

app = Flask(__name__)
websocket_server_thread = None
current_messages_queue = None

# Define a route for the AI request
@app.route('/ai', methods=['POST'])
def handle_ai_request():
    global websocket_server_thread
    global current_messages_queue

    # Get JSON data from the request
    request_data = request.get_json()

    # Perform some AI-related processing here...
    global central_hub_json
    response_json, _central_hub_json, new_messages_queue = role_playing(request_json=request_data, central_hub_json=central_hub_json)
    central_hub_json = _central_hub_json
    print(f"messages_queue, handle_ai_request:\n{new_messages_queue}")

    # Start WebSocket server in a separate thread if not already running
    current_messages_queue = new_messages_queue
    if websocket_server_thread is None or not websocket_server_thread.is_alive():
        websocket_server_thread = threading.Thread(target=run_websocket_server, args=(current_messages_queue,), daemon=True)
        websocket_server_thread.start()

    # Send JSON response
    return jsonify(response_json)


if __name__ == '__main__':
    # Configure logging
    logging.basicConfig(level=logging.INFO)

    # Running on http://0.0.0.0:5000/ without threading even in debug mode
    app.run(debug=True, host='0.0.0.0', port=5000, threaded=False)
