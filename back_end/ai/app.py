from flask import Flask, request, jsonify

from multi_agent_communication_supply_chain import role_playing

app = Flask(__name__)

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

# Define a route for the AI request
@app.route('/ai', methods=['POST'])
def handle_ai_request():
    # Define global variables
    global central_hub_json

    # Get JSON data from the request
    request_data = request.get_json()

    # Perform some AI-related processing here...
    response_data, _central_hub_json = role_playing(request_json=request_data, central_hub_json=central_hub_json)
    central_hub_json = _central_hub_json

    # Send JSON response
    return jsonify(response_data)

if __name__ == '__main__':
    # Running on http://0.0.0.0:5000/ without threading even in debug mode
    app.run(debug=True, host='0.0.0.0', port=5000, threaded=False)
