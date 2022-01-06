# Tibero Active Cluster

Tibero Active Cluster(TAC)는 **확장성, 고가용성을 목적으로 제공**하는 Tibero RDBMS의 주요 기능이다. TAC 환경에서 실행 중인 모든 인스턴스는 공유된 데이터베이스를 통해 트랜잭션을 수행하며, 공유된 데이터에 대한 접근은 데이터의 일관성과 정합성 유지를 위해 상호 통제하에 이뤄진다.

큰 업무를 작은 업무의 단위로 나누어 여러 노드 사이에 분산하여 수행할 수 있기 때문에 업무 처리 시간을 단축할 수 있다.

여러 시스템이 공유 디스크를 기반으로 데이터 파일을 공유한다. TAC 구성에 필요한 데이터 블록은 노드 간을 연결하는 고속 사설망을 통해 주고 받음으로써 노드가 하나의 공유 캐시(shared cache)를 사용하는 것처럼 동작한다.

운영 중에 한 노드가 멈추더라도 동작 중인 다른 노드들이 서비스를 지속하게 된다. 이러한 과정은 투명하고 신속하게 처리된다.



다음은 TAC의 간략한 구조를 나타내는 그림이다	[출처 : http://www.gurubee.net/lecture/2107]

![Tibero Active Cluster](http://www.gurubee.net/imgs/tibero/TAC_C_1.jpg)





# Installation

우선 사용하고자하는 노드의 숫자만큼 생성을 생성을 합니다. 저는 master와 node 2개를 생성했습니다. 



### 1. 라이센스 신청

- 라이센스 신청하는 방법은 [게시물](https://github.com/jaden7856/TIL/blob/master/Tibero/1_Tibero-License.md)을 참조합니다. 여기서 하나의 Tibero를 만드는 것과 차이는 **TAC Cluster를 구성하기 위해서 Standard가 아닌 Enterprise 라이센스를 신청해야합니다.**
- 각각 node들의 hostname에 맞게 발급받으세요.





### 2. 설치 전 환경 구성

> **master, node 2개 모두 설치**

##### 2-0. hosts 설정

```
# vi /etc/hosts
```

```
192.168.1.x master
192.168.1.x node1
192.168.1.x node2
```



##### 2-1. JDK 설치

Tibero를 설치하기 위해서는 `JDK 1.5.17` 이상 버전이 필요합니다. 저는 Openjdk 8 버전을 설치했습니다.

```
# yum install java-1.8.0-openjdk-devel.x86_64
```



##### 2-2 패키지 설치

Tibero 설치 시 필요한 패키지들을 설치

```
# yum install gcc gcc-c++ libgcc libstdc++ libstdc++-devele compat-libstdc++ libaio libaio-devel
```



##### 2-3 커널 파라미터 설정

**master와 node들의 메모리 설정이 다르다면 메모리 설정에 따른 값을 넣어주세요**

`sysctl.conf`파일에 들어간 후 맨밑 하단에 밑의 값들을 추가합니다.

```
# vi /etc/sysctl.conf
```

```
#tibero
kernel.sem = 10000 32000 10000 10000
kernel.shmmax = 8589934592		# 시스템의 물리적인 메모리 절반 (byte 단위)
kernel.shmall = ceil(SHMMAX/PAGE_SIZE)값보다 크게	# Linux 기본 PAGE_SIZE는 4096
kernel.shmmni = 4096
fs.file-max = 6815744
net.ipv4.ip_local_port_range = 1024 65000
```

시스템 메모리가 16GB이고 Tibero RDBMS가 사용할 최대 메모리가 8GB인 경우를 기준으로 설정

 `sysctl.conf`에 추가한 후에 **`sysctl -p`명령어로 동적 적용**



##### 2-4 .bash_profile 수정

> **master, node들에서 `TB_SID`와 `CM_SID`는 각각 다르게 이름을 지어준다. 
>
>  ex) master - tibero1, node1 - tibero2...

```
$ vi .bash_profile
```

```shell
# Tibero 6 Env
export TB_HOME=[티베로 설치 위치]
export TB_SID=[사용하고자하는 DB명] ## DB Instance 별 ID. 두번째 노드는 tibero2
export PATH=.:$TB_HOME/bin:$TB_HOME/client/bin:$PATH
export LD_LIBRARY_PATH=$TB_HOME/lib:$TB_HOME/client/lib:$LD_LIBRARY_PATH

export TB_PROF_DIR=$TB_HOME/bin/prof

# Tibero TBCM Env
export CM_SID=[Cluster ID Name]	## Cluster Manager ID. 두번째 노드는 tbcm2
export CM_HOME=[티베로 설치 위치]

# Tibero aliases		## aliases는 안해도 됩니다.
alias tbhome='cd $TB_HOME'
alias tbbin='cd $TB_HOME/bin'
alias tblog='cd /logs/tibero'
alias tbcfg='cd $TB_HOME/config'
alias tbcfgv='vi $TB_HOME/config/$TB_SID.tip'
alias tbcli='cd ${TB_HOME}/client/config'
alias tbcliv='vi ${TB_HOME}/client/config/tbdsn.tbr'
```

```
$ source .bash_profile
```





### 3. 설치하기

> master, node 모두 설치

1) [다운로드](https://technet.tmaxsoft.com/ko/front/download/viewDownload.do?cmProductCode=0301&version_seq=PVER-20150504-000001&doc_type_cd=DN#binary) 클릭후 자신의 운영체제에 맞게 저는 **Linux (x86) 64-bit** 를 다운받았습니다.

2. Linux server에 `license.xml`파일과 `tibero6....tar.gz`파일을 넣어주고 압축을 해제

   - ```
     # tar -zxvf [파일이름]
     ```

3. `tibero6`폴더안에 license폴더에 `license.xml`파일을 넣어주겠습니다.





### 4. 티베로 환경파일 생성

> master, node 모두 생성

환경변수로 설정한 `$TB_HOME`밑에 config폴더에서 `gen_tip.sh`를 실행하면 위와 같이 티베로 환경파일이 생성됩니다.

```shell
$ $TB_HOME/config/gen_tip.sh		# alias를 했다면 tbcfg 하고 ./gen_tip.sh
```





### 5. raw device 생성

Tibero TAC 는 공유 디스크를 기반으로 하기 때문에 데이터파일이 저장될 경로는 모든 DB 노드가 접근가능해야한다. 공유 파일시스템으로도 구성할 수 있지만 raw device 를 기반으로 설치하겠습니다.

1. 아래 사진과 같이 Hard Disk 2를 Virtual Machine에 추가합니다. 

![image-20220106102311090](2_Tibero-TAC.assets/image-20220106102311090.png)



2. 생성 확인

루트 계정으로 접속 후 `/dev/sdb`가 생성되었다면 정상 (개인 Disk 추가 갯수에따라 다르다.)

```
# ls -l /dev/sd*
brw-rw----. 1 root disk 8,  0 Jan  6 01:11 /dev/sda
brw-rw----. 1 root disk 8,  1 Jan  6 01:11 /dev/sda1
brw-rw----. 1 root disk 8, 16 Jan  6 01:11 /dev/sdb
```



3.  block device (logical volume) 생성

> fdisk 프롬프트에서 sdb1 primary partition 생성

**만약 `pvcreate`나 `vgcreate` 명령어가 안된다면 아래 패키지 설치**

```
yum install lvm2 -y
```



```
# fdisk /dev/sdb
# pvcreate /dev/sdb1			# 2개 이상이면 ex) /dev/sdb1 /dev/sdc1 식으로 생성
# vgcreate raws /dev/sdb1		# 2개 이상이면 disk를 합쳐서 용량을 늘릴 수 있다.
```



최소 10개 이상 만들수있는 용량으로 하는것이 좋다.

```
[root@master ~]# lvcreate -L 512M -n lv500g00 raws
[root@master ~]# lvcreate -L 512M -n lv500g01 raws
[root@master ~]# lvcreate -L 512M -n lv500g02 raws
[root@master ~]# lvcreate -L 512M -n lv500g03 raws
[root@master ~]# lvcreate -L 512M -n lv500g04 raws
[root@master ~]# lvcreate -L 512M -n lv500g05 raws
[root@master ~]# lvcreate -L 512M -n lv500g06 raws
[root@master ~]# lvcreate -L 512M -n lv500g07 raws
[root@master ~]# lvcreate -L 512M -n lv500g08 raws
[root@master ~]# lvcreate -L 512M -n lv500g09 raws
[root@master ~]# lvcreate -L 1G -n lv1g01 raws
[root@master ~]# lvcreate -L 1G -n lv1g02 raws
[root@master ~]# lvcreate -L 1G -n lv1g03 raws
[root@master ~]# lvcreate -L 1G -n lv1g04 raws
[root@master ~]# lvcreate -L 1G -n lv1g05 raws
[root@master ~]# lvcreate -L 5G -n lv5g00 raws
[root@master ~]# lvcreate -L 5G -n lv5g01 raws
```



###### raw

```
[root@master ~]# raw /dev/raw/raw1 /dev/raws/lv500g00
[root@master ~]# raw /dev/raw/raw1 /dev/raws/lv500g01
[root@master ~]# raw /dev/raw/raw2 /dev/raws/lv500g01
[root@master ~]# raw /dev/raw/raw1 /dev/raws/lv500g00
[root@master ~]# raw /dev/raw/raw3 /dev/raws/lv500g02
[root@master ~]# raw /dev/raw/raw4 /dev/raws/lv500g03
[root@master ~]# raw /dev/raw/raw5 /dev/raws/lv500g04
[root@master ~]# raw /dev/raw/raw6 /dev/raws/lv500g05
[root@master ~]# raw /dev/raw/raw7 /dev/raws/lv500g06
[root@master ~]# raw /dev/raw/raw8 /dev/raws/lv500g07
[root@master ~]# raw /dev/raw/raw9 /dev/raws/lv500g08
[root@master ~]# raw /dev/raw/raw10 /dev/raws/lv500g09
[root@master ~]# raw /dev/raw/raw11 /dev/raws/lv1g01
[root@master ~]# raw /dev/raw/raw12 /dev/raws/lv1g02
[root@master ~]# raw /dev/raw/raw13 /dev/raws/lv1g03
[root@master ~]# raw /dev/raw/raw14 /dev/raws/lv1g04
[root@master ~]# raw /dev/raw/raw15 /dev/raws/lv1g05
[root@master ~]# raw /dev/raw/raw16 /dev/raws/lv5g00
```





### 6. Tibero Tip File 수정

> master, node 모두 수정

1번 노드를 위한 `CM TIP` 파일을 1번 노드의 `$TB_HOME/config` 아래에`tbcm1.tip`으로, 2 번 노드를 위한 `CM TIP` 파일을 2번 노드의 `$TB_HOME/config` 아래에 `tbcm2.tip`으로 저장하였으며, 다음 과 같이 TIP 파일을 작성하였습니다. (config 폴더에 저장해야 한다).

- <tbcm1.tip>

```
CM_NAME=tbcm1						## 각각 이름 다르게
CM_UI_PORT=10010
CM_RESOURCE_FILE=/home/vagrant/tbcm1_resources
```

- <tbcm2.tip>

```
CM_NAME=tbcm2						## 각각 이름 다르게
CM_UI_PORT=10010
CM_RESOURCE_FILE=/home/vagrant/tbcm2_resources
```



기존에 있던 모든 값들은 삭제하거나 주석처리 합니다.

```
$ vi $TB_HOME/config/$TB_SID.tip	## alias를 했다면 tbcfgv
```

- <tibero.tip>

```v
### TAC ENV ###
DB_NAME=tibero		## master, node 모두 같은 이름
LISTENER_PORT=8629
DB_CREATE_FILE_DEST="/home/vagrant/tibero6/database/tibero"
LOG_ARCHIVE_DEST="/home/vagrant/tibero6/database/tibero/archivel"

MAX_SESSION_COUNT=20
TOTAL_SHM_SIZE=2G
MEMORY_TARGET=4G
UNDO_TABLESPACE=UNDO0				## node1 = UNDO1, node2 = UNDO2 각각 다르게
THREAD=0							## node1 = 1, node2 = 2 각각 숫자가 다르게

CLUSTER_DATABASE=Y					## Cluster 실행여부
LOCAL_CLUSTER_ADDR=192.168.x.131	## master, node 각 IP주소
LOCAL_CLUSTER_PORT=18631
CM_PORT=10010
```

- <tibero2.tip>

```v
### TAC ENV ###
DB_NAME=tibero
LISTENER_PORT=8629
CONTROL_FILES="dev/raws/rlv500g01"
DB_CREATE_FILE_DEST="/home/vagrant/tibero6/database/tibero"
LOG_ARCHIVE_DEST="/home/vagrant/tibero6/database/tibero/archivel"

MAX_SESSION_COUNT=20
TOTAL_SHM_SIZE=2G
MEMORY_TARGET=4G
UNDO_TABLESPACE=UNDO1
THREAD=1

CLUSTER_DATABASE=Y
LOCAL_CLUSTER_ADDR=192.168.x.132
LOCAL_CLUSTER_PORT=18631
CM_PORT=10010
```



```
[vagrant@master ~]$ tbcm -b
CM Guard daemon started up.

TBCM 6.1.1 (Build 186930)

TmaxData Corporation Copyright (c) 2008-. All rights reserved.

Tibero cluster manager started up.
Local node name is (tbcm1:10010).
```

```
# cmrctl add network --nettype private --ipaddr [IP] --portno 10013 --name prvnet
# cmrctl add network --nettype public --ifname eth0 --name pubnet
# cmrctl add cluster --incnet prvnet --pubnet pubnet --cfile "/dev/raws/rlv1g01" --name cls
# cmrctl start cluster --name cls
```



`Failed to start the resource 'cls'` 에러

```
# systemctl stop firewalld
```

