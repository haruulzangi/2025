FROM python:3.9-slim AS builder

WORKDIR /app

RUN python3 -m venv /app/venv && chmod +x /app/venv/bin/activate
ENV PATH="/app/venv/bin:$PATH"
COPY requirements.txt .
RUN pip install --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt

COPY app ./app

RUN chown -R 65532:65532 /app/venv

FROM python:3.9-slim

USER 65532:65532

WORKDIR /app

COPY --from=builder --chmod=755 /app/venv /app/venv
COPY --from=builder /app /app

ENV PATH="/app/venv/bin:$PATH" \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

EXPOSE 8080
CMD ["gunicorn", "--bind=0.0.0.0:8080", "app.main:app"]
