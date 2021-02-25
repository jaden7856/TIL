# 시크릿(secret)

SSH 키, 비밀번호 등과 같이 **민감한 정보를 저장**하는 용도로 사용



### 타입

- Opaque 타입
  - 기본 타입
  - 사용자가 정의하는 데이터를 저장하는 일반적인 목적의 시크릿
- 비공개 레지스트리(private registry)에 접근할 때 사용하는 인증 설정 시크릿



### 시크릿 생성 방법

**--from-literal**

- `generic` :  Opaque 타입의 시크릿 생성을 명시

```sh
[vagrant@master ~]$ kubectl create secret generic my-password --from-literal password=p@ssw0rd
secret/my-password created
```



**--from-file**

```sh
[vagrant@master ~]$ echo mypassword > pw1 && echo yourpassword > pw2
[vagrant@master ~]$ ls pw*
pw1  pw2

[vagrant@master ~]$ kubectl create secret generic our-password --from-file pw1 --from-file pw2
secret/our-password created
```



- 확인

```sh
[vagrant@master ~]$ kubectl get secrets
NAME                  TYPE                                  DATA   AGE
default-token-mmbs5   kubernetes.io/service-account-token   3      8d
my-password           Opaque                                1      9m35s
our-password          Opaque                                2      8s
```





### 시크릿 내용 확인

```sh
[vagrant@master ~]$ kubectl get secrets my-password -o yaml
apiVersion: v1
data:
  password: cEBzc3cwcmQ=			⇐ BASE64로 인코딩
kind: Secret
metadata:
  creationTimestamp: "2021-02-24T07:05:44Z"
  name: my-password
  namespace: default
  resourceVersion: "200708"
  selfLink: /api/v1/namespaces/default/secrets/my-password
  uid: 81d33e98-db84-4f5d-9998-55077e2a1507
type: Opaque
```

- **`BASE64`** : 아스키문자중에서 가시화할 수 있는 데이터 문자입니다.
  - 2진 데이터를 문자 데이터로 바꾸기 위해서 쓴다.
  - 문자가 눈에 보이기때문에 관리하기 편하다.



#### <u>BASE64 인코딩 및 디코딩</u>

- 인코딩

```sh
[vagrant@master ~]$ echo p@ssw0rd | base64
cEBzc3cwcmQK
```

- 디코딩

```sh
[vagrant@master ~]$ echo cEBzc3cwcmQ= | base64 -d
p@ssw0rd
```





### 시크릿에 저장된 값을 포드의 환경변수로 참조

- **시크릿에 저장된 모든 값을 가져오기**

- **`[vagrant@master ~]$ vi env-from-secret.yaml`**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: secret-env-example
spec:
  containers:
    - name: my-container
      image: busybox
      args: ["tail", "-f", "/dev/null"]
      envFrom:
        - secretRef:
            name: my-password
```

- `my-password` 시크릿에 저장된 모든 키-값을 환경변수로 설정



#### 시크릿에 저장된 일부 값을 가져오기

- **`[vagrant@master ~]$ vi selective-env-from-secret.yaml`**

```yaml
piVersion: v1
kind: Pod
metadata:
  name: selective-secret-env-example
spec:
  containers:
    - name: my-container
      image: busybox
      args: ["tail", "-f", "/dev/null"]
      env: 
        - name: YOUR_PASSWORD
          valueFrom:
            secretKeyRef:
              name: out-password
              key: pw2
```



#### 시크릿에 저장된 값 전체를 포드에 볼륨 마운트

- **`[vagrant@master ~]$ vi volume-mount-secret.yaml`**

```sh
apiVersion: v1
kind: Pod
metadata:
  name: secret-volume-example
spec:
  containers:
    - name: my-container
      image: busybox
      args: ["tail", "-f", "/dev/null"]
      volumeMounts:
        - name: secret-volume
          mountPath: /etc/secret
  volumes:
    - name: secret-volume
      secret:
        secretName: our-password
```



#### 시크릿에 저장된 값 일부를 포드에 볼륨 마운트

- **`[vagrant@master ~]$ vi selective-volume-mount-secret.yaml`**

```sh
apiVersion: v1
kind: Pod
metadata:
  name: selective-secret-volume-example
spec:
  containers:
    - name: my-container
      image: busybox
      args: ["tail", "-f", "/dev/null"]
      volumeMounts:
        - name: secret-volume
          mountPath: /etc/secret
  volumes:
    - name: secret-volume
      secret:
        secretName: our-password
        items:
        - key: pw1
          path: password1
```



#### 포드 생성 후 시크릿 확인

```sh
[vagrant@master ~]$ kubectl apply -f selective-volume-mount-secret.yaml
pod/selective-secret-volume-example created

[vagrant@master ~]$ kubectl exec selective-secret-volume-example -- ls -al /etc/secret
total 0
drwxrwxrwt    3 root     root           100 Feb 24 08:05 .
drwxr-xr-x    1 root     root            20 Feb 24 08:05 ..
drwxr-xr-x    2 root     root            60 Feb 24 08:05 ..2021_02_24_08_05_38.422941904
lrwxrwxrwx    1 root     root            31 Feb 24 08:05 ..data -> ..2021_02_24_08_05_38.422941904
lrwxrwxrwx    1 root     root            16 Feb 24 08:05 password1 -> ..data/password1
```

```sh
[vagrant@master ~]$ kubectl exec selective-secret-volume-example -- cat /etc/secret/password1
mypassword
```

- 시크릿을 포드의 환경변수나 볼륨파일로 가져오면 `BASE64`로 디코딩된 값을 사용





# nginx의 인증정보가 담긴 파일을 시크릿으로 관리

### **접근통제 3단계**

- 식별(identification)
- 인증(authentication)
  - TYPE1 : 알고 있는 정보 (지식기반) - 패스워드
  - TYPE2 : 가고 있는 정보 (소유기반) - 주민등록증, OTP, 인증서, 스마트폰, …
  - TYPE3 : 특징 - 홍채, 지문, 성문, 정맥, … , 필기체 서명(싸인)
  - 2가지 이상을 혼용 : 2-factor 인증, multi-factor 인증(=다중인증방식)
  - 멀티 디바이스 인증 = 멀티 채널 인증
- 인가(authorization)



### HTTP 기본인증(HTTP Basic Authentication)

![img](04-6_Kubernetes_Object_Secret.assets/IM4LJq9baXyrqHAQmGxLSX9Sd2CxJ29xuQwq8_10JJ2ZRlJMIMceRJKcjfNrCJBLboBnl8ZXB53RJhi3_w6wdtUPNNKE5ynURllx5wIfdYxVJkLjMYpgNKpzj0dQQT7EIHeZg5pZ)



##### 1. `openssl` 모듈을 이용해서 사용명과 암호화한 패스워드를 BASE64로 인코딩

```sh
[vagrant@master ~]$ echo "your_name:$(openssl passwd -quiet -crypt your_password)" | base64
eW91cl9uYW1lOjVOU1h3NTFsd1pmOHcK
```

- `"your_name:암호화된패스워드"` : 문자열을 생성
- `your_password` : 문자열을 암호화



##### 2. 시크릿을 생성하는 YAML 파일을 작성

- `[vagrant@master ~]$ vi nginx-secret.yaml`
- `your_password`의 암호화된 코드를 `.htpasswd`에 입력

```sh
apiVersion: v1
kind: Secret
metadata:
  name: nginx-secret
type: Opaque
data:
  .htpasswd: eW91cl9uYW1lOjVOU1h3NTFsd1pmOHcK   ⇐ basic authentication에 사용할 사용자 파일
```



- `create`명령어를 사용하여 바로 생성

```sh
[vagrant@master ~]$ kubectl create secret generic nginx-secret --from-literal .htpasswd=eW91cl9uYW1lOjVOU1h3NTFsd1pmOHcK --dry-run -o yaml > test.yaml
```

```sh
[vagrant@master ~]$ cat test.yaml
apiVersion: v1
data:
  .htpasswd: ZVc5MWNsOXVZVzFsT2pWT1UxaDNOVEZzZDFwbU9IY0s=
kind: Secret
metadata:
  creationTimestamp: null
  name: nginx-secret
```



##### 3. 시크릿 생성 및 확인

```sh
[vagrant@master ~]$ kubectl apply -f nginx-secrect.yaml
secret/nginx-secret created

[vagrant@master ~]$ kubectl get secrets
NAME                  TYPE                                  DATA   AGE
default-token-mmbs5   kubernetes.io/service-account-token   3      8d
my-password           Opaque                                1      136m
nginx-secret          Opaque                                1      13s
our-password          Opaque                                2      127m
```



##### 4. 시크릿을 활용해 기본인증을 적용하는 nginx 파드를 구성

- `[vagrant@master ~]$ vi basic-auth.yaml`

```sh
apiVersion: v1
kind: Service
metadata:
  name: basic-auth
spec:
  type: NodePort
  selector: 
    app: basic-auth
  ports:
    - protocol: TCP
      port: 80
      targetPort: http
      nodePort: 30060                ⇐ 클러스터 외부에서 해당 포트로 서비스를 이용
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: basic-auth
  labels:
    app: basic-auth
spec:
  replicas: 1
  selector:
    matchLabels:
      app: basic-auth
  template:
    metadata:
      labels:
        app: basic-auth
    spec:
      containers:
        - name: nginx
          image: "gihyodocker/nginx:latest"      
          imagePullPolicy: Always
          ports:
          - name: http
            containerPort: 80
          env:
          - name: BACKEND_HOST
            value: "localhost:8080"
          - name: BASIC_AUTH_FILE
            value: "/etc/nginx/secret/.htpasswd" 	⇐ (1)
          volumeMounts:
            - mountPath: /etc/nginx/secret          ⇐ (2)
              name: nginx-secret                 	⇐ 볼륨 이름
              readOnly: true
        - name: echo
          image: "gihyodocker/echo:latest"
          imagePullPolicy: Always
          ports: 
          - containerPort: 8080
          env:
          - name: HTTP_PORT
            value: "8080"
      volumes:
      - name: nginx-secret                        ⇐ 볼륨 이름
        secret:
          secretName: nginx-secret                ⇐ 시크릿 이름

```

- (1) : 기본인증에 사용하는 인증정보가 담긴 파일

- (2) : 볼륨과 연결된 디렉터리 아래에 시크릿 키 이름의 파일이 생성 (1)



##### 5. 생성 및 확인

```sh
[vagrant@master ~]$ kubectl apply -f basic-auth.yaml
service/basic-auth created
deployment.apps/basic-auth created

[vagrant@master ~]$ kubectl get pods,deployments,services
NAME                              READY   STATUS    RESTARTS   AGE
pod/basic-auth-745c944d45-mq747   2/2     Running   0          5m7s

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.extensions/basic-auth   1/1     1            1           5m7s

NAME                     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
service/basic-auth       NodePort    10.100.206.209   <none>        80:30060/TCP     5m7s

```



##### 6. 인증 정보 없이 호출 시 401 오류를 반환

```sh
[vagrant@master ~]$ curl -i http://127.0.0.1:30060
HTTP/1.1 401 Unauthorized		⇐ 응답(response) 시작 ← 인증 정보가 없는 경우 401 오류를 반환
Server: nginx/1.13.12					⇐ 응답 헤더 시작
Date: Wed, 24 Feb 2021 10:11:03 GMT
Content-Type: text/html; charset=utf-8
Content-Length: 196
Connection: keep-alive
WWW-Authenticate: Basic realm="Restricted"		⇐ 기본인증 방식으로 인증정보를 요청
												⇐ 응답 헤더 끝
<html>											⇐ 응답 본문 → 사용자 브라우저에 출력되는 내용
<head><title>401 Authorization Required</title></head>
<body bgcolor="white">
<center><h1>401 Authorization Required</h1></center>
<hr><center>nginx/1.13.12</center>
</body>
</html>
```



##### 7. 인증 정보와 함께 호출 시 정상 처리

```sh
[vagrant@master ~]$ curl -i --user your_name:your_password http://127.0.0.1:30060

HTTP/1.1 200 OK					⇐ 정상 처리
Server: nginx/1.13.12
Date: Wed, 24 Feb 2021 10:16:44 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 14
Connection: keep-alive

Hello Docker!!
```

- `--user your_name:사용자 비밀번호` : 인증 정보