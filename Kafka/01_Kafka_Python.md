# Python에서 Kafka 사용

python에서 kafka를 쓰기위해 다운로드하는데, 저는 conda를 사용하여 가상환경에 다운로드 하겠습니다.

```cmd
> conda activate base
> pip install kafka-python
```



- 로컬에 원하는 폴더에서 `kafka_consumer.py`파일을 하나 만들겠습니다.

```python
from kafka import KafkaConsumer
from json import loads
import time

consumer = KafkaConsumer('<TOPIC_NAME>', 
            bootstrap_servers=['127.0.0.1:9092'],
            auto_offset_reset='earlisest',
            enable_auto_commit=True,
            group_id='my-group',
            # value_deserializer=lambda x: loads(x.decode('utf-8')),
            consumer_timeout_ms=1000)

start = time.time()

for message in consumer:
    topic = message.topic
    partition = message.partition
    offset = message.offset
    key = message.key
    value = message.value
    print("Topic:{}, Partition:{}, Offset:{}, Key:{}, Value:{}".format(
        topic, partition, offset, key, value))

print("Elapsed: ", (time.time() - start))
```



이제 코드를 실행시키면 Python에서도 Kafka명령어와 같은 출력값이 나오게 됩니다. Producer에서 어떤 문장을 입력했는지에 따라 출력값이 다 다르게 나오겠죠.

```cmd
> python kafka_consumer.py
Topic:quickstart-events, Partition:0, Offset:0, Key:None, Value:b''
Topic:quickstart-events, Partition:0, Offset:1, Key:None, Value:b'Hello, World!'
Elapsed:  1.187957525253296
```

```cmd
> bin\windows\kafka-console-consumer.bat --topic <TOPIC_NAME> --from-beginning --bootstrap-server localhost:9092

Hello, World!
```



Producer도 Python으로 작업 해 보도록 하겠습니다. `kafka_producer.py`파일을 생성합니다.

kafka data는 원래 json 타입으로 저장이 되어야 하기때문에 dict타입으로 저장 하도록 하겠습니다.

```python
from kafka import KafkaProducer
from json import dumps
import time


# dict (key, value) -> object
# str -> string
producer = KafkaProducer(acks=0, 
            compression_type='gzip',
            bootstrap_servers=['127.0.0.1:9092'],
            value_serializer=lambda x : dumps(x).encode('utf-8'))

start = time.time()
for i in range(10):
    data = {'name': 'Dowon-' + str(i)}
    producer.send('<TOPIC_NAME>', value=data)
    producer.flush()

print("Doen. Elapsed time: ", (time.time() - start))
```



기존의 topic은 직렬형태의 data가 이미 저장이 되어 있어 새로운 topic을 생성하겠습니다.

```cmd
> bin\windows\kafka-topics.bat --create --topic <TOPIC_NAME> --bootstrap-server localhost:9092
```



파일 실행 

```cmd
> python kafka_producer.py
```



확인 해보시면 두군데 다 같은 출력값이 나와야 정상입니다.

```cmd
> python kafka_consumer.py
> bin\windows\kafka-console-consumer.bat --topic <TOPIC_NAME> --from-beginning --bootstrap-server localhost:9092
```



