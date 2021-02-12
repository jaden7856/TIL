# Network

Docker 컨테이너(container)는 격리된 환경에서 돌아가기 때문에 기본적으로 다른 컨테이너와의 통신이 불가능합니다. 하지만 여러 개의 컨테이너를 하나의 Docker 네트워크(network)에 연결시키면 서로 통신이 가능해집니다. 이번 포스트에서는 컨테이너 간 네트워킹이 가능하도록 도와주는 Docker 네트워크에 대해서 알아보도록 하겠습니다.



### 네트워크 종류

Docker 네트워크는 `bridge`, `host`, `overlay` 등 목적에 따라 다양한 종류의 네트워크 드라이버(driver)를 지원하는데요.

- `bridge` 네트워크는 하나의 호스트 컴퓨터 내에서 여러 컨테이너들이 서로 소통할 수 있도록 해줍니다.
- `host` 네트워크는 컨터이너를 호스트 컴퓨터와 동일한 네트워크에서 컨테이너를 돌리기 위해서 사용됩니다.
- `overlay` 네트워크는 여러 호스트에 분산되어 돌아가는 컨테이너들 간에 네트워킹을 위해서 사용됩니다.



# 컨테이너 간 통신

디폴트 네트워크 안에서 컨테이너 간의 통신에서는 서비스의 이름이 호스트명으로 사용됩니다.

예를 들어, `Django(WEB)` 서비스의 컨테이너에서 `MriaDB(DB)` 서비스의 컨테이너를 대상으로 `ping` 명령어를 날릴 수 있습니다. 

##### `$ docker-compose exec [Django_NAME] ping [MariaDB_NAME]`



컨테이넌 간 통신에서 주의할 점은 접속하는 위치가 디폴트 네트워크 내부냐 외부냐에 따라서 포트(port)가 달라질 수 있다는 것입니다.

##### `$ run -p 8001:800 [Image_name]:[Tag]`

호스트 컴퓨터에서 접속할 때는 `8001` 포트를 사용해야 하고, 같은 디폴트 네트워크 내의 다른 컨테이너에서 접속할 때는 포트 `8000`을 사용해야 합니다.



- 호스트 컴퓨터에서 `web` 서비스 컨테이너 접속
  - `$ curl -I localhost:8001`
- 같은 네트워크 내의 다른 컨테이너에서 `web` 서비스 컨테이너 접속
  - `$ docker-compose exec [CONTAINER_NAME] curl -I web:8000`



# 외부 네트워크 사용

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