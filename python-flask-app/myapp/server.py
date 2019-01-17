import time

from flask import Flask, Response, jsonify, request
from prometheus_client import Counter, Histogram, generate_latest, CONTENT_TYPE_LATEST

LATENCY_HISTOGRAM = Histogram('myapp_request_latency_seconds', 'myapp request latency', ['method', 'endpoint'])
REQUEST_COUNTER = Counter('myapp_request_count', 'myapp HTTP request count', ['method', 'endpoint', 'status'])

app = Flask(__name__)

def before_request():
    request.start_time = time.time()

def after_request(response):
    request_latency = time.time() - request.start_time
    LATENCY_HISTOGRAM.labels(request.method, request.path).observe(request_latency)
    REQUEST_COUNTER.labels(request.method, request.path, response.status_code).inc()
    return response

@app.route("/metrics")
def metrics():
    return generate_latest(), 200

@app.route("/<path:path>")
def other_routes(path):
    return "Hello world", 200

if __name__ == '__main__':
    app.before_request(before_request)
    app.after_request(after_request)
    app.run(host='0.0.0.0')