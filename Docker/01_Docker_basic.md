# 기본 명령어

> `$ docker`가 기본 시작 명령어이다.

- `$ docker image ls`

  - 자신이 가지고 있는 Image 정보

  

- `$ docker container ls`
  - container 세부 정보. 뒤에 `-a`를 붙이면 종료된 container도 보여준다.
  - STATUS : Up - 실행중, Exited - 중지됨

![도커명령어1](01_Docker_basic.assets/도커명령어1.PNG)



- `$ docker stop <CONTAINER ID>` : container 중지
- `$ docker start <CONTAINER ID>` : container 시작

- `$ docker rm <CONTAINER ID>` : container 삭제

- `$ docker image rm <IMAGE ID>` : Image 삭제