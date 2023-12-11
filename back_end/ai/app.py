from flask import Flask, request, jsonify
import json

from multi_agent_communication_supply_chain import role_playing

app = Flask(__name__)

# Define a route for the AI request
@app.route('/ai', methods=['POST'])
def handle_ai_request():
    # Get JSON data from the request
    request_data = request.get_json()

    # Perform some AI-related processing here...
    response_data = role_playing(request_json=request_data)

    # Send JSON response
    return jsonify(response_data)

if __name__ == '__main__':
    # Running on http://0.0.0.0:5000/ without threading even in debug mode
    app.run(debug=True, host='0.0.0.0', port=5000, threaded=False)
