from flask import Flask, render_template_string
from flask import request

app = Flask(__name__)


def is_good(input):
    allowlist = ["{%", "{{", "%}", "}}"]
    return all(token in input for token in allowlist)


@app.route("/", methods=["GET", "POST"])
def hello():
    if request.method == "POST":
        username = request.form.get("username", "")
        if not is_good(username):
            return "Not good enough!", 400
        return render_template_string(f"Hello, {username}!")
    return "Hello! Let me know your username via post, shall ya?"


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, threaded=False)
