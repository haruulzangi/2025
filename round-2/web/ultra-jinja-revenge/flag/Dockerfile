FROM python:3.12-slim

WORKDIR /app

COPY --from=ghcr.io/astral-sh/uv:0.7.3 /uv /uvx /bin/
RUN uv venv
ENV PATH="/app/.venv/bin:$PATH"
RUN uv pip install flask gunicorn

RUN touch __init__.py && cat <<EOF > main.py
#!/srv/app/.venv/bin/python
import os
import base64
import logging
import secrets
from flask import Flask, request, render_template_string

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)
FLAG = os.getenv("FLAG", "The flag was not set. Please create a ticket in Discord.")

@app.route("/")
def index():
	response = FLAG.replace('SIGNATURE', secrets.token_urlsafe(16))
	logger.info(f"Generated flag: {response}")
	return response

if __name__ == "__main__":
	app.run(host="0.0.0.0", port=5000)
EOF

RUN useradd -m app && chown -R app:app /app

EXPOSE 1337

USER app

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1
CMD ["gunicorn", "--bind=0.0.0.0:1337", "main:app"]
