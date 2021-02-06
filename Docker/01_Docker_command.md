# 기본 명령어

> `$ docker`가 기본 시작 명령어이다.

#### `$ docker image ls`

- 자신이 가지고 있는 Image 정보



#### `$ docker container ls`

- container 세부 정보. 뒤에 `-a`를 붙이면 종료된 container도 보여준다.
- STATUS : Up - 실행중, Exited - 중지됨



#### `$ docker stop [CONTAINER ID]` 

-  container 중지



#### `$ docker start [CONTAINER ID]`
- container 시작



#### `$ docker rm [CONTAINER ID, IMAGE ID]` 
- container , image 삭제



#### `$ docker run [option] image[:tag|@digest][command][arg...]`
- `run` : pull, create, start가 합쳐진 기능
  - `-d` : detached mode 흔히 말하는 백그라운드 모드
  - `-p` : port forward 호스트와 컨테이너의 포트를 연결 (포워딩)
  - `-v` : 호스트와 컨테이너의 디렉토리를 연결 (마운트)
  - `-e` : 컨테이너 내에서 사용할 환경변수 설정
  - `-it` : -i와 -t를 동시에 사용하는 것으로 터미널에 접속하고 입력을 위한 옵션
  - `--name` : 컨테이너 이름 설정
  - `--rm` : 프로세스 종료시 컨테이너 자동 제거
  - `--link` : 컨테이너연결 [컨테이너명:별칭]



#### `$ docker inspect [container_id, name]`

- 컨테이너 상세정보