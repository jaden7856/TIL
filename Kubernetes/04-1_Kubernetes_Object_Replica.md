# 레플리카셋(Replica Set)

정해진 수의 동일한 포드가 항상 실행되도록 관리합니다.

노드 장애 등의 이유로 포드를 사용할 수 없다면 다른 노드에서 포드를 다시 생성



## 레플리카셋 생성 및 삭제

### 1. nginx 컨테이너 두 개를 실행하는 포드를 정의

```sh
apiVersion: v1
kind: Pod
metadata:
  name: my-nginx-pod-a
spec:
  containers:
  - name: my-nginx-container
    image: nginx
    ports:
    - containerPort: 80
      protocol: TCP
---
apiVersion: v1
kind: Pod
metadata:
  name: my-nginx-pod-b
spec:
  containers:
  - name: my-nginx-container
    image: nginx
    ports:
    - containerPort: 80
      protocol: TCP
```

- 동일한 포드를 일일이 정의하는 것은 매우 비효율적인 작업이다. 

- 포드가 삭제되거나, 포드가 위치한 노드에 장애가 발생해서 포드에 접근할 수 없는 경우에는 관리자가 직접 포드를 삭제하고 다시 생성해야는 문제가 있다.



### 2. 레플리카셋을 정의

- **`[vagrant@master ~]$ vi replicaset-nginx.yaml`**

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: replicaset-nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-nginx-pods-label
  template:                         ⇐ 포드 스펙, 포드 템플릿 → 생성할 포드를 명시
    metadata:
      name: my-nginx-pod
      labels:
        app: my-nginx-pods-label
    spec:
      containers:
        - name: my-nginx-container
          image: nginx:latest
          ports:
          - containerPort: 80
            protocol: TCP
```

- `replicas` : 이 명령어가 같은 포드를 몇개만들지 정한다.
- `matchLabels` : `app` 에서 정한 이름이 포드의 이름과 같을때 레플리카셋으로 정의한다.
  - ex) `matchLabels`와 `labels`의 이름이 같으므로 같은 이름 3개의 포드를 생성



### 3. 레플릿카셋을 생성 및 확인

```sh
vagrant@master ~]$ kubectl apply -f replicaset-nginx.yaml
replicaset.apps/replicaset-nginx created
```

```sh
[vagrant@master ~]$ kubectl get pods,replicaset
NAME                         READY   STATUS    RESTARTS   AGE
pod/replicaset-nginx-dkw42   1/1     Running   0          16s
pod/replicaset-nginx-hftsb   1/1     Running   0          16s
pod/replicaset-nginx-srcwq   1/1     Running   0          16s

NAME                                     DESIRED   CURRENT   READY   AGE
replicaset.extensions/replicaset-nginx   3         3         3       17s
```



### 4. 레플리카셋을 삭제

```sh
[vagrant@master ~]$ kubectl delete rs replicaset-nginx
replicaset.extensions "replicaset-nginx" deleted
```



## 레플리카셋 동작원리

### 1. `app: my-nginx-pods-label` 라벨을 가지는 포드를 생성

- **`[vagrant@master ~]$ vi nginx-pod-without-rs.yaml`**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-nginx-pod
  labels:
    app: my-nginx-pods-label
spec:
  containers:
    - name: my-nginx-container
      image: nginx:latest
      ports:
      - containerPort: 80
```

```sh
[vagrant@master ~]$ kubectl apply -f nginx-pod-without-rs.yaml
pod/my-nginx-pod created
```

```sh
[vagrant@master ~]$ kubectl get pods --show-labels			⇐ 라벨을 함께 출력
NAME           READY   STATUS    RESTARTS   AGE     LABELS
my-nginx-pod   1/1     Running   0          14s     app=my-nginx-pods-label
```



### 2. `app: my-nginx-pods-label` 라벨을 가지는 포드를 3개 생성하는 레플리카셋을 정의 후 생성

- **`[vagrant@master ~]$ cat replicaset-nginx.yaml`**

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: replicaset-nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-nginx-pods-label
  template:
    metadata:
      name: my-nginx-pod
      labels:
        app: my-nginx-pods-label
    spec:
      containers:
        - name: my-nginx-container
          image: nginx:latest
          ports:
          - containerPort: 80
            protocol: TCP
```

```sh
[vagrant@master ~]$ kubectl apply -f replicaset-nginx.yaml
replicaset.apps/replicaset-nginx created
```

```sh
[vagrant@master ~]$ kubectl get pods --show-labels
NAME                     READY   STATUS    RESTARTS   AGE     LABELS
my-nginx-pod             1/1     Running   0          6m16s   app=my-nginx-pods-label
replicaset-nginx-d2fw8   1/1     Running   0          11s     app=my-nginx-pods-label
replicaset-nginx-d9np5   1/1     Running   0          11s     app=my-nginx-pods-label
```

- **yaml 파일에서는 3개를 생성한다고 했지만 기존의 같은 라벨을 가지고있던 1개를 포함하여 2개의 포드만 새로 생성되었다.**



### 3. #1에서 수동으로 생성한 포드를 삭제

```sh
[vagrant@master ~]$ kubectl delete pods my-nginx-pod
pod "my-nginx-pod" deleted
```

```sh
[vagrant@master ~]$ kubectl get pods --show-labels
NAME                     READY   STATUS    RESTARTS   AGE     LABELS
replicaset-nginx-d2fw8   1/1     Running   0          4m28s   app=my-nginx-pods-label
replicaset-nginx-d9np5   1/1     Running   0          4m28s   app=my-nginx-pods-label
replicaset-nginx-m4kx2   1/1     Running   0          16s     app=my-nginx-pods-label
```

- **기존의 것을 제거하였더니 레플리카 셋이 새로운 포드를 생성하였다.**



### 4. 레플리카셋이 생성한 포드의 라벨을 변경 후 포드 정보를 조회

```sh
[vagrant@master ~]$ kubectl edit pods replicaset-nginx-d2fw8
```

- 아래 두 부분을 주석 처리 후 저장
  `#labels:`
  `#app: my-nginx-pods-label`



```sh
[vagrant@master ~]$ kubectl get pods --show-labels
NAME                     READY   STATUS    RESTARTS   AGE     LABELS
replicaset-nginx-4ljt7   1/1     Running   0          3m24s   app=my-nginx-pods-label
replicaset-nginx-d2fw8   1/1     Running   0          12m     <none>
replicaset-nginx-d9np5   1/1     Running   0          12m     app=my-nginx-pods-label
replicaset-nginx-m4kx2   1/1     Running   0          8m23s   app=my-nginx-pods-label
```

- 첫번째 부분에 새로운 포드가 새로생성 되었습니다.
- 두번째 포드가 라벨이 없어지면서 관리 대상에서 제외된 모습입니다.
- 관리 대상에서 제외되면 레플리카 셋을 삭제하여도 제외된 포드는 **지워지지 않고 남아있습니다.**