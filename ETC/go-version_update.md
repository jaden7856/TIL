# Go version update

회사에서 사용하는 라이브러리들과 기능들을 전체적으로 업데이트 해야하는 업무를 가지게 되었습니다.
처음엔 가볍게 생각했지만 회사내의 서비스들이 Microservice인데다가 방대하여 고치는데 많은 애를 먹었는데요.
간단하게 작업 순서와 기타 방법들을 정리했습니다.

## 1. golang 이미지 파일 만들기

```
### go 이미지 풀 및 실행
[root]# docker pull golang:1.18.4
[root]# docker run -d -it --name go184 golang:1.18.4
[root]# docker exec -it go184 /bin/bash

### 최신 업데이트
[root@go183]# cd
[root@go183]# apt update
[root@go183]# apt upgrade
[root@go183]# go version
go version go1.18.4 linux/amd64

### google-api 설치 (go는 /go/ 에 설치되어 있음)
[root@go183]# git clone https://github.com/googleapis/googleapis.git
[root@go183]# cp -arpvi googleapis/google /go/

### protobuf 설치
[root@go183]# apt install protobuf-compiler
[root@go183]# protoc --version
libprotoc 3.12.4

### protobuf 라이브러리 설치
[roor@go184]# mkdir test
[roor@go184]# cd test
[roor@go184]# go mod init test
[root@go184]# go get -d -u google.golang.org/protobuf
[root@go184]# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
[root@go184]# go get -d -u github.com/micro/micro/v3/cmd/protoc-gen-micro
[root@go184]# go install github.com/micro/micro/v3/cmd/protoc-gen-micro
[root@go184]# go get -d -u github.com/golang/protobuf/protoc-gen-go
[root@go184]# go get -d -u github.com/micro/micro/v2/cmd/protoc-gen-micro
[root@go184]# go get -d -u go-micro.dev/v4
[root@go184]# go get golang.org/x/lint/golint
[root@go184]# go get github.com/t-yuki/gocover-cobertura
[root@go184]# exit

### 실행중인 컨테이너를 이미지로 만들기
[root]# docker commit -a "devops" -m "golang1.18.4" go184 registry.datacommand.co.kr/golang:1.18.4

### 레지스트리 푸쉬
[root]# docker push registry.datacommand.co.kr/golang:1.18.4 
```
우선 테스트에선 `protobuf`와 `go-micro`를 **v2**와 **v3**, **v4**버전을 다 `go get` 했습니다. 실 사용 golang에선 현재 최신버전인
**v4**만 사용합니다.

<br>
<br>

## 2. Dockerfile Build

### 2-1.변경해야 할 파일

뭔저 build를 하기전에 변경해야 할 파일들이 있습니다. 아래의 파일들을 바뀐 이미지에 맞게 변경합니다.

- go.mod (go 1.14 -> go 1.18 로 변경)
- .gitlab-ci.yml (FROM 1.14 -> FROM 1.18.4 로 변경)
- Dockerfile (FROM 1.14 -> FROM 1.18.4 로 변경)

<br>

### 2-2. go.mod 

go version이 `1.14`에서 `1.18`로 바뀌면서 그냥 build를 하면 `go.mod` 내의 바뀐 버전을 적용시키기위해 `go mod tidy`와 `go mod vendor`를 해줍니다. 

- connection refused 에러가 뜬다면 변수설정 
```dockerfile
export GOFLAGS='-mod=vendor'
export GOPRIVATE=10.1.1.220/**
export GOINSECURE=10.1.1.220/**
```

<br>

### 2-3. build

```
[root]# docker build -t registry.datacommand.co.kr/<NAME>:<VERSION> .
```

각 서비스마다 `docker build`를 하여 image를 생성하고 test 하도록 하겠습니다. 특정 서비스만 서비스하고싶다면 빌드 후 실행을 **cdm-cloud** -> **cdm-center**
-> **cdm-dr** -> **cdm-replicator** 순서로 하시면 됩니다.

<br>

### 2-4 error
```dockerfile
The command '/bin/sh -c GOFLAGS='-mod=vendor' CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o registerd cmd/registerd/registerd.go' returned a non-zero code: 2
```
만약 `docker build`할때 위와같은 오류가 발생한다면 common폴더에 파일이 없는것입니다.

<br>
<br>

## 3. ymal파일