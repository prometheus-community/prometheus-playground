FROM python:3.7.0-alpine3.8

WORKDIR /usr/src/myapp

COPY server.py requirements.txt ./

RUN pip install -r requirements.txt

EXPOSE 5000

CMD ["python", "server.py"]