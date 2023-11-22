from flask import Flask, jsonify, request

app = Flask(__name__)

data = {
    "events": [
        {"id": 1, "name": "New Year's Day", "date": "2023-01-01"},
        {"id": 2, "name": "Valentine's Day", "date": "2023-02-14"},
        {"id": 3, "name": "St. Patrick's Day", "date": "2023-03-17"}
    ]
}

@app.route('/events', methods=['GET'])
def get_events():
    return jsonify(data)

@app.route('/events', methods=['POST'])
def create_event():
    event_data = request.json
    if event_data is None:
        return jsonify({"message": "No input data provided"}), 400
    data['events'].append(event_data)
    return jsonify(event_data), 201

if __name__ == '__main__':
    app.run(debug=True)
