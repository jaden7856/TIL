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
[root@go184]# go get -d -u go-micro.dev/v4
[root@go183]# go get -d -u github.com/micro/micro/v3/cmd/protoc-gen-micro
[root@go183]# go install github.com/micro/micro/v3/cmd/protoc-gen-micro
[root@go184]# go get -d -u golang.org/x/lint/golint
[root@go184]# go install golang.org/x/lint/golint
[root@go184]# go get -d -u github.com/t-yuki/gocover-cobertura
[root@go184]# go install github.com/t-yuki/gocover-cobertura
[root@go184]# exit

### 실행중인 컨테이너를 이미지로 만들기
[root]# docker commit -a "devops" -m "golang1.18.4" go184 <DIR>/golang:1.18.4

### 레지스트리 푸쉬
[root]# docker push <DIR>/golang:1.18.4 
```
저희 회사에서 gloang 이미지에 `google api` 라이브러리와 `go-micro`, `protobuf`를 넣어 사용하기때문에 설치를 해서 이미지를 만들었습니다.

하지만 현재 `go-micro`의 v4.8.0버전에서는 `protoc-gen-micro`가 없는 관계로 v3버전으로 다운을 받았습니다.

그 다음 정상작동 확인을 위해 `docker run -d -it --name go184 <DIR>/golang:1.18.4` 명령어로 확인합니다.

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

<br>

### 2-3. build

```
[root]# docker build -t <DIR>/<NAME>:<VERSION> .
```

<br>
<br>

## 3. ymal파일