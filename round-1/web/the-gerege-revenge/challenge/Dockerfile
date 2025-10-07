FROM python:3.13.1-alpine

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY templates/index.html templates/index.html
COPY app.py .

EXPOSE 5000
CMD ["python", "app.py"]
