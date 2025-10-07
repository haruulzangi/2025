FROM python:3.12-slim

WORKDIR /app

COPY --from=ghcr.io/astral-sh/uv:0.7.3 /uv /uvx /bin/
RUN uv venv
ENV PATH="/app/.venv/bin:$PATH"

RUN uv pip install psycopg2-binary

COPY main.py ./

RUN useradd -m app && chown -R app:app /app

EXPOSE 5000

USER app

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

CMD ["python3", "main.py"]
