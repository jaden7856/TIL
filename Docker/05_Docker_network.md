# Network

Docker 컨테이너(container)는 격리된 환경에서 돌아가기 때문에 기본적으로 다른 컨테이너와의 통신이 불가능합니다. 하지만 여러 개의 컨테이너를 하나의 Docker 네트워크(network)에 연결시키면 서로 통신이 가능해집니다. 이번 포스트에서는 컨테이너 간 네트워킹이 가능하도록 도와주는 Docker 네트워크에 대해서 알아보도록 하겠습니다.



## 네트워크 종류

Docker 네트워크는 `bridge`, `host`, `overlay` 등 목적에 따라 다양한 종류의 네트워크 드라이버(driver)를 지원하는데요.

- `bridge` 네트워크는 하나의 호스트 컴퓨터 내에서 여러 컨테이너들이 서로 소통할 수 있도록 해줍니다.
- `host` 네트워크는 컨터이너를 호스트 컴퓨터와 동일한 네트워크에서 컨테이너를 돌리기 위해서 사용됩니다.
- `overlay` 네트워크는 여러 호스트에 분산되어 돌아가는 컨테이너들 간에 네트워킹을 위해서 사용됩니다.



### 네트워크 조회

- `$ docker network ls`



### network 생성

- `$ docker network create [network_name]`



### network 연결

- `$ docker network connect [network_name] [container_name, id]`
  - 새로 생성한 network에 container를 연결
  - `run`을 할때 처음부터 network를 연결 할 수도 있다.
    - `$ docker run --network [network_name] [Image_name]:[Tag]` 



### network의 정보를 검색

- `$ docker inspect [network_name]`





### 접속

- `$ docker exec -it [container_id, name] /bin/sh`

  - 연결한 container에 접속

  

- `$ ping [connected_container_name]`
  - network에 연결된 여러 container중에 연결이 되어있는지 확인을 위해 ping을 걸어본다.



#### Network 연결 해제 및 삭제

- `$ docker network disconnect [network_name] [container_name, id]` 
  - network에 연결된 container를 해제
- `$ docker network rm [network_name]` 
  - network 삭제
- `$ docker network ls` 
  - 상태 확인