FROM python:3.10-slim
RUN addgroup --system ctf && adduser --system --ingroup ctf ctf
WORKDIR /app
RUN pip install --no-cache-dir flask
COPY --chown=ctf:ctf app.py /app/
COPY --chown=ctf:ctf templates/ /app/templates/
COPY --chown=ctf:ctf static /app/static/
COPY flag.txt /
RUN chmod 444 /flag.txt
EXPOSE 5000
USER ctf
CMD ["python", "app.py"]
