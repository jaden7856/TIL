# MariaDB를 적용시킨 Deployment.yaml 생성

만약 `10_AWS_Kubernetes_kops`에서 작업을 하고 왔다면 `pod`와 `service`는 지우고 하겠습니다.

### 퍼시스턴트 볼륨

*퍼시스턴트볼륨* (`PV`)은 관리자가 프로비저닝하거나 [스토리지 클래스](https://kubernetes.io/ko/docs/concepts/storage/storage-classes/)를 사용하여 동적으로 프로비저닝한 클러스터의 스토리지이다. 노드가 클러스터 리소스인 것처럼 `PV`는 클러스터 리소스이다. `PV`는 Volumes와 같은 볼륨 플러그인이지만, `PV`를 사용하는 개별 파드와는 별개의 라이프사이클을 가진다. 이 API 오브젝트는 NFS, iSCSI 또는 클라우드 공급자별 스토리지 시스템 등 스토리지 구현에 대한 세부 정보를 담아냅니다.

만약 `Pod`가 삭제가된다면 그 내부에 있던 파일이 다 삭제가 되지만 퍼시스턴트 볼륨으로인해 다른 파일에도 저장이 되면서 데이터유실을 막을 수 있습니다.

- **`vi persistentVolume.yaml`**

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv0001
  labels:
    type: local
spec:
  capacity:
    storage: 2Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/data001/pv0001"
```

```cmd
$ kubectl apply -f persistentVolume.yaml
```



### *퍼시스턴트볼륨클레임* 

*퍼시스턴트볼륨클레임* (`PVC`)은 사용자의 스토리지에 대한 요청이다. 파드와 비슷하다. 파드는 노드 리소스를 사용하고 `PVC`는 `PV` 리소스를 사용한다. 파드는 특정 수준의 리소스(CPU 및 메모리)를 요청할 수 있다. 클레임은 특정 크기 및 접근 모드를 요청할 수 있다(예: `ReadWriteOnce`, `ReadOnlyMany` 또는 `ReadWriteMany`로 마운트 할 수 있음)

- **`vi persistentVolumeClaim.yaml`**

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-volumeclaim
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  selector:
    matchLabels:
      type: local
```

```cmd
$ $ kubectl apply -f persistentVolumeClaim.yaml
```



- 조회

```cmd
$ kubectl get pv
NAME     CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                	 ...
pv0001   2Gi        RWO            Retain           Bound    default/my-volumeclaim
```

```cmd
$ kubectl get pvc
NAME             STATUS   VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS   AGE
my-volumeclaim   Bound    pv0001   2Gi        RWO                           96s
```



#### Error

만약 `pvc` STATUS가 계속 `Pending`상태일때,  우선 `describe`로 상태를 확인하고 밑의 오류와 같다면 명령어를 따라서 치고 다시 `pv`, `pvc`를 `delete`했다가 다시 만들어봅시다.

```cmd
$ kubectl describe pvc my-volumeclaim
Type    Reason      ...         Message
----    ------                  -------
Normal  WaitForFirstConsumer    waiting for first consumer to be created before binding
```

```cmd
kubectl get sc
kubectl delete sc kops-ssd-1-17
```



##### `Pod`로 volumeMount를 적용시켜 새로 하나 만들겠습니다.

**`$ vi pod.yaml`**

```cmd
apiVersion: v1
kind: Pod
metadata:
  name: my-user-app
  labels:
    app: my-user-app
spec:
  containers:
  - name: my-user-ms
    image: edowon0623/my-user-service:1.1
    ports:
      - containerPort: 8088
    volumeMounts:
      - mountPath: /my-volume
        name: my-hostpath
  volumes:
  - name: my-hostpath
    # hostPath:						# 주석처리 한 부분은 Linux에 직접 volumeMount를
    #   path: /tmp					# 한것이지만 우리는 PVC를 만들어서 연결을 합니다.
    #   type: Directory
    persistentVolumeClaim:
      claimName: my-volumeclaim
```



---

 `hostPath`로 했을때의 테스트방법입니다. 

##### 만들어진 pod내부에 접속하여 volume이 정상적으로 걸렸는지 test를 위해 비어있는 파일을 만들겠습니다.

```cmd
$ kubectl exec -it my-user-app sh
# cd /my-volume
# touch test.txt
```

##### `ip-172-20-51-89.ec2.internal`에 해당하는 노드를 찾아 `PublicIP`에 접속합니다.

```cmd
$ ssh 35.153.105.2
```

##### 저희가 `yaml`에서 `volumeMount`로 설정했던 `/tmp`로 이동을 합니다.  -- `$ cd /tmp`

- `ll`명령어로 하위 파일들을 검색하여 `test.txt`파일이 생성이 되어있는지 확인합니다.

```cmd
/tmp$ ll
-rw-r--r--  1 root root    0 Mar 24 07:55 test.txt
```

------



##### `Service`는 전에 작업했던 것에서 `name`부분만 추가를 했습니다.

```cmd
apiVersion: v1
kind: Service
metadata:
  name: my-user-service
spec:
  selector:
    app: my-user-app
  ports:
  - name: "808"
    port: 8080
    targetPort: 8080
  type: NodePort
```

```cmd
$ kubectl get po -o wide
NAME          READY   STATUS    RESTARTS   AGE   IP          NODE                       
my-user-app   1/1     Running   0          42m   100.96.1.6  ip-172-20-51-89.ec2.internal
```



- `$ vi mysql-deployment.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql-deployment
  labels:
    app: mysqldb
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mysqldb
  template:
    metadata:
      labels:
        app: mysqldb
    spec:
      containers:
      - name: mysql
        image: mysql:5.7
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: <ROOT_PASS>				
        - name: MYSQL_DATABASE
          value: <DATABASE_NAME>
        - name: MYSQL_USER
          value: <USER_NAME>
        - name: MYSQL_PASSWORD
          value: <USER_PASS>
        - name: MYSQL_ROOT_HOST
          value: '%'						# 모두 허용 ex)ec2에서 0.0.0.0:0과 같은 의미
        ports:
        - containerPort: 3306
        volumeMounts:
        - name: mysql-pv-storage
          mountPath: /var/lib/mysql
      volumes:
      - name: mysql-pv-storage
        persistentVolumeClaim:
          claimName: my-volumeclaim			# pvc에서 생성한 이름
```



##### 여기서 문제가 `PASSWORD`값을 `value`에 직접적으로 입력을 한다면 **보안적으로 문제가 있습니다.** 그렇기때문에 저희는 String형태로 보이는 것이 아닌 bit형태로 보이도록 `secret`을 만들어 적용하겠습니다.

```cmd
$ kubectl create secret generic mysql-password --from-literal=password=mysql 
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
mysql-password        Opaque                                1      7s
```



**`secret`을 수정할땐 3가지 방법이 있습니다.**

1. `edit`을 사용합니다. 저희가 `yaml`파일을 생성하지 않고 만들었지만 `edit`으로 가능합니다.

```cmd
$ kubectl edit secret mysql-password
```



2.  `yaml`파일이 없으면 output으로 만들 수 있습니다.

```cmd
$ kubectl get secret mysql-password -o yaml > secret.yaml
```



3. 새로 `kubectl create`를 하고 `pod`를 지우는 방법도 있지만 제일 않좋은 방법입니다.



##### `yaml`로 접속해서 바꾸고싶은 비밀번호를 `base64`로 바꾼 다음 `password`에 고치면 됩니다.

```yaml
apiVersion: v1
data:
  password: bXlzcWw=
kind: Secret
	:
```

```cmd
$ echo bXlzcWw= | base64 --decode
mysql

$ echo -n "hello" | base64
aGVsbG8=
```



##### `secret`을 만들었다면 다시 `$ vi mysql-deployment.yaml`명령어로 들어가서 `MYSQL_ROOT_PASSWORD`의 `value`를 지우고 새로 입력하겠습니다.

```yaml
- name: MYSQL_ROOT_PASSWORD
  valueFrom:
	secretKeyRef:
	  name: mysql-password
	  key: password
- name: MYSQL_DATABASE
  value: mydb
- name: MYSQL_USER
  value: k8suser
- name: MYSQL_PASSWORD
  value: k8spass
- name: MYSQL_ROOT_HOST
  value: '%'
```

```cmd
$ kubectl apply -f mysql-deployment.yaml
```



- 조회

```cmd
$ kubectl get deploy
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
mysql-deployment   1/1     1            1           2m1s
```

```cmd
$ kubectl get po
NAME                                READY   STATUS    RESTARTS   AGE
my-user-app                         1/1     Running   0          90m
mysql-deployment-5b8dbc58f6-vlzwg   1/1     Running   0          117s
```



##### Service 생성

- **`$ vi mysql-service.yaml`**

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql-service
  labels:
    app: mysqldb-svc
spec:
  type: NodePort
  ports:
    - port: 3306
  selector:
    app: mysqldb
```

```cmd
$ kubectl apply -f mysql-service.yaml
```



- 조회

```cmd
$ kubectl get svc
NAME              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
mysql-service     NodePort    100.68.190.214   <none>        3306:30622/TCP   4s
```





### Test

```cmd
$ yum install -y mysql
```

```cmd
$ mysql -uroot -p -h127.0.0.1 --port 30622
Enter password: <ROOT_PASSWORD>

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mydb               |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)
```



