# Django를 Docker로 연결

제가 소소하게 작업한 Django를 활용하여 웹사이트를 만들고 Scrapy를 사용한 프로젝트를 Docker로 옮기고 Docker에서 사용하는 것을 알아보겠습니다.



### 기존 django 프로젝트를 복사

Local에서 작업하던 프로젝트를 복사하여 docker 폴드로 옮기던지 그 프로젝트 폴더에 Docker를 연결해도 상관없습니다. 저는 MySQL을 사용하여 Django를 연결할 예정입니다. 그러므로 우선 docker에서 mysql을 설치하도록 하겠습니다.

```shell
docker run -d -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=true -e MYSQL_DATABASE=[DB_NAME] --name mysql_server -v [Local 작업한 폴더 URL]:/var/lib/mysql mysql:5.7
```

- mysql을 `--volume`을 사용하여 Local과 연결하도록 합니다.
- `MYSQL_ALLOW_EMPTY_PASSWORD`: 비밀번호를 입력하지 않아도 되도록 설정
- `MYSQL_DATABASE` : 새로운 Database를 SQL단에서 만들지 않고 바로 생성하도록 설정



## 데이터베이스 변경

> sqlite3 -> mysql(docker) 
>
> MySQL을 쓰고있거나 다른 DB를 쓰고있다면 하지않아도 된다.

1. Network 생성
2. MySQL, Django 컨테이너를 같은 네트워크에 추가

3. django 기본 project의 settings.py에서 DATABASES부분을 찾는다.

```python
DATABASES = {
    'default': {
        # 'ENGINE': 'django.db.backends.sqlite3',
        # 'NAME': os.path.join(BASE_DIR, 'db.sqlite3'),
        'ENGINE': 'django.db.backends.mysql',
        'NAME': '[DB_NAME]',
        'USER': '[USER_NAME]',
        'PASSWORD': '',
        'HOST': '[MySQL_CONTAINER_NAME]',
        'PORT': '3306',
    }
}
```

- 기존에 연결하고있던 SQLite3를 주석 처리하고 MySQL로 연결을 한다.



4. Model을 바꿨으니 적용하도록 합시다.

   - `$ python manage.py migrations`
   - `$ python manage.py migrate`
   - `$ python manage.py createsuperuser`

   

### Django 정상작동 확인작업

Django Project파일이 Docker에서 실행을 할때 필요한 Library들과 작업들이 무엇이 필요한지 확인을해야합니다.

우선 확인 작업을 위해 Image `build`를 합니다.

0. Dockerfile 생성

   - Dockerfile에서 base가 될 것을 선택하고 확인작업 이므로 간단히 적고 넘어간다.

     ```dockerfile
     FROM python:3.7.9-stretch
     
     CMD ["df"]
     ```

     

1. Image 생성

   - **Dockerfile이 위치한 경로에서 시작**

   - `$ docker build -t [Image_name] .`

     

2. container 생성

   - `$ docker run -v [LOCAL_URL]:[DOCKER_URL] -it -P [Image_name] /bin/bash`

     

3. django 실행을 위한 Lib 확인 후 설치

   **Project에서 했던 작업이 많을수록 설치할 것이 많아지기에 미리 메모해두거나 기억하자**

   - `# pip install django`
   - `# pip install mysqlclient`

   

4. 실행

   - `python manage.py runserver 0.0.0.0:8000` 
     - `0.0.0.0:8000` 을 해 줘야 Django에서 외부 ip접속에대한 권한을 허용한다.



### dockerfile 생성

정상 작동한다면  `Dockerfile`에  Project를 실행하기 위해 설치했던 것들을  추가합시다. 저는 `django`, `mysqlclient`를  설치하면 정상작동하기에 두개를 추가하고 Docker 내에서 수정할 수도 있기에 vim을 설치하였습니다.

```dockerfile
FROM python:3.7.9-stretch

WORKDIR /mydjango

COPY bookmark ./bookmark
COPY mysite2 ./mysite2
COPY manage.py ./manage.py
COPY db.sqlite3 ./db.sqlite3

RUN apt-get update
RUN apt-get install -y vim
RUN pip install django
RUN pip install mysqlclient

EXPOSE 8000

CMD [ "python", "manage.py", "runserver", "0.0.0.0:8000" ]
```



#### Dockerfile을 만든 경로로 이동을 하고 Image를 `build`합니다.

- `docker build -it [Image_name] .`

- `docker run -v [Dockerfile을 만든 경로 URL]:/mydjango -P [Image_name]`
  - `-P` 의 옵션으로 자동 연결된 Port번호를 검색하여 **127.0.0.1:[port_number] **에 접속!



## Network 연결

Django의 Container와 MySQL Container를 같은 Network에 연결을 해야한다.



#### Network 생성

- `$ docker network create [NETWORK_NAME]`



#### NETWORK 연결

MySQL, Django 두개의 Container를 생성한 Network에 연결

- `$ docker network connect [NETWORK_NAME] [CONTAINER_ID, NAME]`



#### 확인

`inspect`하여 Container가 두개 다 정상적으로 연결이 되어있는지 확인한다.

- `$ docekr network inspect [NETWORK_NAME]`



그다음 Django Container를 실행하여 웹사이트에 접속해 확인작업





## Hub 사이트에 공개

Hub사이트에 배포를 하기위해선 Image의 이름앞에 자신의 계정명을 입력해야한다. `[계정명]/[Image_name]` 

처음부터 Image를 생성할때 이름을 규칙에맞게 생성하거나 기존의 Image를 바꾸는 방법은 아래처럼 기존 이미지이름과 뒤에 바꿀 이름을 입력한다.

```dockerfile
docker tag [Image_name]:[tag] [rename_Image]:[tag]
```



배포는 Git과 유사하다.

```dockerfile
docker push [MY_ID]/[Image_name]:[Tag]
```



### Private Registry 사용

`$ docker pull registry`

`$ docker run -d -p 5000:5000 --restart always --name registry registry:latest`

```dockerfile
$ docker pull ubuntu
$ docker tag ubuntu localhost:5000/ubuntu:latest
$ docker push localhost:5000/ubuntu:latest
```

