# Kafka 란

Apache Kafka는 실시간으로 기록 스트림을 게시, 구독, 저장 및 처리할 수 있는 분산 데이터 스트리밍 플랫폼입니다. 이는 여러 소스에서 데이터 스트림을 처리하고 여러 사용자에게 전달하도록 설계되었습니다. 간단히 말해, A지점에서 B지점까지 이동하는 것뿐만 아니라 A지점에서 Z지점을 비롯해 필요한 모든 곳에서 대규모 데이터를 동시에 이동할 수 있습니다.

Apache Kafka는 전통적인 엔터프라이즈 메시징 시스템의 대안입니다. 하루에 1조4천억 건의 메시지를 처리하기 위해 LinkedIn이 개발한 내부 시스템으로 시작했으나, 현재 이는 다양한 기업의 요구사항을 지원하는 애플리케이션을 갖춘 오픈소스 데이터 스트리밍 솔루션이 되었습니다.



# Kafka 설치

[Kafka](https://www.apache.org/dyn/closer.cgi?path=/kafka/2.7.0/kafka_2.13-2.7.0.tgz)에서 다운로드를 받습니다. 다운 받을때의 OS는 WindowOS, MacOS의 차이는 없습니다. 저는 HTTP에서 다운을 하겠습니다.

[JDK Download](https://www.oracle.com/kr/java/technologies/javase-downloads.html)를  하겠습니다. 자신의 운영체제와 원하는 버전을 다운로드 후에 압축을 풀고나서 폴더이름에서 뒤에 숫자부분을 지워주셔야 합니다.





# Kafka 실행 

> 실행 후 터미널을 닫지마세요.
>
> Linux or MacOS는 Window 명령어에서 `/windows/`와 `.bat`을 `.sh`로 변경하시면 됩니다.



### STEP 1: START THE KAFKA ENVIRONMENT

**Windows**

- `> bin\windows\zookeeper-server-start.bat config\zookeeper.properties`

**Linux or MacOS**

- `$ bin/zookeeper-server-start.sh config/zookeeper.properites`



새로운 터미널을 오픈

**Windows**

- `> bin\windows\kafka-server-start.bat config/server.properties`





### STEP 2: CREATE A TOPIC TO STORE YOUR EVENTS

새로운 터미널을 오픈

`Kafka default port -- 9092`, `zookeeper default port -- 2181`

**Windows**

- `> bin\windows\kafka-topics.bat --create --topic <TOPIC_NAME> --bootstrap-server localhost:9092`



- `> bin\windows\kafka-topics.bat --list --bootstrap-server localhost:9092` -- 해당 server의 kafka list 출력



- `> bin\windows\kafka-topics.bat --describe --topic <TOPIC_NAME> --bootstrap-server localhost:9092`  -- 해당 kafka의 상세정보





### STEP 3: WRITE AND READ SOME EVENTS INTO THE TOPIC

Consumer를 실행시키고 새로운 터미널 창을열어서 Producer를 실행시킵니다.

Producer 창에서 원하는 문장을 입력하면 Consumer창에서 문장이 출력이 되는것을 볼 수 있습니다.

**Windows**

- `> bin\windows\kafka-console-consumer.bat --topic <TOPIC_NAME> --from-beginning --bootstrap-server localhost:9092`
  - `--from-beginning` -- 새로운 consumer 창을 실행시키면 기존의 데이터를 다 가져와서 출력시킨다.



- `> bin\windows\kafka-console-producer.bat --topic <TOPIC_NAME> --bootstrap-server localhost:9092`





### STEP 4: DELETE TOPIC

```cmd
> bin\windows\kafka-topics.bat --delete --topic [topic-name] --zookeeper localhost:2181
```

Windows에서 `topic`을 삭제하면 삭제가 되고나서 kafka server가 강제 종료가 되어 버립니다.  kafka server를 다시 실행해도 shutdown이 되기 때문에 `C:\tmp`에서 `kafka-logs`파일을 삭제하고 zookeeper server도 종료 한뒤 `zookeeper`파일도 삭제 해 주셔야 합니다.

이러한 일이 발생하는 이유는 Kafka자체가 Linux기반이기 때문에 Window에서 작은 오류가 발생하는 것 같습니다. 이러한 문제를 해결 하기 위해서는 `topic`을 삭제하기 보다는 새로 만드는 것을 추천 드립니다.





