# Dockerfile 자주쓰는 명령어

### FROM

```dockerfile
FROM [--platform=<platform>] <image> [AS <name>]
```

또는

```dockerfile
FROM [--platform=<platform>] <image>[:<tag>] [AS <name>]
```

또는

```dockerfile
FROM [--platform=<platform>] <image>[@<digest>] [AS <name>]
```

`FROM` 명령어는 새로운 빌드 단계를 준비하고 다음 명령어들의 *기본 이미지* 를 지정한다. 따라서 유효한 `Dockerfile` 은 `FROM` 명령어로 시작해야 한다.



### RUN

`RUN` 은 두 가지 형태가 있다.

- `RUN <command>` (기본적으로 리눅스에서는 `/bin/sh` 로 실행되고 윈도우에서는 `cmd /S /C` 로 실행된다. 따라서 커맨드는 그 형식에 맞게 작성되어야 한다)
- `RUN ["executable", "param1", "param2"]` (*exec* 형식)

`RUN` 명령어는 현재 이미지의 위 새 래이어에서 실행되고 결과를 커밋한다. 커밋된 이미지의 결과물은 `Dockerfile` 의 다음 스탭에서 사용된다.



### CMD

`CMD` 명령어는 다음 세 가지 형식이 있다.

- `CMD ["executable", "param1", "param2"]` (*exec* 형식, 이 형식을 자주 사용한다)
- `CMD ["param1", "param2"]` (*기본 파라미터를 `ENTRYPOINT` 로 갖는다.*)
- `CMD command param1 param2` (*shell* 형식)

**`CMD` 명령의 주 목적은 실행중인 컨테이너에 기본 환경을 제공하기 위함이다.**





# Dockerfile 작성

1. Dockerfile 생성

   - `$ touch Dockerfile`

   

2. FROM

   - `$ gedit Dockerfile`
     - base의 위치를 수정



#### Image build

- `$ docker image build --tag fromtest:1.0 .`



- 만약 같은 tag를 생성할때 CACHED를 지우고싶다면..
  - `$ docker build --no-cache=true -t addtest:1.0 .`



**바로 VisualStudio를 이용하여 편리하게 사용하자! **

1. `$ code Dockerfile`

   

2. VScode에서 Docker를 설치



3. Dockerfile에서 `FROM [Image_name]:[Tag]`

- ex) `FROM ubuntu:1.0`

  

4. `$ docker image build --tag fromtest:1.0 .`



### 추가 작업

- `$ code Dockerfile`

  ```dockerfile
    FROM ubuntu:latest
    
    RUN mkdir /mydata
    COPY test.txt /mydata   # test.txt안에 Hello, Docker!
    
    CMD ["df", "-h"]  # ,는 command에서 df -h 처럼 한칸 띄우기랑 똑같다.
  ```

- `ADD`는 단순히 파일을 호스트에서 컨테이너로 복사하는 기능뿐만 아니라
  추가기능이 있었는데, 그 추가기능이 문제가되서 단순히 복사만 하는 `COPY` 명령어가 만들어짐

- `CMD` 명령어는 제일 마지막에 작성해야하며 프로세서를 작동시키는 명령어이다.

  - `ENTRYPOINT`도 비슷한 명령어 이지만 `CMD`가 파라미터를 사용하거나 가변데이터 사용에 적합하다.



## Volume

Volume Mount는 windows나 MacOS같이 HostPC에서 작업한 파일이 수정, 삭제, 추가 같은 작업이 발생하면 Docker Image파일과 Container파일도 새로 `build`를 하고 `run`을 해 줘야하는 것을 연결된 디렉토리 끼리는 자동으로 업데이트가 되어 **HostPC에서 변경된 작업이 Docker내에서 자동으로 변경이 된다.**

1. dockerfile을 작성합니다.

```dockerfile
FROM python:3.7.9-stretch

WORKDIR /mydata
RUN pip install numpy

CMD python ${EXEC_FILE}
```

2. `$ docker build -t [Image_name]:[Tag] .`



3. `$ docker run -it -v [Windows 파일 URL]:[Container 파일 URL] -e EXEC_FILE=test.py [Image_name]`
   - ex) `-v c:/work/my_test:/mydata`



# Node.js를 실행하는 Dockerfile 작성

1. Node project 폴더 생성

   - `$ mkdir node_project`

   

2. package.json 과 index.js 파일 작성

   - ```json
     // package.json
     
     {
         "dependencies": {
             "express": "*"
         },
         "scripts": {
             "start": "node index.js"
         }
     }
     ```

   - ```js
     // index.js
     
     const express = require('express');
     
     const app = express();
     
     app.get('/', (req, res) => {
         res.send('How are you doing');
     });
     
     app.listen(8080, () => {
         console.log('Listening on port 8080');
     });
     ```



4. nodejs 설치

   - `$ apt-get install nodejs`
   - `$ yum install nodejs`

   - `$ npm install`
   - `$ npm start`

   

5. 적상작동 한다면 Dockerfile 생성

   ```dockerfile
   FROM node:alpine
   
   WORKDIR /mydata
   
   COPY ./package.json ./package.json
   COPY ./index.js ./index.js
   
   EXPOSE 8080
   RUN npm install
   
   ENTRYPOINT ["npm", "start" ]
   ```

- `EXPOSE`를 사용하면 같은 Networks를 쓰는 컨테이너끼리는 연결 할 수 있다.



6. Image 생성

   - `$ cd 파일 생성한 장소`

   - `$ docker build -t mynodejs .`

   

7. Container 실행

   - HostPC와 container가 연결이 되기위해 `-p` 옵션을 사용한다.
- `$ docker run -p 8080:8080 mynodejs`



## 배포

Image의 이름앞에 자신의 계정명을 입력해야한다. `[계정명]/[Image_name]` 

기존의 Image를 바꾸는 방법은 아래처럼 기존 이미지이름과 뒤에 바꿀 이름을 입력한다.

```dockerfile
docker tag [Image_name]:[tag] [rename_Image]:[tag]
```



배포는 Git과 유사하다.

```dockerfile
docker push [ID]/[Image_name]:[Tag]
```

