# 기본 명령어

> `$ docker`가 기본 시작 명령어이다.

#### `$ docker image ls`

- 자신이 가지고 있는 Image 정보



#### `$ docker container ls`

- container 세부 정보. 뒤에 `-a`를 붙이면 종료된 container도 보여준다.
- STATUS : Up - 실행중, Exited - 중지됨



#### `$ docker ps [OPTION]`

- `docker container ls` 랑 같은 기능을 한다.



#### `$ docker images [OPTION]`

- `docker image ls`랑 같은 기능을 한다.



#### `$ docker stop [CONTAINER ID]` 

-  container 중지



#### `$ docker start [CONTAINER ID]`
- container 시작



#### `$ docker rm [CONTAINER ID]` 
- container  삭제



#### `$ docker rmi [IMAGE ID]`

- image 삭제



#### `$ docker run [option] image[:tag|@digest][command][arg...]`
- `run` : pull, create, start가 합쳐진 기능
  - `-d` : detached mode 흔히 말하는 백그라운드 모드
  - `-p` : port forward 호스트와 컨테이너의 포트를 연결 (포워딩)
  - `-v` : 호스트와 컨테이너의 디렉토리를 연결 (마운트)
  - `-e` : 컨테이너 내에서 사용할 환경변수 설정
  - `-it` : `-i`(interective)와 `-t`(tty)를 동시에 사용하는 것으로 터미널에 접속하고 입력을 위한 옵션
  - `--name` : 컨테이너 이름 설정
  - `--rm` : 프로세스 종료시 컨테이너 자동 제거
  - `--link` : 컨테이너연결 [컨테이너명:별칭]



#### `$ docker exec -it [container_id] /bin/bash`

- container 내부에 접속
  
   - `-i`가 없이 `-t`만 있으면 컨테이너 안에서 입력이 불가능
   
   - `-t`없이 `-i`만 있으면 접속이아닌 일반 문자입력만 가능 그러므로 `-it` 같이 쓰는게 정석
   - `/bin/bash`는 컨테이너 실행 명령어



#### `$ docker inspect [container_id, name]`

- 컨테이너 상세정보



#### `$ docker logs [container_id, name]`

- container를 만들때 생기는 logs들과 데이터logs들을 출력해준다.



#### `$ docker system prune`

- 쓰지않는(Stoped) container, image, volume data를 모두 삭제
  - 하나만 삭제하고싶다면 `docker volume prune`처럼 입력



#### `$ docker network [OPTION]`

- docker 에서 container를 연결하고 관리하는 기능으로써 예를들어`ls`를 붙이면 목록을 보여주고 `inpect [network_name]`을 붙이면 해당 network의 상세정보를 보여준다.
- 자세한 사용법은 [여기](https://github.com/jaden7856/TIL/blob/master/Docker/05_Docker_network.md)에 적어놓도록 하겠습니다.