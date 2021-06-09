#  Micro Service

### Consul 설치

```cmd
> pip install python-consul
```



저희가 실습할 데이터를 가져오겠습니다.

```cmd 
> git clone https://github.com/joneconsulting/flask_msa.git
```



```cmd
consul_demo> docker-compose up
```



dig 를 설치한 경로로 커맨드창을 열고

```cmd
BIND9.16.13> curl http://localhost:8500/v1/catalog/service/order-ms
```

데이터가 정상적으로 나오는지 확인



`service_list.py` 파일을 실행하여 서비스가 정상적으로 작동되는지 확인하기위해 실행을 해야 하지만 서버 자체가 docker로 실행이 되기때문에 local환경에서 `python service_list.py`를 하면 오류가 발생합니다. 때문에 코드를 수정 하겠습니다.

```python
			:
        	:
service_address = client.catalog.service(serviceName)[1][0]['ServiceAddress']
service_port = client.catalog.service(serviceName)[1][0]['ServicePort']

print(service_address)
print(service_port)

# 나머지는 다 주석 처리
```

```cmd
consul_demo> python service_list.py
10.5.0.5
5000
```





# Store Key/Value

```cmd
consul_demo> python kvs_put.py
```

```cmd
consul_demo> python kvs_get.py
b'bar'
```

