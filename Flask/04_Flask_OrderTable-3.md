# Flask에서 MariaDB와 Kafka를 연동-3

> MariaDB



```python
from flask import Flask, jsonify, request
from datetime import datetime
from flask_restful import reqparse
import flask_restful
import mariadb
import json
import uuid


app = Flask(__name__)
api = flask_restful.Api(app)

config = {
    'host': '127.0.0.1',
    'port': 13306,
    'user': 'root',
    'database': 'mydb'
}

@app.route('/')
def index():
    return "Welcome to Delivery Microservice!"


class Delivery(flask_restful.Resource):
    def __init__(self):
        self.conn = mariadb.connect(**config)
        self.cursor = self.conn.cursor()
    
    def get(self):
        sql = '''SELECT delivery_id, order_json, status, created_at
                FROM delivery_status order by id desc
        '''
        self.cursor.execute(sql)
        
        result_set = self.cursor.fetchall()

        row_headers = [x[0] for x in self.cursor.description]

        json_data = []
        for result in result_set:
            json_data.append(dict(zip(row_headers, result)))

        return jsonify(json_data)


class DeliveryStatus(flask_restful.Resource):
    def __init__(self):
        self.conn = mariadb.connect(**config)
        self.cursor = self.conn.cursor()

    def put(self, delivery_id):
        json_data = request.get_json()
        status = json_data['status']

        # DB INSERT
        sql = '''UPDATE delivery_status SET status=? WHERE delivery_id=?'''
        self.cursor.execute(sql, [status, delivery_id])
        self.conn.commit()

        json_data['updated_at'] = str(datetime.today())

        response = jsonify(json_data)
        response.status_code = 200

        return response

api.add_resource(Delivery, '/delivery-ms/deliveries')
api.add_resource(DeliveryStatus, '/delivery-ms/deliveries/<string:delivery_id>')

if __name__ == '__main__':
    app.run()
```

