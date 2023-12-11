from flask import Flask, request, jsonify

app = Flask(__name__)

# Define a route for the AI request
@app.route('/ai', methods=['POST'])
def handle_ai_request():
    # Get JSON data from the request
    request_data = request.get_json()
    print("Received request data: ", request_data)

    # Perform some AI-related processing here...
    # For example, let's just echo back the received JSON data
    response_data = {
        "status": "success",
        "message": "Request processed",
        "data": request_data  # Echoing back the data for demonstration purposes
    }

    # Send JSON response
    return jsonify(response_data)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)  # Running on http://0.0.0.0:5000/
