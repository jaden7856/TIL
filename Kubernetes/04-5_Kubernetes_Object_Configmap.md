# 컨피그맵(configmap)

컨피그맵은 포드에 설정값을 전달하고 키-값 쌍으로 **기밀이 아닌** 데이터를 저장하는 데 사용하는 API 오브젝트



### 컨피그맵 생성

- `kubectl create configmap` 명령어로 생성

```sh
[vagrant@master ~]$ kubectl create configmap log-level-configmap --from-literal LOG_LEVEL=DEBUG

[vagrant@master ~]$ kubectl create configmap start-k8s --from-literal k8s=kubernetes --from-literal container=docker
```

- `k8s=kubernetes`(key=value) 형식의 정보를 저장



## 방법1) 포드에서 컨피그맵을 사용

### envFrom

- 컨피그맵에 정의된 **모든 키=값**쌍을 가져와서 환경변수로 설정



##### 1. YAML 파일을 작성

- **`[vagrant@master ~]$ vi all-env-from-configmap.yaml`**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: container-env-example
spec:
  containers:
    - name: my-container
      image: busybox
      args: ['tail', '-f', '/dev/null']
      envFrom:
      - configMapRef:
          name: log-level-configmap		<- configmap 이름
      - configMapRef:
          name: start-k8s

```



##### 2. 생성 및 컨테이너 환경변수를 확인

```sh
[vagrant@master ~]$ kubectl apply -f all-env-from-configmap.yaml
pod/container-env-example created
```

```sh
[vagrant@master ~]$ kubectl exec container-env-example -- env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=container-env-example
LOG_LEVEL=DEBUG						<- configmap : log-level-configmap 
container=docker					<- configmap : start-k8s
k8s=kubernetes						<- configmap : start-k8s
NGINX_TEST_PORT_80_TCP_PORT=80
```





### valueFrom, configMapKeyRef

- 컨피그맵에 존재하는 키=값 쌍중에서 **원하는 데이터만** 선택해서 환경변수로 설정



##### 1. YAML 파일을 작성

- **`[vagrant@master ~]$ vi selective-env-from-configmap.yaml`**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: container-env-example
spec:
  containers:
    - name: my-container
      image: busybox
      args: ['tail', '-f', '/dev/null']
      env:
      - name: ENV_KEYNAME_1            		⇐ 새롭게 정의될 환경변수 이름
        valueFrom:
          configMapKeyRef:
            key: LOG_LEVEL             		⇐ 해당 컨피그맵에 정의된 변수의 이름(키)
            name: log-level-configmap  		⇐ 참조할 컨피그맵 이름 
      - name: ENV_KEYNAME_2
        valueFrom:
          configMapKeyRef:
            key: k8s
            name: start-k8s

```



##### 2. 포드 생성 및 환경변수 확인

```sh
[vagrant@master ~]$ kubectl apply -f selective-env-from-configmap.yaml
pod/container-env-example created

[vagrant@master ~]$ kubectl exec container-env-example -- env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=container-env-example
ENV_KEYNAME_1=DEBUG
ENV_KEYNAME_2=kubernetes
KUBERNETES_PORT_443_TCP_PORT=443
```





## 방법2) 컨피그맵의 값을 포드 내부의 파일로 마운트해 사용

### volumeMounts

- 모든 키-값 쌍 데이터를 포드에 마운트



##### 1. YAML 파일을 생성

- **`[vagrant@master ~]$ vi volume-mount-configmap.yaml`**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: configmap-volume-pod
spec:
  containers:
    - name: my-container
      image: busybox
      args: ["tail", "-f", "/dev/null"]
      volumeMounts:
        - name: configmap-volume		<- volume 파일을 생성
          mountPath: /etc/config		<- 생성할 경로
  volumes:
    - name: configmap-volume			<- volume 연결할 이름
      configMap:
        name: start-k8s
```



##### 2. 포드를 생성

```sh
[vagrant@master ~]$ kubectl apply -f volume-mount-configmap.yaml
pod/configmap-volume-pod created
```



##### 3. 포드 내부의 /etc/config 디렉터리를 조회

```sh
[vagrant@master ~]$ kubectl exec configmap-volume-pod -- ls -al /etc/config
total 0
drwxrwxrwx    3 root     root   87 Feb 24 05:46 .
drwxr-xr-x    1 root     root   20 Feb 24 05:46 ..
drwxr-xr-x    2 root     root   34 Feb 24 05:46 ..2021_02_24_05_46_50.402227761
lrwxrwxrwx    1 root     root   31 Feb 24 05:46 ..data -> ..2021_02_24_05_46_50.402227761
lrwxrwxrwx    1 root     root   16 Feb 24 05:46 container -> ..data/container	
lrwxrwxrwx    1 root     root   10 Feb 24 05:46 k8s -> ..data/k8s
```

- 밑에 부분 `container`와 `k8s` 부분 컨피그맵의 키이름의 파일이 생성



```sh
[vagrant@master ~]$ kubectl exec configmap-volume-pod -- cat /etc/config/container
docker		<- 파일 내용이 컨피그맵의 값과 동일
```





### 원하는 키-값 쌍의 데이터만 선택해서 포드에 마운트

##### 1. YAML 생성

- **`[vagrant@master ~]$ vi selective-volume-configmap.yaml`**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: configmap-volume-pod
spec:
  containers:
    - name: my-container
      image: busybox
      args: ["tail", "-f", "/dev/null"]
      volumeMounts:
        - name: configmap-volume
          mountPath: /etc/config
  volumes:
    - name: configmap-volume
      configMap:
        name: start-k8s                   ⇐ 컨피그맵의 이름
        items:
        - key: k8s                        ⇐ 컨피그맵에서 가져올 값의 키
          path: k8s_fullname              ⇐ 컨피그맵의 키에 해당하는 값을 저장할 파일 이름
```



##### 2. 포드 생성 후 마운트 디렉터리 및 파일을 확인

```sh
[vagrant@master ~]$ kubectl apply -f selective-volume-configmap.yaml
pod/configmap-volume-pod created

[vagrant@master ~]$ kubectl exec configmap-volume-pod -- ls -al /etc/config
total 0
drwxrwxrwx    3 root     root   79 Feb 24 05:59 .
drwxr-xr-x    1 root     root   20 Feb 24 05:59 ..
drwxr-xr-x    2 root     root   26 Feb 24 05:59 ..2021_02_24_05_59_16.332588447
lrwxrwxrwx    1 root     root   31 Feb 24 05:59 ..data -> ..2021_02_24_05_59_16.332588447
lrwxrwxrwx    1 root     root   19 Feb 24 05:59 k8s_fullname -> ..data/k8s_fullname

[vagrant@master ~]$ kubectl exec configmap-volume-pod -- cat /etc/config/k8s_fullname
kubernetes
```





### 파일로부터 컨피그맵을 생성

nginx.conf (nginx 서버의 설정 파일) 또는 mysql.conf (MySQL 설정 파일) 등의 내용 전체를 컨피그맵에 저장할 때 사용



##### 1. 설정 파일을 생성

```sh
[vagrant@master ~]$ echo Hello, World! >> index.html
```



##### 2. index.html 파일의 내용으로 컨피그맵을 생성 및 확인

```sh
[vagrant@master ~]$ kubectl create configmap index-file --from-file index.html
configmap/index-file created

[vagrant@master ~]$ kubectl describe configmap index-file
Name:         index-file
Namespace:    default
Labels:       <none>
Annotations:  <none>
Data
====
index.html:				⇐ 파일 명이 컨피그맵의 키(key)로 사용
----
Hello, World!			⇐ 파일 내용이 컨피그맵의 값(value)로 사용

Events:  <none>
```



##### 3. 키 이름을 직접 지정해서 컨피그맵을 생성

```sh
[vagrant@master ~]$ kubectl create configmap index-file-customkey --from-file myindex=index.html
configmap/index-file-customkey created

[vagrant@master ~]$ kubectl describe configmap index-file-customkey
Name:         index-file-customkey
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
myindex:			⇐ 컨피그맵 생성시 사용한 키
----
Hello, World!

Events:  <none>
```





### 여러 개의 키-값 형식의 내용으로 구성된 설정 파일을 한꺼번에 컨피그맵으로 가져오기

##### 1. 여러 개의 키-값 형식의 내용으로 구성된 설정 파일을 생성

- **`$ vi multiple-keyvalue.env`**

```sh
mykey1=myvalue1
mykey2=myvalue2
mykey3=myvalue3
mykey4=myvalue4
```



- **`$ kubectl create configmap from-env-1 --from-file multiple-keyvalue.env`**

```sh
[vagrant@master ~]$ kubectl get configmap from-env-1 -o yaml
apiVersion: v1
data:
  multiple-keyvalue.env: |+		⇐ 하나의 키
    mykey1=myvalue1			⇐ 하나의 값	
    mykey2=myvalue2
    mykey3=myvalue3
    mykey4=myvalue4

kind: ConfigMap
metadata:
  creationTimestamp: "2021-02-24T06:15:00Z"
  name: from-env-1
  namespace: default
  resourceVersion: "196301"
  selfLink: /api/v1/namespaces/default/configmaps/from-env-1
  uid: 6dcc3748-17de-4af2-af06-e0099c08f90b
```

위의 경우는 값을 여러개 적었지만 하나의 키값 안에 모두 들어있어 실질적인 데이터값은 1개입니다.



- 한번에 여러개의 키/값으로 만들기

```sh
[vagrant@master ~]$ kubectl create configmap from-env-2 --from-env-file multiple-keyvalue.env
configmap/from-env-2 created

[vagrant@master ~]$ kubectl get configmap from-env-2 -o yaml
apiVersion: v1
data:
  mykey1: myvalue1			⇐ 네개의 키/값으로 구성
  mykey2: myvalue2
  mykey3: myvalue3
  mykey4: myvalue4
kind: ConfigMap
metadata:
  creationTimestamp: "2021-02-24T06:17:40Z"
  name: from-env-2
  namespace: default
  resourceVersion: "196534"
  selfLink: /api/v1/namespaces/default/configmaps/from-env-2
  uid: 72086a95-ca18-4a78-96b1-51ff431c4e27
[vagrant@master ~]$ kubectl describe configmap from-env-2
Name:         from-env-2
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
mykey2:
----
myvalue2
mykey3:
----
myvalue3
mykey4:
----
myvalue4
mykey1:
----
myvalue1
Events:  <none>
```