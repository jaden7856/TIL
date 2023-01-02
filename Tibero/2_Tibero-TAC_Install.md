# Installation

> 본 문서에서는 TAC-TAS, TAC, 1:1 Tibero 이 3가지 설치중에서 **TAC만 있는 구성을 설치하는 문서입니다.**

<br>
<br>

## 1. VM 구성

**VM 구성에서는 [TAS-TAC 설치방법](2_Tibero-TAC-TAS_Install.md)문서에서 (4)번 'VMware 설정' 까지 1~4번 내용을 따라하면 된다.**

<br>
<br>

## 2. Raw Device 생성

Raw Device 설정은 [TAS-TAC 설치방법](2_Tibero-TAC-TAS_Install.md)문서에서 (5)번처럼 `.rules`파일을 생성해서 symlink를 걸어서 사용해도 되고
지금부터 설명할 lvm을 사용해서 해도 상관없습니다.

<br>

### 2-1. 공유 Disk 할당
초기 VM을 설정하면 Disk 추가한 개수에 따라 sdb, sdc... 가 생성이 된다.

두개 이상 추가했다면 각 디스크마다 똑같이 수행을 하면 된다.
```
$ fdisk -l /dev/sdb /dev/sdc

Disk /dev/sdb: 10.7 GB, 10737418240 bytes, 20971520 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes


Disk /dev/sdc: 10.7 GB, 10737418240 bytes, 20971520 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
```

```
$ fdisk /dev/sdb
n
p
Enter
Enter
Enter
w

$ fdisk /dev/sdc
n
p
Enter
Enter
Enter
w
```

```
$ pvcreate /dev/sdb1
  Physical volume "/dev/sdb1" successfully created.
  
$ pvcreate /dev/sdc1
  Physical volume "/dev/sdc1" successfully created.
  
## 확인
$ fdisk -l /dev/sdb /dev/sdc

Disk /dev/sdb: 10.7 GB, 10737418240 bytes, 20971520 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disk label type: dos
Disk identifier: 0xdd7aa996

   Device Boot      Start         End      Blocks   Id  System
/dev/sdb1            2048    20971519    10484736   83  Linux

Disk /dev/sdc: 10.7 GB, 10737418240 bytes, 20971520 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disk label type: dos
Disk identifier: 0x9a6db0ca

   Device Boot      Start         End      Blocks   Id  System
/dev/sdc1            2048    20971519    10484736   83  Linux
```

<br>

### 2-2. Volumn Group 생성

```
$ vgcreate vg1 /dev/sdb1
  Volume group "vg1" successfully created
  
$ vgcreate vg2 /dev/sdc1
  Volume group "vg2" successfully created
```

<br>

### 2-3. Create Logical Volume
Logical Volume 크기는 물리디스크 크기에 따라 유기적으로 변경하시면 됩니다.

```
# CM File
lvcreate -L 128M -n lvtbcm_128m_01 vg1
lvcreate -L 128M -n lvtbcm_128m_01 vg2

# Control File
lvcreate -L 128M -n lvctl_128m_01 vg1
lvcreate -L 128M -n lvctl_128m_01 vg2

# Redo Log File
lvcreate -L 512M -n redo_log_lv512m_01 vg1
lvcreate -L 512M -n redo_log_lv512m_02 vg1
lvcreate -L 512M -n redo_log_lv512m_03 vg1
lvcreate -L 512M -n redo_log_lv512m_01 vg2
lvcreate -L 512M -n redo_log_lv512m_02 vg2
lvcreate -L 512M -n redo_log_lv512m_03 vg2

# Undo Data File
lvcreate -L 512M -n undo_data_lv512m_01 vg1
lvcreate -L 512M -n undo_data_lv512m_01 vg2

# Data File
lvcreate -L 1024M -n datafile_lv1024m_01 vg1
lvcreate -L 1024M -n datafile_lv1024m_02 vg1
lvcreate -L 1024M -n datafile_lv1024m_03 vg1
lvcreate -L 1024M -n datafile_lv1024m_04 vg1
lvcreate -L 1024M -n datafile_lv1024m_01 vg2
lvcreate -L 1024M -n datafile_lv1024m_02 vg2
lvcreate -L 1024M -n datafile_lv1024m_03 vg2
lvcreate -L 1024M -n datafile_lv1024m_04 vg2
```

<br>

### 2-4. View logical volume in LVM

밑의 명령어 2개 모두 LV가 Physical Volume과 어떻게 연결이 되어있는지 확인하는 명령어이다.

```
$ lvs -o +devices

  LV                  VG     Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert Devices
  root                centos -wi-ao---- <44.00g                                                     /dev/sda2(1280)
  swap                centos -wi-ao----   5.00g                                                     /dev/sda2(0)
  datafile_lv1024m_01 vg1    -wi-a-----   1.00g                                                     /dev/sdb1(576)
  datafile_lv1024m_02 vg1    -wi-a-----   1.00g                                                     /dev/sdb1(832)
  datafile_lv1024m_03 vg1    -wi-a-----   1.00g                                                     /dev/sdb1(1088)
  datafile_lv1024m_04 vg1    -wi-a-----   1.00g                                                     /dev/sdb1(1344)
  lvctl_128m_01       vg1    -wi-a----- 128.00m                                                     /dev/sdb1(32)
  lvtbcm_128m_01      vg1    -wi-a----- 128.00m                                                     /dev/sdb1(0)
  redo_log_lv512m_01  vg1    -wi-a----- 512.00m                                                     /dev/sdb1(64)
  redo_log_lv512m_02  vg1    -wi-a----- 512.00m                                                     /dev/sdb1(192)
  redo_log_lv512m_03  vg1    -wi-a----- 512.00m                                                     /dev/sdb1(320)
  undo_data_lv512m_01 vg1    -wi-a----- 512.00m                                                     /dev/sdb1(448)
  datafile_lv1024m_01 vg2    -wi-a-----   1.00g                                                     /dev/sdc1(576)
  datafile_lv1024m_02 vg2    -wi-a-----   1.00g                                                     /dev/sdc1(832)
  datafile_lv1024m_03 vg2    -wi-a-----   1.00g                                                     /dev/sdc1(1088)
  datafile_lv1024m_04 vg2    -wi-a-----   1.00g                                                     /dev/sdc1(1344)
  lvctl_128m_01       vg2    -wi-a----- 128.00m                                                     /dev/sdc1(32)
  lvtbcm_128m_01      vg2    -wi-a----- 128.00m                                                     /dev/sdc1(0)
  redo_log_lv512m_01  vg2    -wi-a----- 512.00m                                                     /dev/sdc1(64)
  redo_log_lv512m_02  vg2    -wi-a----- 512.00m                                                     /dev/sdc1(192)
  redo_log_lv512m_03  vg2    -wi-a----- 512.00m                                                     /dev/sdc1(320)
  undo_data_lv512m_01 vg2    -wi-a----- 512.00m                                                     /dev/sdc1(448)
```
```
$ lsblk

NAME                           MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda                           8:0    0   50G  0 disk
├─sda1                        8:1    0    1G  0 part /boot
└─sda2                        8:2    0   49G  0 part
  ├─centos-root             253:0    0   44G  0 lvm  /
  └─centos-swap             253:1    0    5G  0 lvm  [SWAP]
sdb                           8:16   0   10G  0 disk
└─sdb1                        8:17   0   10G  0 part
  ├─vg1-lvtbcm_128m_01      253:2    0  128M  0 lvm
  ├─vg1-lvctl_128m_01       253:4    0  128M  0 lvm
  ├─vg1-redo_log_lv512m_01  253:6    0  512M  0 lvm
  ├─vg1-redo_log_lv512m_02  253:7    0  512M  0 lvm
  ├─vg1-redo_log_lv512m_03  253:8    0  512M  0 lvm
  ├─vg1-undo_data_lv512m_01 253:12   0  512M  0 lvm
  ├─vg1-datafile_lv1024m_01 253:14   0    1G  0 lvm
  ├─vg1-datafile_lv1024m_02 253:15   0    1G  0 lvm
  ├─vg1-datafile_lv1024m_03 253:16   0    1G  0 lvm
  └─vg1-datafile_lv1024m_04 253:17   0    1G  0 lvm
sdc                           8:32   0   10G  0 disk
└─sdc1                        8:33   0   10G  0 part
  ├─vg2-lvtbcm_128m_01      253:3    0  128M  0 lvm
  ├─vg2-lvctl_128m_01       253:5    0  128M  0 lvm
  ├─vg2-redo_log_lv512m_01  253:9    0  512M  0 lvm
  ├─vg2-redo_log_lv512m_02  253:10   0  512M  0 lvm
  ├─vg2-redo_log_lv512m_03  253:11   0  512M  0 lvm
  ├─vg2-undo_data_lv512m_01 253:13   0  512M  0 lvm
  ├─vg2-datafile_lv1024m_01 253:18   0    1G  0 lvm
  ├─vg2-datafile_lv1024m_02 253:19   0    1G  0 lvm
  ├─vg2-datafile_lv1024m_03 253:20   0    1G  0 lvm
  └─vg2-datafile_lv1024m_04 253:21   0    1G  0 lvm
sr0                          11:0    1 1024M  0 rom
```

<br>

### 2-5. 유저생성

만약 Linux에 자기 계정(root 이외)이 없다면 하나 생성 후 시작하겠습니다.

```
$ groupadd dba -g 1024
$ useradd tibero -g 1024 -u 1024
```

<br>

### 2-6. `.rules` 파일 생성

- 모든 노드 적용
```rules
$ vi /etc/udev/rules.d/99-tibero.rules

ACTION!="add|change", GOTO="raw_end"

ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="lvtbcm_128m_01", RUN+="/usr/bin/raw /dev/raw/raw1 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="lvtbcm_128m_01", RUN+="/usr/bin/raw /dev/raw/raw2 %N"

ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="lvctl_128m_01", RUN+="/usr/bin/raw /dev/raw/raw3 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="lvctl_128m_01", RUN+="/usr/bin/raw /dev/raw/raw4 %N"

ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="redo_log_lv512m_01", RUN+="/usr/bin/raw /dev/raw/raw5 %N"
ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="redo_log_lv512m_02", RUN+="/usr/bin/raw /dev/raw/raw6 %N"
ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="redo_log_lv512m_03", RUN+="/usr/bin/raw /dev/raw/raw7 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="redo_log_lv512m_01", RUN+="/usr/bin/raw /dev/raw/raw8 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="redo_log_lv512m_02", RUN+="/usr/bin/raw /dev/raw/raw9 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="redo_log_lv512m_03", RUN+="/usr/bin/raw /dev/raw/raw10 %N"

ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="undo_data_lv512m_01", RUN+="/usr/bin/raw /dev/raw/raw11 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="undo_data_lv512m_01", RUN+="/usr/bin/raw /dev/raw/raw12 %N"

ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="datafile_lv1024m_01", RUN+="/usr/bin/raw /dev/raw/raw13 %N"
ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="datafile_lv1024m_02", RUN+="/usr/bin/raw /dev/raw/raw14 %N"
ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="datafile_lv1024m_03", RUN+="/usr/bin/raw /dev/raw/raw15 %N"
ENV{DM_VG_NAME}=="vg1", ENV{DM_LV_NAME}=="datafile_lv1024m_04", RUN+="/usr/bin/raw /dev/raw/raw16 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="datafile_lv1024m_01", RUN+="/usr/bin/raw /dev/raw/raw17 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="datafile_lv1024m_02", RUN+="/usr/bin/raw /dev/raw/raw18 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="datafile_lv1024m_03", RUN+="/usr/bin/raw /dev/raw/raw19 %N"
ENV{DM_VG_NAME}=="vg2", ENV{DM_LV_NAME}=="datafile_lv1024m_04", RUN+="/usr/bin/raw /dev/raw/raw20 %N"

KERNEL=="raw*", OWNER=="tibero", GROUP=="dba", MODE=="0660"

LABEL="raw_end"
```

<br>

### 2-7. udev rules 적용

```
$ udevadm control --reload-rules
$ udevadm trigger
```

- 확인 및 재부팅
```
$ reboot

$ ll /dev/raw
```

<br>
<br>

## 3. 설치 전 환경 구성
> 모든 노드 설정 필요

### 3-1. hosts 설정

```
$ vi /etc/hosts

192.168.x.x tibero1
192.168.x.x tibero2
```

<br>

### 3-2. JDK 적용

앞의 [TAS-TAC 설치방법](2_Tibero-TAC-TAS_Install.md)문서 (2)에서 JDK 다운 받았던 bin 파일을 vm 서버로 옮겨주고 설치하겠습니다.

```
$ tar -zxvf jdk-17_linux-x64_bin.tar.gz

$ mv jdk-17.0.5 /usr/lib/

$ vi /etc/profile

export JAVA_HOME=/usr/lib/jdk-17.0.5   # 맨 밑에 적습니다.
export PATH=$PATH:$JAVA_HOME/bin

=====================================
$ source /etc/profile
```

<br>

### 3-3 패키지 설치

Tibero 설치 시 필요한 패키지들을 설치

```
$ yum install -y gcc gcc-c++ libgcc libstdc++ libstdc++-devele compat-libstdc++ libaio libaio-devel pstack libpthread librt libm libaio libdl
```

<br>

### 3-4 커널 파라미터 설정
> **참고**
>
> SELinux를 enforce mode로 사용할 경우 프로세스가 비정상적인 동작을 할 수 있어 Tibero를 사용하는 경우 enforce mode로 설정하는 것을 권장하지 않는다.

- 방화벽 해제
```
$ systemctl stop firewalld && systemctl disable firewalld
```

**master와 node들의 메모리 설정이 다르다면 메모리 설정에 따른 값을 넣어주세요**

```
$ vi /etc/sysctl.conf

#tibero
kernel.sem = 10000 32000 10000 10000
kernel.shmall = 2097152       # ceil(shmmax/PAGE_SIZE) Linux 기본 PAGE_SIZE는 4096
kernel.shmmax = 8589934592    # 16GB이면 8589934592. 8GB이면 4294967296.
kernel.shmmni = 4096
fs.file-max = 67108864
fs.aio-max-nr = 1048576
net.ipv4.ip_local_port_range = 1024 65000
net.ipv4.tcp_rmem = 4194304
net.ipv4.tcp_wmem = 1048576
net.core.rmem_default = 262144
net.core.wmem_default = 262144
net.core.rmem_max = 4194304
net.core.wmem_max = 67108864
```

시스템 메모리가 16GB이고 Tibero RDBMS가 사용할 최대 메모리가 8GB인 경우를 기준으로 설정

`sysctl.conf`에 추가한 후에 **`sysctl -p`명령어로 동적 적용**

<br>

### 3-5. Shell Limits 파라미터

```
# nofile = MAX_SESSION_COUNT / WTHR_PROC_CNT 이상으로 설정한다.
#        - Soft Limit : 65536
#        - Hard Limit : 65536

# nproc = MAX_SESSION_COUNT+10000 이상으로 설정한다.
#        - Soft Limit : 65536
#        - Hard Limit : 65536
```

```
$ vi /etc/security/limits.conf

tibero soft nproc 65536
tibero hard nproc 65536
tibero soft nofile 65536
tibero hard nofile 65536
```

<br>

### 3-6 .bash_profile 수정

> node 들에서 `TB_SID`와 `CM_SID`는 각각 다르게 이름을 지어준다.

- root 계정

```shell
[노드 1번]
# Tibero 7 Env
export TB_HOME=/tibero/tibero7
export PATH=.:$TB_HOME/bin:$TB_HOME/client/bin:$PATH
export LD_LIBRARY_PATH=$TB_HOME/lib:$TB_HOME/client/lib:$LD_LIBRARY_PATH

# Tibero7 CM ENV
export CM_HOME=$TB_HOME
export CM_SID=tbcm1

[노드 2번]
# Tibero 7 Env
export TB_HOME=/tibero/tibero7
export PATH=.:$TB_HOME/bin:$TB_HOME/client/bin:$PATH
export LD_LIBRARY_PATH=$TB_HOME/lib:$TB_HOME/client/lib:$LD_LIBRARY_PATH

# Tibero7 CM ENV
export CM_HOME=$TB_HOME
export CM_SID=tbcm2
```
```
$ source .bash_profile
```

<br>

- tibero 계정

앞서 계정을 만든 `tibero`에 접속(`su - tibero`)하여 환경변수 설정하겠습니다.

```shell
$ vi .bash_profile

[노드 1번]
# Tibero 7 Env
export TB_HOME=/tibero/tibero7
export TB_SID=tac1
export TB_PROF_DIR=$TB_HOME/bin/prof
export PATH=.:$TB_HOME/bin:$TB_HOME/client/bin:$PATH
export LD_LIBRARY_PATH=$TB_HOME/lib:$TB_HOME/client/lib:$LD_LIBRARY_PATH

export CM_SID=tbcm1
export CM_HOME=$TB_HOME

# Tibero aliases
alias tbhome='cd $TB_HOME'
alias tbbin='cd $TB_HOME/bin'
alias tblog='cd /logs/tibero'
alias tbcfg='cd $TB_HOME/config'
alias tbcfgv='vi $TB_HOME/config/$TB_SID.tip'
alias tbcli='cd ${TB_HOME}/client/config'
alias tbcliv='vi ${TB_HOME}/client/config/tbdsn.tbr'

=========

[노드 2번]
# Tibero 7 Env
export TB_HOME=/tibero/tibero7
export TB_SID=tac2
export TB_PROF_DIR=$TB_HOME/bin/prof
export PATH=.:$TB_HOME/bin:$TB_HOME/client/bin:$PATH
export LD_LIBRARY_PATH=$TB_HOME/lib:$TB_HOME/client/lib:$LD_LIBRARY_PATH

export CM_SID=tbcm2
export CM_HOME=$TB_HOME

# Tibero aliases
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

<br>
<br>

## 4. 티베로 설치하기
> 모든 노드

앞서 (2. JDK 및 Tibero 다운로드)에서 자신의 운영체제에 맞게 저는 **Linux (x86) 64-bit** 를 다운받았습니다.

라이센스발급받았던 `license.xml`파일과 `tibero7....tar.gz`파일을 넣어주고 압축을 해제 후 `tibero7/license`폴더에 `license.xml`파일을 넣어주겠습니다.
```
$ su -
$ mkdir /tibero					                  # 경로생성

$ mv tibero7....tar.gz /tibero/		          # tibero binary 옮기기

$ cd /tibero 					                  # 경로에 가서 tibero binary 풀어줌
$ tar -zxvf tibero7....tar.gz


$ mv /root/license.xml /tibero/tibero7/license/   # license 인스턴스의 license 경로 아래 옮겨줌

$ chown -R tibero:dba /tibero			          # tibero 하위 경로 오너쉽 변경
```

<br>
<br>

## 5. 티베로 환경파일 생성

> 모든 노드

환경변수로 설정한 `$TB_HOME`밑에 config폴더에서 `gen_tip.sh`를 실행하면 위와 같이 티베로 환경파일이 생성됩니다.

```shell
$ su - tibero
$ $TB_HOME/config/gen_tip.sh		# alias를 했다면 tbcfg 하고 ./gen_tip.sh

Using TB_SID "tac1"   # Node 2 에서는 "tac2"
/home/vagrant/tibero6/config/tac1.tip generated
/home/vagrant/tibero6/config/psm_commands generated
/home/vagrant/tibero6/client/config/tbdsn.tbr generated.
Running client/config/gen_esql_cfg.sh
Done.
```

<br>
<br>

## 6. Tibero Tip File 수정

> 모든 노드

1번 노드를 위한 `CM TIP` 파일을 1번 노드의 `$TB_HOME/config` 아래에`tbcm1.tip`으로, 2 번 노드를 위한 `CM TIP` 파일을 2번 노드의 `$TB_HOME/config` 아래에 `tbcm2.tip`으로 저장하였으며, 다음 과 같이 TIP 파일을 작성하였습니다. (config 폴더에 저장해야 한다).

### 6-1. `cm.tip` 설정

- Node 1
```
$ vi $TB_HOME/config/$CM_SID.tip

CM_NAME=tbcm1
CM_UI_PORT=8635
CM_RESOURCE_FILE="/tibero/tibero7/config/tbcm1_res.crf"
```

- Node 2
```
$ vi $TB_HOME/config/$CM_SID.tip

CM_NAME=tbcm2
CM_UI_PORT=8655
CM_RESOURCE_FILE="/tibero/tibero7/config/tbcm2_res.crf"
```

<br>

### 6-2. `tac.tip` 설정

**TAC 설치 파라미터입니다.** 기존에 있던 모든 값들은 삭제하거나 주석처리 합니다.

- Node 1
```
$ vi $TB_HOME/config/$TB_SID.tip	## alias를 했다면 tbcfgv

### TAC ENV ###
DB_NAME=tac
LISTENER_PORT=8629
CONTROL_FILES="/dev/raw/raw3","/dev/raw/raw4"
DB_CREATE_FILE_DEST="/tibero/tbdata"

MAX_SESSION_COUNT=50
MEMORY_TARGET=2G
TOTAL_SHM_SIZE=1G

CLUSTER_DATABASE=Y
THREAD=0
UNDO_TABLESPACE=UNDO0

LOCAL_CLUSTER_ADDR=192.168.40.130
LOCAL_CLUSTER_PORT=21100
CM_PORT=8635
```

- Node 2
```
$ vi $TB_HOME/config/$TB_SID.tip	## alias를 했다면 tbcfgv

### TAC ENV ###
DB_NAME=tac
LISTENER_PORT=8629
CONTROL_FILES="/dev/raw/raw3","/dev/raw/raw4"
DB_CREATE_FILE_DEST="/tibero/tbdata"

MAX_SESSION_COUNT=50
MEMORY_TARGET=2G
TOTAL_SHM_SIZE=1G

CLUSTER_DATABASE=Y
THREAD=1
UNDO_TABLESPACE=UNDO1

LOCAL_CLUSTER_ADDR=192.168.40.131
LOCAL_CLUSTER_PORT=21110
CM_PORT=8655
```

<br>

### 6-3. `tbdns.tbr` 설정

- Node 1
```
tac1=(
    (INSTANCE=(HOST=localhost)
              (PORT=8629)
              (DB_NAME=tac)
    )
)

tac=(
    (INSTANCE=(HOST=192.168.33.150)
    (PORT=8629)
    (DB_NAME=tac)
    )
    (INSTANCE=(HOST=192.168.33.151)
    (PORT=8629)
    (DB_NAME=tac)
    )
    (LOAD_BALANCE=Y)
    (USE_FAILOVER=Y)
)
```

- Node 2
```
tac2=(
    (INSTANCE=(HOST=localhost)
              (PORT=8629)
              (DB_NAME=tac)
    )
)

tac=(
    (INSTANCE=(HOST=192.168.33.151)
    (PORT=8629)
    (DB_NAME=tac)
    )
    (INSTANCE=(HOST=192.168.33.150)
    (PORT=8629)
    (DB_NAME=tac)
    )
    (LOAD_BALANCE=Y)
    (USE_FAILOVER=Y)
)
```

<br>
<br>

## 7. DB Cluster 구성

- Node 1

```shell
ens33: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.33.138  netmask 255.255.255.0  broadcast 192.168.33.255
        inet6 fe80::109f:6f6d:9546:8903  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:68:87:d1  txqueuelen 1000  (Ethernet)
        RX packets 449527  bytes 646455486 (616.5 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 34876  bytes 3161917 (3.0 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

ens34: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.40.130  netmask 255.255.255.0  broadcast 192.168.40.255
        inet6 fe80::70c4:de2c:ec94:308  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:68:87:db  txqueuelen 1000  (Ethernet)
        RX packets 476  bytes 57588 (56.2 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 16  bytes 1202 (1.1 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```

```
# root 로 tbcm을 시작
$ su -

$ tbcm -b

$ chown -R tibero:dba /tibero

$ su - tibero

$ cmrctl add network --name inc1 --nettype private --ipaddr 192.168.40.130 --portno 29000
Resource add success! (network, inc1)

$ cmrctl add network --name pub1 --nettype public --ifname ens33
Resource add success! (network, pub1)

$ cmrctl add cluster --name cls1 --incnet inc1 --pubnet pub1 --cfile "/dev/raw/raw1,/dev/raw/raw2"
Resource add success! (cluster, cls1)

$ cmrctl start cluster --name cls1
MSG SENDING SUCCESS!

$ cmrctl add service --name tac --cname cls1
Resource add success! (service, tac)

$ cmrctl add db --name tac1 --svcname tac --dbhome $TB_HOME --envfile /home/tibero/.bash_profile
Resource add success! (db, tac1)
```

cluster 추가에서 저는 `cfile`저장 장소를 2개 했기때문에 밑의 결과처럼 `cls1:0`, `cls1:1` 두개가 생성되었습니다.
```
$ cmrctl show

Resource List of Node tbcm1
=====================================================================
  CLUSTER     TYPE        NAME       STATUS           DETAIL
----------- -------- -------------- -------- ------------------------
     COMMON  network           inc1       UP (private) 192.168.40.130/29000
     COMMON  network           pub1       UP (public) ens33
     COMMON  cluster           cls1       UP inc: inc1, pub: pub1
       cls1     file         cls1:0       UP /dev/raw/raw1
       cls1     file         cls1:1       UP /dev/raw/raw2
       cls1  service            tac     DOWN Database, Active Cluster (auto-restart: OFF)
       cls1       db           tac1     DOWN tac, /tibero/tibero7, failed retry cnt: 0
=====================================================================
```

<br>

- Node 2

```
ens33: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.33.139  netmask 255.255.255.0  broadcast 192.168.33.255
        inet6 fe80::f4ae:f51b:3459:7b9f  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:51:96:81  txqueuelen 1000  (Ethernet)
        RX packets 448157  bytes 646336507 (616.3 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 34087  bytes 3027293 (2.8 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

ens34: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.40.131  netmask 255.255.255.0  broadcast 192.168.40.255
        inet6 fe80::449e:3f63:2235:9b3b  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:51:96:8b  txqueuelen 1000  (Ethernet)
        RX packets 468  bytes 57099 (55.7 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 18  bytes 1382 (1.3 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```

```
$ su -

$ tbcm -b

$ chown -R tibero:dba /tibero

$ su - tibero

$ cmrctl add network --name inc2 --nettype private --ipaddr 192.168.40.131 --portno 19629
Resource add success! (network, inc1)

$ cmrctl add network --name pub2 --nettype public --ifname ens33
Resource add success! (network, pub1)

$ cmrctl add cluster --name cls1 --incnet inc2 --pubnet pub2 --cfile "/dev/raw/raw1,/dev/raw/raw2"
Resource add success! (cluster, cls1)

$ cmrctl start cluster --name cls1
MSG SENDING SUCCESS!
```

<br>
<br>

## 8. TAC를 위한 데이터베이스 생성

- Node 1

logfile, datafile, tempfile 등 경로를 지정할 때 자기가 저장하고 싶은 공유디스크 경로 어디든 해도 상관은 없습니다.

```
$ vi create_db.sql

create database "tac"
user sys identified by tibero
character set UTF8
national character set utf16
logfile group 0 ('/dev/raw/raw5') size 512M,
        group 1 ('/dev/raw/raw6') size 512M,
        group 2 ('/dev/raw/raw7') size 512M
maxdatafiles 4096
maxlogfiles 100
maxlogmembers 8
noarchivelog
datafile '/dev/raw/raw13' size 1024M autoextend off
default tablespace USR
datafile '/dev/raw/raw14' size 1024M autoextend off
default temporary tablespace TEMP
tempfile '/dev/raw/raw15' size 1024M autoextend off
extent management local AUTOALLOCATE
undo tablespace UNDO0
datafile '/dev/raw/raw11' size 512M autoextend off
extent management local AUTOALLOCATE;
```
```
$ cmrctl start db --name tac1 --option "-t nomount"
$ tbsql sys/tibero @create_db.sql

tbSQL 7

TmaxTibero Corporation Copyright (c) 2020-. All rights reserved.

Connected to Tibero.


Database created.
```

<br>
<br>

### 8-2. Node 2 Thread1 추가
- Node 1
```
$ tbboot

Change core dump dir to /tibero/tibero7/bin/prof.
Listener port = 8629

Tibero 7

TmaxTibero Corporation Copyright (c) 2020-. All rights reserved.
Tibero instance started up (NORMAL mode).
```

```
$ tbsql sys/tibero

SQL> create undo tablespace undo1 datafile '/dev/raw/raw12' size 512M autoextend off
extent management local autoallocate; 

Tablespace 'UNDO1' created.

SQL> alter database add logfile thread 1 group 3 '/dev/raw/raw8' size 512M;
Database altered.

SQL> alter database add logfile thread 1 group 4 '/dev/raw/raw9' size 512M;
Database altered.

SQL> alter database add logfile thread 1 group 5 '/dev/raw/raw10' size 512M;
Database altered.

SQL> alter database enable public thread 1;
Database altered.

SQL> q
Disconnected.
```

<br>

### 8-3. tpr 테이블스페이스 생성 구문 수정
```
$ vi /tibero/tibero7/scripts/tpr_create_obj.sh

tpr_create_ts_sql="create tablespace syssub datafile 'syssub001.dtf' size 10m reuse autoextend on next 10m;"

## 위에 부분을 아래처럼 변경

tpr_create_ts_sql="create tablespace syssub datafile '/dev/raw/raw16' size 1024m reuse autoextend off;"
```

<br>

### 8-4. system shell

```
$ sh $TB_HOME/scripts/system.sh -p1 tibero -p2 syscat -a1 y -a2 y -a3 y -a4 y
```

<br>

### 8-5. Node 2에서 db 생성

```
$ cmrctl add db --name tac2 --svcname tac --dbhome $TB_HOME --envfile /home/tibero/.bash_profile
Resource add success! (db, tac2)

$ cmrctl start db --name tac2

Listener port = 8629

Tibero 7

TmaxTibero Corporation Copyright (c) 2020-. All rights reserved.
Tibero instance started up (NORMAL mode).
BOOT SUCCESS! (MODE : NORMAL)
```
```
$ cmrctl show

Resource List of Node tbcm2
=====================================================================
  CLUSTER     TYPE        NAME       STATUS           DETAIL
----------- -------- -------------- -------- ------------------------
     COMMON  network           inc2       UP (private) 192.168.40.131/19629
     COMMON  network           pub2       UP (public) ens33
     COMMON  cluster           cls1       UP inc: inc2, pub: pub2
       cls1     file         cls1:0       UP /dev/raw/raw1
       cls1     file         cls1:1       UP /dev/raw/raw2
       cls1  service            tac       UP Database, Active Cluster (auto-restart: OFF)
       cls1       db           tac2 UP(NRML) tac, /tibero/tibero7, failed retry cnt: 0
=====================================================================
```

<br>
<br>

## 9. 복제

Tibero를 복제하여 정상 가동하기 위해서는 CM file, Control file, Redo file, Datafiles, Undo file이 필요합니다.

이것들의 저장위치는 `(8. TAC를 위한 데이터베이스 생성)`에서 저장 위치를 지정하였으며 CM file은 Cluster를 등록할때 `/dev/raw/raw1`,
`/dev/raw/raw2` 이 두곳에 저장하는것으로 지정을 했고, Control file은 `tbcfgv`명령어를 입력하여 tac 설정파일에서 `CONTROL FILES`
ENV를 통햬 저장장소를 지정했습니다.

 