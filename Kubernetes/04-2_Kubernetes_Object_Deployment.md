# 디플로이먼트(Deployment)

레플리카셋, 포드의 **배포**, **업데이트** 등을 관리합니다.



## 디플로이먼트 생성, 삭제

### 1. 디플로이먼트 정의, 생성, 확인

- **`[vagrant@master ~]$ vi deployment-nginx.yaml`**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-nginx
  template:
    metadata:
      name: my-nginx-pod
      labels:
        app: my-nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.10
          ports:
          - containerPort: 80
```

```sh
[vagrant@master ~]$ kubectl apply -f deployment-nginx.yaml
deployment.apps/my-nginx-deployment created
```

```sh
[vagrant@master ~]$ kubectl get deployment,replicaset,pod  <-- 한번에 불러올 수도 있습니다.
NAME                                        READY   UP-TO-DATE   AVAILABLE   AGE
deployment.extensions/my-nginx-deployment   0/3     3            0           19s

NAME                                                 DESIRED   CURRENT   READY   AGE
replicaset.extensions/my-nginx-deployment-9b5988dd   3         3         0       19s

NAME                                     READY   STATUS    RESTARTS   AGE
pod/my-nginx-deployment-9b5988dd-667kq   1/1     Running   0          19s
pod/my-nginx-deployment-9b5988dd-nxg4p   1/1     Running   0          19s
pod/my-nginx-deployment-9b5988dd-v7x67   1/1     Running   0          19s
```



### 2. 디플로이먼트 삭제 → 레플리카셋, 포드도 함께 삭제되는 것을 확인

```sh
[vagrant@master ~]$ kubectl delete deployment my-nginx-deployment
deployment.extensions "my-nginx-deployment" deleted
```



## 디플로이먼트를 사용하는 이유

어플리케이션의 **업데이트와 배포를 편하게 만들기 위해서** 사용합니다.



### 1. `--record` 옵션을 추가해 디플로이먼트를 생성

```sh
[vagrant@master ~]$ kubectl apply -f deployment-nginx.yaml --record
deployment.apps/my-nginx-deployment created
```

```sh
[vagrant@master ~]$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
my-nginx-deployment-9b5988dd-24cp8   1/1     Running   0          12s
my-nginx-deployment-9b5988dd-8g557   1/1     Running   0          12s
my-nginx-deployment-9b5988dd-jh4d9   1/1     Running   0          12s
```



### 2. `kubectl set image` 명령으로 포드의 이미지를 변경

```sh
[vagrant@master ~]$ kubectl set image deployment my-nginx-deployment nginx=nginx:1.11 --record
deployment.extensions/my-nginx-deployment image updated
```

```sh
[vagrant@master ~]$ kubectl get pods
NAME                                  READY   STATUS              RESTARTS   AGE
my-nginx-deployment-9b5988dd-24cp8    0/1     Terminating         0          3m32s
my-nginx-deployment-9b5988dd-8g557    1/1     Running             0          3m32s
my-nginx-deployment-9b5988dd-jh4d9    1/1     Running             0          3m32s
my-nginx-deployment-d4659856c-qxflc   1/1     Running             0          16s
my-nginx-deployment-d4659856c-rjt94   0/1     ContainerCreating   0          5s	
```

- 첫번째 기존의 1.0 버전은 삭제되고있는중이고 맨 밑의 새로운 1.11 버전은 생성되는 중입니다.



```sh
[vagrant@master ~]$ kubectl get replicaset
NAME                            DESIRED   CURRENT   READY   AGE
my-nginx-deployment-9b5988dd    0         0         0       9m46s
my-nginx-deployment-d4659856c   3         3         3       6m29s
```

- 첫번째는 처음에 생성했던 레플리카셋
- 두번째는 새롭게 생성된 레플리카셋입니다.



### 3. 리버전 정보 확인

```sh
[vagrant@master ~]$ kubectl rollout history deployment my-nginx-deployment
deployment.extensions/my-nginx-deployment
REVISION  CHANGE-CAUSE
1         kubectl apply --filename=deployment-nginx.yaml --record=true
2         kubectl set image deployment my-nginx-deployment nginx=nginx:1.11 --record=true
```

- 지금까지 입력했던 명령어의 history를 알 수 있습니다.



### 4. 이전 버전의 레플리카셋으로 롤백

```sh
[vagrant@master ~]$ kubectl rollout undo deployment my-nginx-deployment --to-revision=1
deployment.extensions/my-nginx-deployment rolled back
```

- 특정 버전으로 이동하고 싶을때 `--to-revision`을 사용합니다.



```sh
[vagrant@master ~]$ kubectl get pods
NAME                                  READY   STATUS        RESTARTS   AGE
my-nginx-deployment-9b5988dd-5kgfz    1/1     Running       0          7s
my-nginx-deployment-9b5988dd-bcz5x    1/1     Running       0          10s
my-nginx-deployment-9b5988dd-lrzbj    1/1     Running       0          13s
my-nginx-deployment-d4659856c-qxflc   0/1     Terminating   0          14m
my-nginx-deployment-d4659856c-rjt94   0/1     Terminating   0          14m
```

```sh
[vagrant@master ~]$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
my-nginx-deployment-9b5988dd-5kgfz   1/1     Running   0          17s
my-nginx-deployment-9b5988dd-bcz5x   1/1     Running   0          20s
my-nginx-deployment-9b5988dd-lrzbj   1/1     Running   0          23s
```

- 위의 결과들을 보면서 알수있는건  차례차례 지워지면서 생성되는 모습을 볼 수 있습니다.



```sh
[vagrant@master ~]$ kubectl get replicasets
NAME                            DESIRED   CURRENT   READY   AGE
my-nginx-deployment-9b5988dd    3         3         3       19m
my-nginx-deployment-d4659856c   0         0         0       16m
```

- 새롭게 생성된 레플리카셋에서 예전의 레플리카셋으로 돌아갔습니다.



```sh
[vagrant@master ~]$ kubectl rollout history deployment my-nginx-deployment
deployment.extensions/my-nginx-deployment
REVISION  CHANGE-CAUSE
2         kubectl set image deployment my-nginx-deployment nginx=nginx:1.11 --record=true
3         kubectl apply --filename=deployment-nginx.yaml --record=true
```

- history도 순서가 바뀐것을 볼 수 있습니다.



### 5. 모든 리소소를 삭제

```sh
[vagrant@master ~]$ kubectl delete deployment,replicaset,pod --all
```