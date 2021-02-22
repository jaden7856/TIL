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



# 기본 오브젝트

### 1. Types (Kinds)

#### 1) Pod

쿠버네티스에서 가장 기본적인 배포 단위다.
하나 이상의 컨테이너들을 포함한다.

예를 들어 다음과 같이 `yaml` 포맷으로 정의할 수 있다.

```
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

Pod 은 다음과 같은 특징을 지닌다.

- Pod 내 컨테이너들은 같은 IP 를 가지고, Port 를 공유한다.
- Pod 내 컨테이너들끼리는 Volume 을 공유할 수 있다.



### 2) Controller

Controller 의 주 역할은 **Pod 을 생성, 관리하는 것**이다.
예를 들어 다음과 같이 `yaml` 포맷으로 정의할 수 있다.

```
apiVersion: v1
kind: ReplicationController
metadata:
  name: nginx
spec:
  replicas: 3
  selector:
    app: nginx
  template:
    metadata:
      name: nginx
      labels:
        app: nginx
      spec:
        containers:
          - name: nginx
            image: nginx
            ports:
            - containerPort: 80
```

종류는 다음과 같다.

- ReplicationController (RC)

  - 지정된 숫자로 Pod 을 기동 시키고 관리한다.

    - `selector` : `label` 기준으로 어떤 Pod 들을 관리할지 정의한다.
    - `replicas` : 선택된 Pod 들을 몇 개의 인스턴스로 띄울지 정의한다.
    - `template` : `Pod` 을 추가로 기동할 때, 어떤 Pod 을 만들지 정의한다.

    

- ReplicaSet

  - ReplicationController 는 Equailty 기반 Selector 를 사용하는 반면
  - ReplicaSet 는 Set 기반 Selector 를 사용한다.

  

- Deployment

  - ReplicationController 와 ReplicaSet 을 좀더 추상화한 개념
  - 실제 배포할 때는 이 컨트롤러를 주로 사용

  

- DaemonSet

  - Pod 이 각각의 노드에서 하나씩만 돌게 한다. (균등하게 배포)
  - 보통 서버 모니터링이나 로그 수집 용도
  - 모든 노드가 아닌 특정 노드들만 선택할 수도 있다.

  

- Job

  - 한번 실행되고 끝나는 Pod 을 관리한다.
  - Job 컨트롤러가 종료되면 Pod 도 같이 종료한다.
  - 컨테이너에서 Job 을 수행하기 위한 별도의 `command` 를 준다.
  - Job `command` 의 성공 여부를 받아 재실행 또는 종료여부를 결정한다.

  

- CronJob

  - 주기적으로 돌아야하는(배치) Pod 을 관리한다.
  - 별도의 `schedule` 을 정의해아 한다.

  

- Stateful

  - DB 와 같이 상태를 가지는 Pod 을 관리한다.

  

### 3) Service

Service 는 같은 어플리케이션을 운용하는 **Pod 간의 로드 밸런싱 역할**을 한다.
또, 동적으로 생성되는 Pod 들의 동적 IP 와 달리 Service 는 **지정된 IP 로 생성가능**하다.
(즉 Pod 접근은 어려운 반면, Service 접근은 더 쉬움)

다음과 같은 기능들이 있다.

- 멀티 포트 지원

  - 예를 들어 80 -> 8080 으로, 443 -> 8082 로 가도록 한번에 설정할 수 있다.

  

- 로드 밸런싱

  - 부하(트래픽)를 여러 Pod 으로 분산한다.
  - Pod 은 기본적으로 랜덤하게 선택된다.

  

- IP 주소 할당 방식과 연동 서비스에 따른 Type

  - Cluster IP

    - 디폴트 값
    - 서비스에 클러스터 내부 IP 를 할당
    - 즉 클러스터 내부 접근 O, 외부 접근 X

  - Load Balancer

    - 외부 IP 를 가지고 있는 로드밸런서를 할당
    - 즉 외부 접근 O

  - NodePort

    - 클러스터 내 노드의 `ip:port` 로도 접근가능하게 함
    - ex. `curl 10.146.0.10:30036`
    - `10.146.0.10` 는 노드의 ip 고, `30036` 는 NodePort 로 설정한 포트임

  - ExternalName

    - 외부 서비스를 쿠네터네스 내부에서 호출하고자 할 때 사용
    - 모든 Pod 들은 Cluster IP 를 가지고 있기 때문에, 외부에서도 접근이 가능함.
    - 요청 -> (외부 서비스) -> 클러스터 내 쿠버네티스. 일종의 프록시 역할

    

- headless 서비스

  - 서비스 디스커버리 솔루션을 제공하는 경우, 서비스의 IP 가 필요 없음
  - `clusterIP: None` 으로 주면 된다.

  

- External IP

  - 서비스에 별도의 외부 IP 를 지정해줄 수 있음




### 2. Spec

#### 1) Volume

Volume 은 **Pod 에 종속**된 디스크다.
따라서 **Pod 내 여러 컨테이너들이 공유해서 사용할 수 있다.**

예를 들면 다음과 같이 Pod 을 정의할 때 `volumes` 를 통해 정의할 수 있다.

```
apiVersion: v1
kind: Pod
metadata:
  name: shared-volumes 
spec:
  containers:
  - name: redis
    image: redis
    volumeMounts:
    - name: shared-storage
      mountPath: /data/shared
  - name: nginx
    image: nginx
    volumeMounts:
    - name: shared-storage
      mountPath: /data/shared
  volumes:
  - name : shared-storage
    emptyDir: {}
```

종류는 다음과 같다.

- 임시 디스크(Pod 단위 공유)

  - emptyDir
  - Pod 이 생성되고 삭제될 때, 같이 생성되고 삭제되는 임시 디스크
    - 생성 당시에는 아무 것도 없는 빈 상태
  - 물리 디스크(노드), 메모리에 저장



- 로컬 디스크(노드 단위 공유)

  - hostPath
  - emptyDir 와 같은 컨셉이지만, 공유 범위가 노드라는 점만 다름



- 네트워크 디스크

  - gitRepo (지금은 deprecated 라고 한다.)
    - 생성시에 지정된 git repo 를 clone 한 후, 디스크 생성
    - emptyDir -> git clone 이라보면 됨
  - 그 외 클라우드 서비스 별로 더 있음.



### 2) Ingress

Ingress 는 api 게이트 웨이, 즉 **url 기반 라우팅 역할**을 한다.
Service 앞에 붙는다.

예를 들어, `/user` 로 들어오는 트래픽은 `service A` 에,
`/products` 로 들어오는 트래픽은 `service B` 로 라우팅 시켜준다.

Ingress 를 Service 앞에 달아두면, Service 는 `NodePort` 타입으로 선언되어야 한다.



## HealthCheck

HealthCheck 는 **각 컨테이너 상태를 주기적으로 문제가 있는지 체크**하는 기능이다.
문제가 있으면 재시작하거나, 서비스에서 제외한다.

### 1) 방법

- Liveness probe

  - 응답이 없으면 컨테이나 자동 재시작
  - kubelet 을 통해서 재시작함.

- Readiness probe

  - 서비스가 일시적으로 작동 불가인 경우 서비스에서 제외

  

### 2) 방식

- Commnad probe

  - `cat /tmp/healthy` 실행. 성공 시 0 리턴 -> 정상 판단

- HTTP probe

  - GET 요청 -> status 200 시 정상 판단

- TCP probe

  - TCP 연결 시도 -> 연결 성공하면 정상 판단

