# 기본 개념 이해

### 1) 구조와 오브젝트

- 클러스터 구조

  - 마스터 / 머신 노드가 존재.
  - 마스터는 하나, 머신은 여러개

- 쿠버네티스 오브젝트

  - 기본적으로 오브젝트와 컨트롤러로 구성

  - 오브젝트

    - yaml 이나 json 으로 정의된 스펙에 따라 만들어지는 객체

  - 예를 들면 다음과 같이 스펙을 정의

  -  

    ```
    # pod_example.yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 8090
    ```

  - `kubectl create -f pod_example.yaml` 로 오브젝트 생성

    

### 2) 기본 오브젝트 종류

- **Pod**

  - 가장 기본적인 배포 단위
  - 컨테이너들을 포함하고 있음

  

- **Volume**

  - 컨테이너의 외부 디스크
  - 컨테이너가 재 실행되어도 Volume 을 사용하면 파일 유지

  

- **Service**

  - 여러 개의 Pod 들을 서비스할 때, 현재 요청이 어느 Pod 으로 갈지 선택하는 오브젝트
  - 부하가 많을 때 이를 분산시키는 로드밸런서 역할

  

- **Controller**

  - 여러 개의 Pod 배포를 적절하게 관리하는 오브젝트
  - Pod 들을 생성, 삭제함.
    - Replication Controller
    - Replication Set
  - Deployment

  

- 그 외 개념

  - **네임 스페이스**

    - 한 클러스터 내에 논리적인 분리 단위
    - 예를 들면 `namespace:billing` 과 `namespace:commerce` 는 같은 클러스터 내에 있지만 논리적으로 분리됨
    - 한 클러스터 자원을 가지고 개발 / 운영 / 테스트 식으로 나눌 수 있음.

    

  - **라벨**

    - 리소스를 선택하는데 사용됨

    - 예를 들면 Pod 안에 다음과 같이 라벨을 달 수 있음.
  
    - ```
      kind: Pod
      ...
      metadata:
        labels:
        app: myapp
      ```

    - selector 를 사용하여 리소르를 선택함

    - 예를 들어, Service 는 다음과 같이 Pod 을 선택함
  
    - ```
      kind: Service
      ...
      spec:
        selector:
          app: myapp
      ```



## 1. 아키텍처

쿠버네티스 내부 구조는 크게 마스터와 노드로 분리될 수 있다.

### 1) 마스터

쿠버네티스 클러스터 전체를 컨트롤하는 시스템이다.
다음과 같이 구성되어 있다.

- API 서버
  - 모든 명령과 통신을 REST API 로 제공
- Etcd
  - 분산형 키-밸류 스토어로, 쿠버네티스 클러스터 상태나 설정 정보 저장
- 스케쥴러
  - Pod, Service 등 각 리소스를 적절한 노드에 할당
- 컨트롤러 매니저
  - 컨트롤러들을 생산, 배포 등 관리
- DNS
  - 동적으로 생성되는 Pod, Service 등의 IP 를 담는 DNS

### 2) 노드

마스터에 의해 명령을 받고 실제 서비스하는 컴포넌트다.
다음과 같이 구성되어 있다.

- Kubelet
  - 마스터 API 서버와 통신하는 노드 에이전트
- Kube-proxy
  - 노드로 오는 네트워크 트래픽을 적절한 컨테이너로 라우팅
  - 네트워크 트래픽 프록시 등 노드-마스터간 네트워크 통신 관리
- Container runtime
  - 컨테이너를 실행 (ex. dcoker)
- cAdvisor
  - 각 노드 내 모니터링 에이전트
  - 노드 내 컨테이너들의 상태, 성능 수집하여 마스터 API서버에 전달



## HealthCheck

HealthCheck 는 **각 컨테이너 상태를 주기적으로 문제가 있는지 체크**하는 기능이다.

- 헬스 체크 결과, 이상이 감지되는 경우에는 컨테이너를 강제 종료하고 재시작할 수 있음
- kubelet이 컨테이너의 헬스 체크를 담당



### 방법

- Liveness probe

  - 컨테이너의 어플리케이션이 정상적으로 실행 중인 것을 검사
  - 검사에 실패하면 포드상의 컨테이너를 강제로 종료하고 재시작
  - 매니페스트에 명시적으로 설정해야 사용할 수 있음

- Readiness probe

  - 컨테이너의 어플리케이션이 요청을 받을 준비가 되었는지 검사
- 검사에 실패하면 서비스에 의한 요청 트래픽 전송을 준비
  - 포드가 기동되고 나서 준비가 될 때까지 요청이 전송되지 않도록 하기 위해 사용
  - 매니페스트에 명시적으로 설정해야 사용할 수 있음
  
  

### 1. 작업 디렉터리 생성 후 매니페스트 파일(YAML) 작성

- **`[vagrant@master]$ vi webapl-pod.yaml`**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: webapl
spec:
  containers:
    - name: webapl
      image: myanjini/webapl:0.1     # 핸들러를 구현한 어플리케이션
      livenessProbe:                 # 어플리케이션이 동작하지 여부를 확인
        httpGet:                     #   지정된 포트와 경로로 HTTP GET 요청을 주기적으로 실행  
          path: /healthz             #     확인 경로
          port: 3000                 #     확인 포트
        initialDelaySeconds: 3       # 검사 개시 대시 시간
        periodSeconds: 5             # 검사 주기
      readinessProbe:                # 어플리케이션이 준비되었는지 확인
        httpGet:
          path: /ready
          port: 3000
        initialDelaySeconds: 15
        periodSeconds: 6

```



### 2. 컨테이너 이미지를 정의

- **`[vagrant@master]$ vi Dockerfile`**

```dockerfile
FROM alpine:latest

RUN apk update && apk add --no-cache nodejs npm

WORKDIR /
ADD ./package.json /
RUN npm install
ADD ./webapl.js /

CMD node /webapl.js
```



- **`[vagrant@master]$ vi package.json`**

```json
{
    "name": "webapl",
    "version": "1.0.0",
    "description": "",
    "main": "webapl.js",
    "scripts": {
      "test": "echo \"Error: no test specified\" && exit 1"
    },
    "author": "",
    "license": "ISC",
    "dependencies": {
      "express": "^4.16.3"
    }
  }
```



- **`[vagrant@master]$ vi webapl.js`**

```js
//  웹 어플리케이션 
const express = require('express')
const app = express()
var start = Date.now()               // 어플리케이션이 시작된 시간

//  http://CONTAINER_IP:3000/healthz 형식으로 요청이 들어왔을 때 수행하는 기능을 정의하는 함수
app.get('/healthz', function(request, response) {
    var msec = Date.now() - start    // 어플리케이션이 시작된 후 경과된 시간
    var code = 200
    if (msec > 40000 ) {      // 경과된 시간이 40초 보다 작으면 200을, 크면 500을 응답코드로 반환
    code = 500
    }
    console.log('GET /healthz ' + code)
    response.status(code).send('OK')
})

app.get('/ready', function(request, response) {
    var msec = Date.now() - start
    var code = 500
    if (msec > 20000 ) {
    code = 200
    }
    console.log('GET /ready ' + code)
    response.status(code).send('OK')
})

app.get('/', function(request, response) {
    console.log('GET /')
    response.send('Hello from Node.js')
})

app.listen(3000);
```



### 3. 컨테이너 이미지 생성 후 도커허브에 등록

```shell
[vagrant@master]$ docker build --tag [USER_ID]/webapl:[TAG] .
```

```sh
[vagrant@master]$ docker push [USER_ID]/webapl:[TAG]
```



### 4. 포드를 배포하고 헬스체크 기능을 확인

```sh
[vagrant@master hc-probe]$ kubectl apply -f webapl-pod.yaml
pod/webapl created
```

```sh
[vagrant@master hc-probe]$ kubectl get pods
NAME           READY   STATUS    RESTARTS   AGE
webapl         1/1     Running   0          30s
```

```sh
[vagrant@master hc-probe]$ kubectl logs webapl
GET /healthz 200
GET /healthz 200
GET /healthz 200
GET /ready 500   	15 + 6초  
GET /healthz 200	20초
GET /ready 200    	15 + 12초 → 20초 초과	⇒ READY 상태 1/1가 설정
GET /healthz 200
GET /ready 200
GET /healthz 200
	:
```

