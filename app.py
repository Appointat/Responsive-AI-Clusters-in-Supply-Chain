from flask import Flask, jsonify

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

@app.route('/events/<int:event_id>', methods=['GET'])
def get_event(event_id):
    event = next((event for event in data['events'] if event['id'] == event_id), None)
    if event is not None:
        return jsonify(event)
    else:
        return jsonify({"message": "Event not found"}), 404

if __name__ == '__main__':
    app.run(debug=True)
