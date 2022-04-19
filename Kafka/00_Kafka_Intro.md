# Kafka 란

Apache Kafka는 실시간으로 기록 스트림을 게시, 구독, 저장 및 처리할 수 있는 분산 데이터 스트리밍 플랫폼입니다. 이는 여러 소스에서 데이터 스트림을 처리하고 여러 사용자에게 전달하도록 설계되었습니다. 간단히 말해, A지점에서 B지점까지 이동하는 것뿐만 아니라 A지점에서 Z지점을 비롯해 필요한 모든 곳에서 대규모 데이터를 동시에 이동할 수 있습니다.

Apache Kafka는 전통적인 엔터프라이즈 메시징 시스템의 대안입니다. 하루에 1조4천억 건의 메시지를 처리하기 위해 LinkedIn이 개발한 내부 시스템으로 시작했으나, 현재 이는 다양한 기업의 요구사항을 지원하는 애플리케이션을 갖춘 오픈소스 데이터 스트리밍 솔루션이 되었습니다.

<br>
<br>

# Kafka 설치

[Kafka](https://kafka.apache.org/downloads.html)에서 최신에 나온 버전을 다운로드를 받습니다. 저는 `3.1.0` 버전을 다운받았습니다. 다운 받을때의 OS는 WindowOS, MacOS의 차이는 없습니다.
```
$ wget http://apache.mirror.cdnetworks.com/kafka/3.1.0/kafka_2.13-3.1.0.tgz

$ tar -zxvf kafka_2.13-3.1.0.tgz

$ mkdir zookeeper & mv kafka_2.13-3.1.0 zookeeper
```
<br>

java 설치
```
$ yum install -y java-11-openjdk.x86_64
```

<br>
<br>

# Kafka 실행

### STEP 1: START THE KAFKA ENVIRONMENT
백그라운드로 실행하려면 `-daemon` 옵션을 추가하여 기동

```
$ bin/zookeeper-server-start.sh -daemon config/zookeeper.properites
```

```
$ bin/kafka-server-start.sh config/server.properties
```

<br>


### STEP 2: CREATE A TOPIC TO STORE YOUR EVENTS
> `Kafka default port -- 9092`, `zookeeper default port -- 2181`

```
$ bin/kafka-topics.sh --create --topic <TOPIC_NAME> /
--bootstrap-server <IP>:9092 /
--partitions 3 /
--replication-factor 3

Created topic <TOPIC_NAME>
```
`--partitions` 와 `--replication-factor`를 설정하지 않으면 클러스터 없이 단일로 사용합니다.

`bootstrap-server`는 활성 Kafka 브로커 중 하나의 주소를 가리킵니다. 모든 브로커는 Zookeeper를 통해 서로에 대해 알고 있으므로 어느 브로커를 선택하든 상관 없습니다.

- `--partitions` : 파티션을 사용하면 데이터를 분할할 브로커 수를 결정할 수 있습니다. 일반적으로 브로커 수와 동일하게 설정합니다. 3개의 브로커를 설정했으므로 이 옵션을 3으로 설정합니다. 
- `--replication-factor` :  하는 데이터 복사본의 수를 나타냅니다(브로커 중 하나가 다운되는 경우에도 다른 브로커에 데이터가 남아 있음). 이 값을 3로 설정했으므로 데이터는 브로커에 복사본 두 개를 더 갖습니다.

<br>

#### Topic list
```
$ bin/kafka-topics.sh --list --bootstrap-server <IP>:9092
```

- 해당 kafka의 상세정보
```
$ bin/kafka-topics.sh --describe --topic <TOPIC_NAME> --bootstrap-server <IP>:9092

Topic: my-kafka-topic   TopicId: 2c7cvTC1QGKy2bu18revoA PartitionCount: 3       ReplicationFactor: 3    Configs: segment.bytes=1073741824
        Topic: my-kafka-topic   Partition: 0    Leader: 3       Replicas: 3,1,2 Isr: 3,1,2
        Topic: my-kafka-topic   Partition: 1    Leader: 1       Replicas: 1,2,3 Isr: 1,2,3
        Topic: my-kafka-topic   Partition: 2    Leader: 2       Replicas: 2,3,1 Isr: 2,3,1
```

주제의 파티션 및 복제본에 대한 세부 정보를 출력합니다. Partition, Leader/follower, Replicas, Isr(In Sync Replica) 정보를 보여줍니다. 

여기서 `Isr`은 kafka 리더 파티션과 팔로워 파티션이 모두 싱크가 된 상태를 나타냅니다. 만일, 브로커 중 1대의 서버가 중지된 상태라면 `Isr` 은 2개 만 표시됩니다. 3번 브로커 서버가 중지되었다면 `Leader는` 1 또는 2가 되고 `Isr` 은 1,2 가 됩니다.

<br>


### STEP 3: WRITE AND READ SOME EVENTS INTO THE TOPIC

Consumer를 실행시키고 새로운 터미널 창을열어서 Producer를 실행시킵니다.

Producer 창에서 원하는 문장을 입력하면 Consumer창에서 문장이 출력이 되는것을 볼 수 있습니다.

**Windows**

- `> bin\windows\kafka-console-consumer.bat --topic <TOPIC_NAME> --from-beginning --bootstrap-server <IP>:9092`
  - `--from-beginning` -- 새로운 consumer 창을 실행시키면 기존의 데이터를 다 가져와서 출력시킨다.



- `> bin\windows\kafka-console-producer.bat --topic <TOPIC_NAME> --bootstrap-server <IP>:9092`


<br>


### STEP 4: DELETE TOPIC

```cmd
$ bin/kafka-topics.sh --delete --topic <TOPIC_NAME> --bootstrap-server <IP>:9092
```

Windows에서 `topic`을 삭제하면 삭제가 되고나서 kafka server가 강제 종료가 되어 버립니다.  kafka server를 다시 실행해도 shutdown이 되기 때문에 `C:\tmp`에서 `kafka-logs`파일을 삭제하고 zookeeper server도 종료 한뒤 `zookeeper`파일도 삭제 해 주셔야 합니다.

이러한 일이 발생하는 이유는 Kafka자체가 Linux기반이기 때문에 Window에서 작은 오류가 발생하는 것 같습니다. 이러한 문제를 해결 하기 위해서는 `topic`을 삭제하기 보다는 새로 만드는 것을 추천 드립니다.





