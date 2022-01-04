# 티베로 설치 - Linux

> 명령어 앞에 `#` 는 root계정이고, `$`는 자기계정입니다.

### 1. 라이센스 신청

- 라이센스 신청하는 방법은 [게시물](https://github.com/jaden7856/TIL/blob/master/Tibero/1_Tibero-License.md)을 참조합니다.





### 2. 설치 전 환경 구성

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

현재 커널 파라미터 값을 확인

```
# ipcs -l
```

```shell
------ Messages Limits --------
max queues system wide = 32000
max size of message (bytes) = 8192
default max size of queue (bytes) = 16384

------ Shared Memory Limits --------
max number of segments = 4096							# SHMMNI
max seg size (kbytes) = 18014398509465599				# SHMMAX
max total shared memory (kbytes) = 18014398442373116	# SHMALL
min seg size (bytes) = 1

------ Semaphore Limits --------
max number of arrays = 128								# SEMMNI
max semaphores per array = 250
max semaphores system wide = 32000
max ops per semop call = 32
semaphore max value = 32767
```



- 설정 파일

`sysctl.conf`파일에 들어간 후 맨밑 하단에 밑의 값들을 추가합니다.

```
# vi /etc/sysctl.conf
```

```
#tibero
kernel.sem = 10000 32000 10000 10000
kernel.shmall = ceil(SHMMAX/PAGE_SIZE)값보다 크게	# Linux 기본 PAGE_SIZE는 4096
kernel.shmmax = 8014398509465599		# 시스템의 물리적인 메모리 절반 (byte 단위)
kernel.shmmni = 4096
fs.file-max = 6815744
net.ipv4.ip_local_port_range = 1024 65000
```

시스템 메모리가 16GB이고 Tibero RDBMS가 사용할 최대 메모리가 8GB인 경우를 기준으로 설정

 `sysctl.conf`에 추가한 후에 **`sysctl -p`명령어로 동적 적용**



##### 2-4 .bash_profile 수정

```shell
$ vi .bash_profile
```

```shell
export TB_HOME=[티베로 설치 위치]		 # ex) /home/centos/tibero6
export TB_SID=[사용하고자하는 DB명]		# ex) tibero
export TB_PROF_DIR=$TB_HOME/bin/prof
export LD_LIBRARY_PATH=$TB_HOME/lib:$TB_HOME/client/lib:$LD_LIBRARY_PATH
export PATH=$PATH:$TB_HOME/bin:$TB_HOME/client/bin:$PATH
```

```shell
$ source .bash_profile
```





### 3. 설치하기

1) [다운로드](https://technet.tmaxsoft.com/ko/front/download/viewDownload.do?cmProductCode=0301&version_seq=PVER-20150504-000001&doc_type_cd=DN#binary) 클릭후 자신의 운영체제에 맞게 저는 **Linux (x86) 64-bit** 를 다운받았습니다.

2. Linux server에 `license.xml`파일과 `tibero6....tar.gz`파일을 넣어주고 압축을 해제

   - ```
     # tar -zxvf [파일이름]
     ```

3. `tibero6`폴더안에 license폴더에 `license.xml`파일을 넣어주겠습니다.





### 4. 티베로 환경파일 생성

환경변수로 설정한 `$TB_HOME`밑에 config폴더에서 `gen_tip.sh`를 실행하면 위와 같이 티베로 환경파일이 생성됩니다.

```shell
$ $TB_HOME/config/gen_tip.sh
```



#### 4-1 base_env.sh: file not found

환경파일 생성(4번) 명령어를 실행했을때 위와같이 오류가 뜬다면 **환경변수에서 `bash_profile`
에 제대로 값을 입력했는지 확인해야 합니다. 이 문제에 직면하신 분들은 반드시 echo $TB_HOME과 cd $TB_HOME으로 제대로 설정해놓았는지 확인해보시길..**





### 5. 기동

- 아직 DB랑 컨트롤 파일, 로그파일등등이 생성 되지 않았기 때문에 nomount모드로 부팅

```
$ tbboot nomount
```



- DB내부 연결할때는 `tbsql` 사용 (sys 사용자로 접속)

```
$ tbsql sys/tibero
```



- tibero라는 database도 새로 만들었다.(몇초정도 기다려야함)

```
SQL> create database "tibero";
```



- sql문을 나갈때는 `quit` 커맨드 사용.

  이제 database를 만들었으니 `tbboot`로만 기동

```
$ tbboot

Change core dump dir to /home/vagrant/tibero6/bin/prof.
Listener port = 8629

Tibero 6

TmaxData Corporation Copyright (c) 2008-. All rights reserved.
Tibero instance started up (NORMAL mode).
```



기본 명령어나 사용숙지는**[tbSQL 사이트](https://technet.tmaxsoft.com/upload/download/online/tibero/pver-20160406-000002/tibero_util/ch01.html)에서 확인하시면 됩니다.**

