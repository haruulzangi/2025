from flask import Flask, request, jsonify, render_template
from urllib.parse import urlparse, unquote
import http.client, base64

BLACKLIST = ['.', '..', '@', '\\', '//', 'flag', '%']

app = Flask(__name__)

@app.route('/', methods=['GET'])
def home():
    return render_template('index.html')

@app.route('/public', methods=['POST'])
def fetch_public():
    gerege_image = request.form.get('p')
    if gerege_image is None or gerege_image == '' or not isinstance(gerege_image, str):
        return jsonify({'message': 'Invalid Gerege fetch', 'data': None}), 400
    if any(_ in gerege_image for _ in BLACKLIST):
        return jsonify({'message': 'Invalid Gerege fetch', 'data': None}), 400
    try:
        gerege_image = unquote(gerege_image)
        if any(_ in gerege_image for _ in BLACKLIST):
            return jsonify({'message': 'Invalid Gerege fetch', 'data': None}), 400
        safe_url = urlparse(f"http://nginx:80{gerege_image}.png")
    except Exception as e:
        return jsonify({'message': str(e), 'data': None}), 400
    if not safe_url.hostname:
        return jsonify({'message': 'Malformed fetch URL', 'data': None}), 400
    try:
        conn = http.client.HTTPConnection(safe_url.hostname, safe_url.port, timeout=5)
        conn.request('GET', safe_url.path)
        r = conn.getresponse()
        if r.status == 200:
            data = base64.b64encode(r.read()).decode('ascii')
            return jsonify({'message': f'Gerege received', 'data': data}), 200
        else:
            return jsonify({'message': 'Failed to fetch Gerege', 'data': None}), 500
    except http.client.HTTPException as e:
        return jsonify({'message': str(e), 'data': None}), 500
    except Exception as e:
        return jsonify({'message': str(e), 'data': None}), 500

@app.route('/gerege', methods=['GET'])
def gerege_query():
    t = request.args.get('type')
    if t is None:
        return jsonify({'message': 'Gerege type is required', 'data': None}), 400
    return gerege(t)

@app.route('/gerege/<string:type>', methods=['GET'])
def gerege(type):
    if not type or type == '' or not isinstance(type, str):
        return jsonify({'message': 'Invalid Gerege type', 'data': None}), 400
    if type == 'gold':
        if request.remote_addr != '127.0.0.1' and request.remote_addr != '::1':
            return jsonify({'message': 'Only authorized officials can get the GOLDEN Gerege', 'data': None}), 403
        conn = http.client.HTTPConnection('nginx', 80, timeout=5)
        conn.request('GET', '/flag.txt')
        r = conn.getresponse()
        if r.status == 200:
            data = r.read().decode()
            return jsonify({'message': 'Gold Gerege received', 'data': data}), 200
        return jsonify({'message': 'Failed to fetch flag', 'data': None}), 500
    elif type == 'silver':
        data = "Under the power of the Eternal Heaven, may the name of the Khan be blessed. Whoever does not show reverence shall be punished by death"
        return jsonify({'message': 'Silver Gerege received', 'data': data}), 200
    elif type == 'bronze':
        data = "Under the power of the Eternal heaven, Under the patronage of the Great Wisdom"
        return jsonify({'message': 'Bronze Gerege received', 'data': data}), 200
    else:
        return jsonify({'message': 'Unknown Gerege type', 'data': None}), 400

if __name__ == '__main__':
    app.run('::', port=5000)
