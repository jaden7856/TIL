# Reflink 사용법
앞서 [Xfs_Reflink](1_Xfs_Reflinks.md)에서 말한것과 같이 시스템에서 CRC 검사를 활성화해야 하며 매우 흥미로운 점은 사용된 데이터 블록 
크기가 1KiB에서 4KiB 사이여야 한다는 점입니다. Linux의 XFS는 페이지 크기 또는 더 작은 블록의 파일 시스템만 마운트할 수 있습니다.

이러한 것들이 정상인지 사전에 확인합니다. 현재 테스트하는 Linux의 사양입니다.

![img.png](2_Reflink_사용법.assets/1.png)

```
$ yum install -y lshw

$ lshw -class disk -class storage -short

H/W path              Device      Class          Description
============================================================
/0/100/7.1            scsi2       storage        82371AB/EB/MB PIIX4 IDE
/0/100/7.1/0.0.0      /dev/cdrom  disk           VMware IDE CDR10
/0/100/7.1/0.0.0/0    /dev/cdrom  disk
/0/100/10             scsi0       storage        53c1030 PCI-X Fusion-MPT Dual Ultra320 SCSI
/0/100/10/0.0.0       /dev/sda    disk           107GB VMware Virtual S
/0/100/10/0.1.0       /dev/sdb    disk           107GB VMware Virtual S
```

<br>

그런 다음 명령으로 원하는 섹션을 만듭니다. 이것을 하는 이유는 이미 생성된 파일 시스템에서는 `-rflink`플래그를 활성화할 수 없기 때문입니다.

```
$ mkfs.xfs -b size=4096 -m crc=1,reflink=1 /dev/sdb
```

명령어에서 `crc=1`, `reflink=1` 와 같은 것을 볼 수 있습니다. 이것이 우리가 필요로 하는 것입니다. 솔직히 `crc=1`이 기본적으로 설정되어 있지만 명확성을 
위해 이 작업을 수행했습니다.

<br>

다음은 백업용 폴더를 만들고 마운트합니다.
```
$ mkdir /backups
$ mount /dev/sdb /backups
```

마지막으로 모든 것이 정상인지 다음 명령어로 확인합니다.
```
$ df -hT

Filesystem              Type      Size  Used Avail Use% Mounted on
/dev/sdb                xfs       100G   33M  100G   1% /backups
```

<br>
<br>

## Test
이제 XFS와 reflink가 어떻게 작동하는지 확인해 보겠습니다. 이를 위해 `urandom`에서 리디렉션하는 가장 좋아하는 방법을 사용하여 임의 콘텐츠가 
포함된 파일을 생성합니다.

```
$ dd if=/dev/urandom of=test bs=1M count=10240

10240+0 records in
10240+0 records out
10737418240 bytes (11 GB) copied, 50.1595 s, 214 MB/s
```

`count=`에서 용량은 자신이 원하는 용량을 넣어도 됩니다. 저는 약 2-3분가량 시간이 지나 완료가 되었습니다. reflink는 기본적으로 시스템에서 
사용되지 않기 때문에 여기서는 중복 제거가 표시되지 않습니다. 지금 우리가 정말로 관심을 갖고 있는 것은 데이터 자체가 차지하는 공간과 
메타데이터에 필요한 공간입니다.

<br>

지금 `df -h`를 수행하면 11GB가 중에 1GB의 메타데이터를 차지하고 새로 만든 10GB가 있음을 알 수 있습니다.
```
$ df -h

Filesystem               Size  Used Avail Use% Mounted on
/dev/sdb                 100G   11G   90G  11% /backups
```

<br>

reflink 매개변수를 사용하여 이 동일한 파일의 복사본을 시작할 것입니다. 결과는 다음과 같습니다.
```
$ cp -v --reflink=always test test-one

‘test’ -> ‘test-one’
cp: failed to clone ‘test-one’ from ‘test’: Operation not supported
```

`--reflink`에서 `always`가 아닌 `auto`를 v8.32 이후 의 주요 릴리스는 기본적으로 `cp`에서 `reflink=auto`를 시도합니다.

하지만 제가 사용하는 CentOS Kernel에선 위와 같이 `always`를 사용하면 `Operation net suppored`라는 에러가 떨어집니다. 저의 
Kernel 버전은 아래와 같으며 알아본 결과 Centos 일부 버전에서 오류가 발생하는걸 확인하였습니다.

```
$ uname -srm

Linux 3.10.0-1160.62.1.el7.x86_64 x86_64
```

<br>

정상적으로 작동을 하였다면 아래의 명령어를 실행했을때 각 파일에 대해 10GB의 디스크 공간이 20GB 사용되었음을 알려줍니다.
```
$ ls -hsl

total 20G
10G -rw-r--r--. 1 root root 10G May 12 20:50 test
10G -rw-r--r--. 1 root root 10G May 12 20:56 test-one
```

<br>

그러나 디스크 공간 사용량을 확인하는 명령을 실행하면 **디스크가 증가하지 않았습니다.**
```
$ df -h

Filesystem               Size  Used Avail Use% Mounted on
/dev/sdb                 100G   11G   90G  11% /backups
```

<br>

아래와 같은 명령어를 사용하여 두개의 파일이 동일한 범위를 가지고 디스크 내부의 물리적 위치가 같음을 알 수 있습니다.
두 파일 모두 3개의 extents를 차지하며 블록의 정확도와 동일하게 위치합니다.
```
$ filefrag -v test test-one

Filesystem type is: 58465342
File size of test is 10737418240 (2621440 blocks of 4096 bytes)
 ext:     logical_offset:        physical_offset: length:   expected: flags:
   0:        0.. 1048559:         24..   1048583: 1048560:             shared
   1:  1048560.. 2097135:    1310733..   2359308: 1048576:    1048584: shared
   2:  2097136.. 2621439:    2624013..   3148316: 524304:    2359309: last,shared,eof
test: 3 extents found
File size of test-one is 10737418240 (2621440 blocks of 4096 bytes)
 ext:     logical_offset:        physical_offset: length:   expected: flags:
   0:        0.. 1048559:         24..   1048583: 1048560:             shared
   1:  1048560.. 2097135:    1310733..   2359308: 1048576:    1048584: shared
   2:  2097136.. 2621439:    2624013..   3148316: 524304:    2359309: last,shared,eof
test-one: 3 extents found
```



---

### 참고
- https://prog.world/xfs-reflink-and-fast-clone-made-for-each-other/
- https://jorgedelacruz.uk/2020/03/19/veeam-whats-new-in-veeam-backup-replication-v10-xfs-reflink-and-fast-clone-repositories-in-veeam/