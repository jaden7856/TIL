# Basic Project 

Linux에서 Node.js를 이용한 작은 파일을 생성하여 Docker를 커쳐 Kubernetes를 사용하기 까지의 작은 프로젝트를 하겠습니다.

우선 처음 root계정에서 바로 만들기는 좀 그러니까 파일을 하나 새로 생성해서 작업하겠습니다.

`# mkdir node_project`



### STATUS가 NotReady인 경우 해당 노드에 접속해서 kubelet 실행 상태 확인 후 실행

```yaml
[vagrant@node1 ~]$ systemctl status kubelet
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/usr/lib/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /usr/lib/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: inactive (dead)
     Docs: https://kubernetes.io/docs/

```

- Active 상태가 dead라면 밑의 명령어를 실행

```yaml
$ sudo systemctl restart kubelet
```



### 1. Node 설치

Linux에서 Node를 설치하겠습니다.

- `# yum install epel-release`
- `# yum install -y gcc-c++ make`
- `# curl -sL http://rpm.nodesource.com/setup_12.x | sudo -E bash -`
- `# yum install nodejs`



### 2. Node 생성

- `# vi hello.js` -- js파일을 생성면서 접속하여 코드 입력

```nginx
var http = require('http');
var content = function(req, res) {
	res.end('Hello Kubernetes! on Docker' + '\n');
	res.writeHead(200);
}

var my_server = http.createServer(content);
my_server.listen(8000);

```



### 3. Dockerfile 생성

- `# vi Dockerfile`


`vi` 명령어를 통해 Dockerfile을 생성하고 내부에 접속후 밑의 간단한 Dockerfile 입력

```dockerfile
FROM node:slim
EXPOSE 8000
COPY hello.js .
CMD node hello.js
```



### 4. Docker

```bash
# docker build -t [user_name]/hello .
```

- docker image를 build하면서 Hub사이트에 push하기위해 자기 계정이름을 입력합니다.



```bash
# docker run -d -p 8001:8000 jhg7856/hello
```

- docker container를 생성하면서 portforward를 하겠습니다.



```bash
# docker push jhg7856/hello
```

- 우선 `docker login`을 하여 docker hub사이트 계정 로그인 후에 push를 합니다.



### 5. Kubernetes

```bash
# vi my_hello_pod.yml
```

- `my_hello_pod.yml`을 생성하여 밑의 코드를 입력하겠습니다.

```dockerfile
apiVersion: v1
kind: Pod
metadata:
  name: hello-Pod
  labels:
    app: hello
spec:
  containers:
  - name: hello-container
    image: jhg7856/hello
    ports:
    - containerPort: 8000
```



```bash
# kubectl apply -f my_hello_pod.yml
```

- 모든 설정 파일들은 디렉토리 안에 들어있고 오브젝트를 생성하거나 패치 한다.



### 5-1 Another YML

두개의 컨테이너를 추가한 다른 `.yml`파일도 만들어보자.

- `# vi my_two_container.yml`

```dockerfile
apiVersion: v1
kind: Pod
metadata:
  name: pod-1
spec:
  containers:
  - name: container1
    image: kubetm/p8000
    ports:
    - containerPort: 8000
  - name: container2
    image: kubetm/p8000
    ports:
    - containerPort: 8080
```

**`kubectl apply -f my_hello_pod.yml` **



### 6. Test

모두 정상적으로 완료가 되었다면 pod list를 확인해보고 실행하여 연결을 해보겠습니다.

- `# kubectl get pods`

- `# kubectl exec -it hello-pod /bin/bash`

- `# apt-get update && apt-get install -y curl`



```
curl -X GET http://127.0.0.1:8000
```

`Hello Kubernetes! on Docker`가 출력된다면 정상



### 7. Service

Pod의 경우에 지정되는 Ip가 랜덤하게 지정이 되고 리스타트 때마다 변하기 때문에 고정된 엔드포인트로 호출이 어렵다, 또한 여러 Pod에 같은 애플리케이션을 운용할 경우 이 Pod 간의 로드밸런싱을 지원해줘야 하는데, 서비스가 이러한 역할을 한다.

서비스는 지정된 IP로 생성이 가능하고, 여러 Pod를 묶어서 로드 밸런싱이 가능하며, 고유한 DNS 이름을 가질 수 있다.



- `# vi my_hello_svc.yml` 

```
apiVersion: v1
kind: Service
metadata:
  name: hello-service
spec:
  selector:
    app: hello
  ports:
    - port: 8001
      targetPort: 8000
  type: NodePort
```

**`# kubectl apply -f my_hello_svc.yml`**



정상적으로 created가 된다면 `# kubectl get svc`로 확인을해보자

```
NAME            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello-service   NodePort    10.97.162.181   <none>        8001:30569/TCP   6m29s
```



저는 PORT번호가 `30569`로 번호를 받았습니다. 테스트를 해보아요

- `# curl -X GET http://127.0.0.1:30569`
- `# curl -X GET http://127.0.0.1:8001`

둘다 정상적으로 출력이 된다면 정상!



- `# kubectl get pods -o wide`

위의 명령어를 통해 node가 어디에 연결되어있는지 확인! 저는 `node1`로 연결이 되었습니다.

```
NAME         READY   STATUS              RESTARTS   AGE   IP        NODE    NOMINATED ...
hello-pod    0/1     ContainerCreating   0          10s   <none>    node1   <none>    
```



- `# kubectl get nodes -o wide`

node1의 ip주소를 확인하기 위해 위의 명령어를 실행합니다.

```
NAME     STATUS   ROLES    AGE    VERSION   INTERNAL-IP     EXTERNAL-IP   OS-IMAGE   ... 
master   Ready    master   3d2h   v1.15.5   192.168.56.10   <none>        CentOS Linux 7 
node1    Ready    <none>   3d2h   v1.15.5   192.168.56.11   <none>        CentOS Linux 7 
node2    Ready    <none>   3d2h   v1.15.5   10.0.2.15       <none>        CentOS Linux 7 
```

 

그럼 이제 windows나 macOS에서 정상작동하는지 확인해봅시다

![image-20210219174906309](05_Kubernetes_Project.assets/image-20210219174906309.png)