from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/ai', methods=['POST'])
def handle_ai_request():
    data = request.json  # 接收 JSON 请求数据
    print("Received AI Request:", data)  # 打印收到的数据（可选）

    # 将接收到的数据原样返回
    return jsonify(data)

if __name__ == '__main__':
    app.run(port=5000)

