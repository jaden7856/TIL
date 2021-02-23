# 서비스(service)

포드를 연결하고 외부에 노출합니다.
YAML 파일에 containerPort 항목을 정의했다고 해서 해당 포트가 바로 외부에 노출되는 것은 아니고 해당 포트로 사용자가 접근하거나, 다른 디플로이먼트의 포드들이 내부적으로 접근하려면 서비스(service) 객체가 필요합니다.



### 서비스의 기능

- 여러 개의 포드에 쉽게 접근할 수 있도록 고유한 도메인 이름을 부여
- 여러 개의 포드에 접근할 때, 요청을 분산하는 로드 밸런서 기능을 수행
- 클라우드 플랫폼의 로드 벨런서, 클러스터 노드의 포트 등을 통해 포드를 외부에 노출



### 서비스의 종류(type)

- **ClusterIP 타입**
  - 쿠버네티스 내부에서만 포드들에 접근할 때 사용
  - 외부로 포드를 노출하지 않기 때문에 쿠버네티스 클러스터 내부에서만 사용되는 포드에 적합
- **NodePort 타입**
  - 포드에 접근할 수 있는 포트를 클러스터의 모든 노드에 동일하게 개방
  - 외부에서 포드에 접근할 수 있는 서비스 타입
  - 접근할 수 있는 포트는 랜덤으로 정해지지만, 특정 포트로 접근하도록 설정할 수 있음
- **LoadBalancer 타입**
  - 클라우드 플랫폼에서 제공하는 로드 벨러서를 동적으로 프로비저닝해 포드에 연결
  - `NodePort` 타입과 마찬가지로 외부에서 포드에 접근할 수 있는 서비스 타입
  - 일반적으로 AWS, GCP 과 같은 클라우드 플랫폼 환경에서 사용



## 디플로이먼트를 생성

### 1. 디플로이먼트 정의, 생성, 확인

- **`[vagrant@master ~]$ vi deployment-hostname.yaml`**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hostname-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: webserver
  template:
    metadata:
      name: my-webserver
      labels:
        app: webserver
    spec:
      containers:
        - name: my-webserver
          image: alicek106/rr-test:echo-hostname  # 포드의 호스트 이름을 반환하는 웹 서버 이미지
          ports:
          - containerPort: 80

```

```sh
[vagrant@master ~]$ kubectl apply -f deployment-hostname.yaml
deployment.apps/hostname-deployment created
```



```sh
[vagrant@master ~]$ kubectl get pods -o wide
NAME                                  READY STATUS   RESTARTS  AGE  IP              NODE
hostname-deployment-6cd58767b4-5mtjb  1/1   Running  0         49s  192.168.166.184 node1 
hostname-deployment-6cd58767b4-8c57s  1/1   Running  0         49s  192.168.104.52  node2 
hostname-deployment-6cd58767b4-s9rrs  1/1   Running  0         49s  192.168.166.183 node1
```



### 2. 클러스터 노드 중 하나에 접속해서 `curl`을 이용해 포드에 접근

```sh
[vagrant@master ~]$ kubectl run -i --tty --rm debug --image=alicek106/ubuntu:curl --restart=Never curl 192.168.166.184 | grep Hello
 
 	<p>Hello,  hostname-deployment-6cd58767b4-5mtjb</p>     </blockquote>
```



## ClusterIP 타입의 서비스 - 쿠버네티스 클러스터 내부에서만 포드에 접근

### 1. `hostname-svc-clusterip.yaml` 파일을 생성

- **`[vagrant@master ~]$ vi hostname-svc-clusterip.yaml`**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: hostname-svc-clusterip
spec:
  ports:
  - name: web-port
    port: 8080               ⇐ 서비스의 IP에 접근할 때 사용할 포트
    targetPort: 80           ⇐ selector 항목에서 정의한 라벨의 포드의 내부에서 사용하고 있는 포트
  selector:                  ⇐ 접근 허용할 포드의 라벨을 정의
    app: webserver
  type: ClusterIP            ⇐ 서비스 타입
```



### 2. 서비스 생성 및 확인

```sh
[vagrant@master ~]$ kubectl apply -f hostname-svc-clusterip.yaml
service/hostname-svc-clusterip created

[vagrant@master ~]$ kubectl get services
NAME                     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
hostname-svc-clusterip   ClusterIP   10.107.79.119    <none>        8080/TCP         7s
kubernetes               ClusterIP   10.96.0.1        <none>        443/TCP          7d4h 
```

- `kuberentes` service는 쿠버네티스 API에 접근하기 위한 서비스입니다.



### 3. 임시 포드를 생성해서 서비스로 요청

```sh
[vagrant@master ~]$ kubectl run -it --rm debug --image=alicek106/ubuntu:curl --restart=Never -- bash
If you don't see a command prompt, try pressing enter.

root@debug:/# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
4: eth0@if42: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1440 qdisc noqueue state UP group default
    link/ether 3e:00:8f:da:36:b4 brd ff:ff:ff:ff:ff:ff
    inet 192.168.104.53/32 scope global eth0
       valid_lft forever preferred_lft forever
```

```sh
root@debug:/# curl 10.107.79.119:8080 --silent | grep Hello
        <p>Hello,  hostname-deployment-6cd58767b4-5mtjb</p>     </blockquote>

root@debug:/# curl 10.107.79.119:8080 --silent | grep Hello			   
        <p>Hello,  hostname-deployment-6cd58767b4-s9rrs</p>     </blockquote>

root@debug:/# curl 10.107.79.119:8080 --silent | grep Hello
        <p>Hello,  hostname-deployment-6cd58767b4-8c57s</p>     </blockquote>
```

- 계속 같은 IP의 명령어를 실행하여도 다른 포드를 조회합니다. 이것이 **로드밸런싱**입니다.



### 4. 서비스 이름으로 접근

```sh
root@debug:/# curl hostname-svc-clusterip:8080 --silent | grep Hello	
        <p>Hello,  hostname-deployment-6cd58767b4-5mtjb</p>     </blockquote>

root@debug:/# curl hostname-svc-clusterip:8080 --silent | grep Hello
        <p>Hello,  hostname-deployment-6cd58767b4-8c57s</p>     </blockquote>
```

- 쿠버네티스는 어플리케이션이 서비스나 포드를 쉽게 찾을 수 있도록 내부 DNS를 구동



### 5. 서비스 삭제

```sh
[vagrant@master ~]$ kubectl delete service hostname-svc-clusterip
service "hostname-svc-clusterip" deleted
```

