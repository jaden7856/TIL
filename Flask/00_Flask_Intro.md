# Flask 란?

**Micro Web Framework** 간단한 웹 사이트, 혹은 간단한 API 서버를 만드는 데에 특화 되어있는 **Python Web Framework** 입니다. 요즘에는 클라우드 컴퓨팅의 발달로 **Docker**, **Kubernetes** 와 접목해서 소규모 컨테이너 단위로 기능 별 개발 후, 한 꺼번에 배포하는 방식, 혹은 배포 후 기능 추가 하는 식으로 자주 사용하고 있습니다.



### 장점

- **가볍게 배울 수 있다!** (Python, HTML + CSS + Javascript만 할 줄 알면 금방 배운다!)
- **가볍게 사용 할 수 있다!** (코드 몇 줄이면 금방 만든다!)
- **가볍게 배포 할 수 있다!** (virtualenv에 Flask 깔고 바로 배포 하면 됨!)



### 단점

- **Django** 에 비해서 자유도는 높으나, 제공해 주는 기능이 덜 하다.
- 복잡한 어플리케이션을 만들려고 할 때 **해야 할 것들이 많다.**





# Flask 설치

```cmd
> conda create -n <virtual_name> python=3.8 flask
> conda activate <virtual_name>
> conda list
```





# Flask 실행

**`app.py`**

```python
from flask import Flask, jsonify, request
from datetime import datetime
import uuid


app = Flask(__name__)

@app.route('/')
def index():
    return "Hello, World!"


@app.route('/health-check')
def health_check():
    return "Server is running on 5000 port"


@app.route('/users')
def users():
    return "** Users List"

@app.route('/users/<userId>')
def users_detail(userId):
    # return "{\"name\":%s}" % (userId)
    return jsonify({"user_id": userId})


@app.route('/users', methods = ['POST'])
def userAdd():
    user = request.get_json()
    user['user_id'] = uuid.uuid4()  # uuid1() ~ uuid5()
    user['created_at'] = datetime.today()
    
    # 200 OK -> 201 Created
    return jsonify(user), 201


@app.route('/users/<userId>/orders')
def order_list(userId):
    return "** Orders List"


@app.route('/users/<userId>/orders/<orderId>')
def order_detail(userId, orderId):
    return jsonify({"order_id" : orderId})


@app.route('/users/<userId>/orders', methods = ['POST'])
def orderAdd(userId):
    order = request.get_json()
    order['order_id'] = uuid.uuid4()
    
    return jsonify(order), 201


if __name__ == "__main__":
    app.run()
```

```cmd
> flask run --port <port_number>
```



#### jsonify

`def users_detail`에서 `return "{\"name\":%s}" % (userId)` 이 코드를 사용하여 `{"name":userId}`가 출력이 되지만, `jsonify`를 사용하여 `return jsonify({"user_id": userId})` 코드 처럼 쉽게 json 형태로 출력 할 수 있습니다.



#### UUID

UUID는 128비트 숫자이며, 32자리의 16진수로 표현된다. 여기에 8-4-4-4-12 글자마다 하이픈을 집어 넣어 5개의 그룸으로 구분한다. ex) `550e8400-e29b-41d4-a716-446655440000`

- `uuid1` -- 타임스탬프를 기준으로 생성
- `uuid3` -- MD5 해쉬를 이용해 생성

- `uuid4` -- random 생성
- `uuid5` -- SHA-1 해쉬를 이용해 생성





## Flask 실행

```cmd
> flask run
```

`set FLASK_APP=<FILE_NAME>.py` -- `flask run`을 할 때 파일명이 `app.py`가 아니라면 설정을 해 주자