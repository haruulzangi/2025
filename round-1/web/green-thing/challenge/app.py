import base64, pickle
from flask import Flask, request, render_template

app = Flask(__name__)

MAX_LEN = 2500
BANNED_SUBSTR = [
    "os", "system", "popen", "subprocess", "exec", "open",
    "__reduce__", "__reduce_ex__", "__getattribute__"
]

def naive_filter(raw: bytes) -> bool:
    if len(raw) > MAX_LEN:
        return False
    s = raw.decode("latin-1", errors="ignore")
    for bad in BANNED_SUBSTR:
        if bad in s:
            return False
    return True

@app.route("/", methods=["GET"])
def index():
    return render_template("index.html")

@app.route("/submit", methods=["POST"])
def submit():
    b64 = request.form.get("payload", "").strip()
    try:
        blob = base64.b64decode(b64, validate=True)
    except Exception:
        return render_template("result.html", error="Invalid Base64")

    if not naive_filter(blob):
        return render_template("result.html", error="Blocked by filter")

    try:
        obj = pickle.loads(blob)
    except Exception as e:
        return render_template("result.html", error=f"Error: {e}")
        #return render_template("result.html", result=f"Aldaa")

   
    if isinstance(obj, tuple) and len(obj) == 2 and callable(obj[0]) and isinstance(obj[1], tuple):
        try:
            obj = obj[0](*obj[1])
        except Exception as e:
            return render_template("result.html", result=f"Execution error: {e}")

    return render_template("result.html", result=obj)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=False)
