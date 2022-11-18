# Where are Docker images stored on the host machine?

> `/var/lib/docker` 디렉토리에 저장이 되며 내용은 Docker가 스토리지에 사용하는 드라이버에 따라 다릅니다.

사용하는 드라이버는 **커널 지원에 따라 `overlay`, `overlay2`, `btrfs`, `devicemapper` 또는 `zfs`가 될 수 있습니다. 


`-s` 또는 `--storage-driver=` 옵션을 사용하여 Docker 데몬에 스토리지 드라이버를 수동으로 설정할 수 있습니다.

- `/var/lib/docker/{driver-name}`에는 이미지 내용에 대한 드라이버별 저장소가 포함됩니다.
- `/var/lib/docker/graph/<id>`에는 json 및 layerize 파일의 이미지에 대한 메타데이터만 포함합니다.

#### In the case of `aufs`
- `/var/lib/docker/aufs/diff/<id>`에는 이미지의 파일 내용이 있습니다.
- `/var/lib/docker/repositories-aufs`는 로컬 이미지 정보가 포함된 JSON 파일입니다. 이 정보는 `docker images` 명령으로 볼 수 있습니다.

#### In the case of `device-mapper`
- `/var/lib/docker/devicemapper/devicemapper/data`는 이미지를 저장합니다.
- `/var/lib/docker/devicemapper/devicemapper/metadata` 메타데이터

이러한 파일은 thin provisioned "sparse" 파일이기 때문에 보기만큼 크지 않습니다

---

<br>

### Docker 이미지 및 컨테이너의 저장 위치
Docker 컨테이너는 네트워크 설정, 볼륨 및 이미지로 구성됩니다. Docker 파일의 위치는 운영 체제에 따라 다릅니다. 다음은 가장 많이 사용되는 운영 체제에 대한 개요입니다.

- **Ubuntu** - `/var/lib/docker/`
- **CentOS** - `/var/lib/docker/`
- **Fedora** - `/var/lib/docker/`
- **Debian** - `/var/lib/docker/`
- **Windows** - `C:\ProgramData\DockerDesktop`
- **MacOS** - `~/Library/Containers/com.docker.docker/Data/vms/0/`


현재 제가 사용하고있는 CentOS 에서 밑의 명령어를 실행하면 `overlay2` Driver 를 사용중인걸 확인할 수 있습니다.

```
$ docker info

---
Storage Driver: overlay2
Docker Root Dir: /var/lib/docker
---
```

---

<br>

## Docker Images
기본 storage driver `overlay2`를 사용하는 경우 Docker 이미지는 `/var/lib/docker/overlay2`에 저장됩니다.
여기서 Docker 이미지의 읽기 전용 레이어와 변경 내용을 포함하는 레이어를 나타내는 다양한 파일을 찾을 수 있습니다.

```
$ docker inspect golang:1.19.3

[
    {
        "Id": "sha256:6bc...bdf",
        "RepoTags": [
            "golang:1.19.3"
        ],
        "RepoDigests": [
            "golang@sha256:435...27b"
        ],
        "Parent": "",
  ...
        "Architecture": "amd64",
        "Os": "linux",
        "Size": 2777932612,
        "VirtualSize": 2777932612,
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/17b...cc8/diff:
                             /var/lib/docker/overlay2/b6d...0ad/diff:
                             ...
                             /var/lib/docker/overlay2/4c2...381/diff:
                             /var/lib/docker/overlay2/5b6...f2f/diff",
                "MergedDir": "/var/lib/docker/overlay2/b8c...152/merged",
                "UpperDir": "/var/lib/docker/overlay2/b8c...152/diff",
                "WorkDir": "/var/lib/docker/overlay2/b8c...152/work"
            },
            "Name": "overlay2",
  ...
```
**LowerDir** 에는 이미지의 읽기 전용 레이어가 포함되어 있습니다. **UpperDir** 에는 변경 사항을 나타내는 읽기-쓰기 계층입니다.

저는 golang:1.19.3 이미지에 따로 옮겼던 파일이 있는데 그런 로그 파일이 포함되어 있는 곳이 **UpperDir**입니다.

```
$ ls -la /var/lib/docker/overlay2/4da...371/diff

total 0
drwxr-xr-x. 5 root root 39 Oct 28 02:17 .
drwx--x---. 4 root root 55 Oct 28 02:19 ..
drwxrwxrwx. 4 root root 28 Nov  3 00:52 go
drwx------. 3 root root 41 Oct 28 02:16 root
drwxrwxrwt. 2 root root  6 Oct 28 02:15 tmp

```

---
