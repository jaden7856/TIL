# Docker-compose

애플리케이션 간의 연동 없이는 실용적 수준의 시스템을 구축할 수 없다. 다시말하면, 도커 컨테이너로 시스템을 구축하면 하나 이상의 컨테이너가 서로 통신하며, 그 사이에 의존관계가 생긴다.

이런 방식으로 시스템을 구축하다 보면 **단일 컨테이너를 다룰 때는 문제가 되지 않던 부분에도 주의가 필요**하다. 컨테이너의 동작을 제어하기 위한 설정 파일이나 환경 변수를 어떻게 전달할지, **컨테이너 간의 의존관계를 고려할 때 포트 포워딩을 어떻게 설정해야 하는지 등의 요소를 적절히 관리해야 한다.**

도커를 사용해서 이런 시스템을 만들려면 기본 사용법만 알아서는 부족할 것이다. 이때 필요한 것이 도커 컴포즈(Docker Compose)다. Compose는 yaml 포맷으로 기술된 설정 파일로, **여러 컨테이너의 실행을 한 번에 관리할 수 있게 해준다.**



# Docker-compose 명령으로 컨테이너 실행

윈도우용/macOS용 도커가 로컬 환경에 설치돼 있다면 docker-compose명령을 바로 사용할 수 있다.

```shell
$ docker-compose version
docker-compose version 1.27.4, build 40524192
```



먼저 컨테이너 하나를 실행해 보고 같은 작업을 docker-composer를 사용해 다시 수행할 것이다. 저는 MySQL을 생성하겠습니다.

```
$ docker run -d -p 3306:3306 -e MYSQL_DATABASE=[생성할 DB이름] -e MYSQL_ROOT_PASSWORD=[PASSWORD] -v [LOCAL URL]:/var/lib/mysql mysql:5.7
```

정상적으로 작동을 하면 컨테이너를 다시 삭제합니다.



임의의 디렉토리에서 `docker-compose.yml` 라는 파일명으로 다음과 같이 임의로 작성한 내용입니다.

```yaml
version: "3.9"
service:
		servicename:
				build: . # image build를 할꺼면 생성
				image: # optioanl
				command: # optional
				environment: # optioanl
				volumes: # optioanl
				...
		servicename2: # if have second service...
		
volumes: # optional
network: # optional
```



저는 방금전에 작성했던 `run`코드를 통하여 `mysql`을 만들어보겠습니다.

```yaml
version: "3.9"

services:
  my-mysql:
    image: mysql:5.7
    volumes:
      - ./mysql-data:/var/lib/mysql
    ports:
      - 3306:3306
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: mydb
```



### Docker compose 실행

docker-compose.yml 파일이 위치한 디렉토리에서 이 정의에 따라 여러 컨테이너를 한꺼번에 시작하려면

```
$ docker-compose up
$ docker-compose up --build  # Dockerfile을 다시 빌드
```



### Docker compose stop

```
$ docker-compose down
```

`docker-compose down`명령을 사용하면 docker-compose.yml 파일에 정의된 모든 컨테이너가 정지 혹은 삭제된다.

