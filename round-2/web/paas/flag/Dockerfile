FROM alpine AS builder

WORKDIR /app

RUN cat <<EOF > main.py
import os
import logging
import secrets
import socketserver

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

FLAG = os.getenv("FLAG", "The flag was not set. Please create a ticket in Discord.")

def generate_flag():
	return FLAG.replace('SIGNATURE', secrets.token_urlsafe(16))

class ThreadedTCPRequestHandler(socketserver.BaseRequestHandler):
	def handle(self):
		flag = generate_flag()
		logger.info(f"Generated flag: {flag}")
		self.request.sendall(flag.encode())
		self.request.close()

if __name__ == "__main__":
	socketserver.ThreadingTCPServer.allow_reuse_address = True
	server = socketserver.ThreadingTCPServer(("0.0.0.0", 1337), ThreadedTCPRequestHandler)
	server.serve_forever()
EOF

FROM gcr.io/distroless/python3-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/main.py .

EXPOSE 1337
CMD ["main.py"]
