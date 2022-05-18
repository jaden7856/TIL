# NFS(Network File System)
NFS(Network File System)는 네트워크에 파일을 저장하는 메커니즘입니다. 이 시스템은 사용자가 원격 컴퓨터에 있는 파일 및 디렉토리에 액세스할 수 있고 
해당 파일 및 디렉토리가 로컬에 있는 것처럼 처리하도록 허용하는 분산 파일 시스템입니다.

예를 들어, 사용자는 운영 체제 명령을 사용하여 원격 파일 및 디렉토리에 대한 파일 속성을 작성, 제거, 읽기, 쓰기, 설정할 수 있습니다.
한마디로 **네트워크를 이용해서 여러 서버간의 마운트를 통해 마치 하나의 파일 스토리지 처럼 사용하는 기술**입니다. 

서버와 클라이언트간의 파일 공유하는 방법은 여러가지가 있습니다. 대표적으로는 Samba가 유명하며, Linux, Window 모든 OS에 호환되어 다양하게 사용됩니다.

하지만 닉스나 리눅스 환경에서는 NFS 서버를 구축하여 서로간의 손쉽게 디렉터리를 공유할 수 있습니다. NFS 로 네트워크 망을 통해 Linux(Unix) 컴퓨터간의 저장 공간을 공유합니다. (Window 제외)

<br>

### 특징.
- NFS로 공유한 파일은 일반적인 소유권과 접근 권한이 그대로 적용
- 원격으로 접속한 유저가 파일 소유 권한이 있다면, UID를 이용해 조작 가능.
- 루트의 권한을 갖는 사용자가 해당 공유 디렉터리에서 모두 조작 가능.
- **단점으로 보안에 취약**

<br>
<br>

## NFS 서버 구축

### 1. NFS 확인 후 설치

```
// 설치 확인
$ rpm -qa | grep nfs-utils
nfs-utils-2.3.3-51.el8.x86_64

// 설치
$ yum install nfs-utils
```

<br>

### 2. 공유폴더 리스트


- `/`아래 디렉터리 생성

```
$ mkdir /backups
```

<br>

#### 2-1. NFS Server exports 옵션
NFS 파일 공유 설정파일은 /etc/exports 파일을 설정하시면 됩니다
- IP, IP대역, *(와일드카드)를 사용하여 IP를 지정

| 옵션  | 설명  |
|-----|-----|
| ro    | 읽기만 허용    |
| rw    | 읽기/쓰기 허용    |
| no-root-squash    | 관리자 권한 부여    |
| sync    | 파일을 쓸 때 서버와 클라이언트 싱크를 맞춘다. 서버는 데이터가 저장소에 안전히 쓰였음을 확인 한 후, 응답을 보낸다    |
| async    | 서버는 데이터가 저장소에 안전히 저장됐는지를 확인 하지 않는다. 클라이언트의 데이터 쓰기 요청이 들어오면 바로 응답을 보낸다 데이터 curruption이 발생할 수 있지만 성능 향상을 기대할 수 있다    |
| noaccess    | 디렉토리를 접근하지 못하게 한다. 공유된 디렉토리의 특정 하위 디렉토리만 접근하지 못하도록 제한할때 사용하는 옵션입니다.    |
| no_root_squash    | 클라이언트가 root 권한 획득가능, 파일 생성시 클라이언트의 권한으로 생성됨    |

```
$ vi /etc/exports

/backups 192.168.33.*(rw,sync)
```

<br>

#### 2-2. 디렉터리 권한 부여
```
$ chmod 707 /share
```

<br>

#### 2-3. 수정내용 반영
```
$ exportfs -r
```

<br>

### 3. 실행 및 확인
```
// 실행
$ systemctl start nfs-server
$ systemctl enable nfs-serve

// 확인
$ exportfs -v
/backups        192.168.33.*(sync,wdelay,hide,no_subtree_check,sec=sys,rw,secure,root_squash,no_all_squash)

// 확인
$ showmount -e
Export list for xfs:
/backups 192.168.33.*
```

<br>

### 4. 방화벽 끄기
```
$ service firewalld stop

Redirecting to /bin/systemctl stop firewalld.service
```

<br>
<br>

## NFS 클라이언트 구축

### 1. NFS 패키지 확인 및 설치
```
// 설치 확인
$ rpm -qa | grep nfs-utils
nfs-utils-2.3.3-51.el8.x86_64

// 설치
$ yum install nfs-utils
```

<br>

### 2. NFS 서버 마운트 확인
```
$ showmount -e 192.168.33.146

Export list for 192.168.33.146:
/backups 192.168.33.*
```

<br>

### 3. 마운트 할 디렉터리 생성
```
$ mkdir /nfs
```

<br>

### 4. 연결 및 확인
```
// 연결
$ mount -t nfs 192.168.33.146:/backups /nfs

// 확인
$ df -hT

Filesystem              Type      Size  Used Avail Use% Mounted on
192.168.33.146:/backups nfs4      100G  2.8G   98G   3% /nfs
```
