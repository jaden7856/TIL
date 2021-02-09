# MySQL 생성 후 접속

> **중요!!**

1. `$ docker run -d -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=true --name [container_name] mysql:5.7`
   - `-d` : detached -> 백그라운드 실행
     - 만약 하지않으면 여러 로그들이 출력이된다.
     - `-p 3306:3306` : [host:container] HostOS와 container의 mysql을 연결
         - 예를들어 window에서 127.0.0.1:3306을 사용하면 container안에 mysql 3306이랑 연결
   - `-e` : 환경 변수 전달, `MYSQL_ALLOW_EMPTY_PASSWORD=true` : 패스워드 생성 없음 



**2. Container 접속**
   - `$ docker container exec -it [container_id, container_name] /bin/bash`

- 접속하면서 뒤에 명령어를 적어도 됩니다.
  - `$ docker container exec -it mysql mysql -h127.0.01 -uroot -p`
    - 만약 error가 발생한다면
  - `$ docker container exec -it mysql bash -c mysql -h127.0.01 -uroot -p`



**3. MySQL 접속**

- `/# mysql -h127.0.0.1 -uroot -p`
  - `-h` 내부의 다른 서버에 접속할때 쓰인다, `-u` : user_name, `-p` : password를 물어보는 옵션
    - `-uroot`처럼 `-u`뒤에 바로 붙여도 되고 `-u root`처럼 띄어도 상관은 없습니다.
  - 비밀번호 입력란에 그냥 엔터를 하면된다.



**다른 Container끼리 접속 방법**

- `docker run`할떄  `-p 13306:3306 --name db1`과 `-p 23306:3306 --name db2`을 만든다.
- `$ docker inspect db1`과 `db2`를 하여 맨 밑부분에 `IPAddress: 172.17.0.?`를 찾는다.
- `db2`에서 `db1`과 연결하기위해 `db2` 컨테이너에 `exec`후에 `mysql -h[db1 IPAddress] -u[] -p[]`를 입력
- `db2`가 `db1`에 접속되었습니다.
