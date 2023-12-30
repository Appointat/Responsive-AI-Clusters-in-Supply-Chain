from flask import Flask, request, jsonify
import threading

app = Flask(__name__)

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
    print(f"Request data:\n{request_data}")

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

    # Return the response from role_playing
    return jsonify(response_json)

def run_flask_app():
    # Running on http://0.0.0.0:5000/ without threading even in debug mode
    app.run(debug=False, host='0.0.0.0', port=5000)

def main():
    flask_thread = threading.Thread(target=run_flask_app)
    flask_thread.start()

if __name__ == "__main__":
    main()
