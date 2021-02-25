# 네임스페이스(namespace)

쿠버네티스 클러스터 안의 가상 클러스터포드, 레플리카셋, 디플로이먼트, 서비스 등과 같은 쿠버네티스 리소스들이 묶여 있는 하나의 가상 공간 또는 그룹리소스를 논리적으로 구하기 위한 오브젝트

예를들어 인사, 개발, 마케팅 등 각각의 부서를 생각하면 조금 이해가 더 편할 것 같습니다.



### 네임스페이스 생성, 조회

```sh
[vagrant@master ~]$ kubectl get namespaces	<- `ns`로 단축어 가능
NAME              STATUS   AGE
default           Active   7d21h	
kube-node-lease   Active   7d21h
kube-public       Active   7d21h
kube-system       Active   7d21h
```



### 특정 네임스페이스에 생성된 오브젝트를 확인

```sh
[vagrant@master ~]$ kubectl get pods --namespace default
NAME                                   READY   STATUS    RESTARTS   AGE
hostname-deployment-6cd58767b4-5mtjb   1/1     Running   1          17h
hostname-deployment-6cd58767b4-8c57s   1/1     Running   1          17h
hostname-deployment-6cd58767b4-s9rrs   1/1     Running   1          17h
```

- 네임스페이스를 지정하지 않으면 기본적으로 default 네임스페이스를 사용합니다.



## 네임스페이스 생성, 확인

### 1. YAML 작성

- **`[vagrant@master ~]$ vi production-namespace.yaml`**

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: production
```

- `$ kubectl create namespace [NAMESPACE_NAME]` - `create`를 통해서 간편하게 만들 수 있다.



### 2. 네임스페이스 생성 및 확인

```sh
[vagrant@master ~]$ kubectl apply -f production-namespace.yaml
namespace/production created
```

```sh
[vagrant@master ~]$ kubectl get ns
NAME              STATUS   AGE
default           Active   7d22h
kube-node-lease   Active   7d22h
kube-public       Active   7d22h
kube-system       Active   7d22h
production        Active   39s
```



### 3. 특정 네임스페이스에 리소스를 생성하는 방법

- **`[vagrant@master ~]$ vi hostname-deploy-svc-ns.yaml`**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hostname-deployment-ns
  namespace: production
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
        image: alicek106/rr-test:echo-hostname
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: hostname-svc-clusterip-ns
  namespace: production
spec:
  ports:
    - name: web-port
      port: 8080
      targetPort: 80
  selector:
    app: webserver
  type: ClusterIP
```

```sh
[vagrant@master ~]$ kubectl apply -f hostname-deploy-svc-ns.yaml
```



- 확인

```sh
[vagrant@master ~]$ kubectl get po,svc --namespace production
NAME                                          READY   STATUS    RESTARTS   AGE
pod/hostname-deployment-ns-6cd58767b4-94l9p   1/1     Running   0          48s
pod/hostname-deployment-ns-6cd58767b4-g5zmr   1/1     Running   0          48s
pod/hostname-deployment-ns-6cd58767b4-m7t9l   1/1     Running   0          48s

NAME                               TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/hostname-svc-clusterip-ns  ClusterIP  10.109.85.128  <none>        8080/TCP   48s
```



### 5. 동일 네임스페이스 내의 서비스에 접근

```sh
[vagrant@master ~]$ kubectl get svc --namespace production
NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
hostname-svc-clusterip-ns   ClusterIP   10.109.85.128   <none>        8080/TCP   4m3s

[vagrant@master ~]$ kubectl run -it --rm debug --image=alicek106/ubuntu:curl --restart=Never --namespace=production -- bash
```

```sh
root@debug:/# curl 10.109.85.128:8080
```

- 내부 IP로 접근할 수도 있지만 같은 네임스페이스끼리는  이름으로도 접근이 가능하다.

```sh
root@debug:/# curl hostname-svc-clusterip-ns:8080
```





#### 네임스페이스에 속하는 오브젝트 종류

- `$ kubectl api-resources --namespaced=true`



#### 네임스페이스에 속하지 않는 오브젝트 종류 → 클러스터 전반에 사용되는 경우가 많음

- `$ kubectl api-resources --namespaced=false`