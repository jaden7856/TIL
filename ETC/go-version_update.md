# Go version update

회사에서 사용하는 라이브러리들과 기능들을 전체적으로 업데이트 해야하는 업무를 가지게 되었습니다.
그중에서도 여기서 정리할 내용은 golang version upgrade 인데요.

처음엔 가볍게 생각했지만 회사내의 서비스들이 Microservice인데다가 방대하여 고치는데 많은 애를 먹었지만, 
간단하게 작업 순서와 기타 방법들을 정리했습니다. 

그래도 Golang은 버전 업데이트나 라이브러리 관리가 엄청 쉬운편이라 좋았습니다.

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
[root@go184]# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
[root@go183]# go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
[root@go184]# go get -d -u go-micro.dev/v4
[root@go184]# go install golang.org/x/lint/golint
[root@go184]# go install github.com/t-yuki/gocover-cobertura
[root@go184]# exit

### 실행중인 컨테이너를 이미지로 만들기
[root]# docker commit -a "devops" -m "golang1.18.4" go184 <DIR(option)>/golang:1.18.4

### 레지스트리 푸쉬
[root]# docker push <DIR>/golang:1.18.4 
```
저희 회사에서 gloang 이미지에 `google api` 라이브러리와 `go-micro`, `protobuf`를 넣어 사용하기때문에 설치를 해서 이미지를 만들었습니다.

하지만 현재 `go-micro`의 v4.8.0버전에서는 `protoc-gen-micro`가 없는 관계로 v3버전으로 다운을 받았습니다.
저야 `protoc-gen-micro`를 회사에서 사용하기 때문에 설치했지만 보통은 `protoc-gen-go-grpc`를 사용하기때문에 유의해주세요.

그 다음 정상작동 확인을 위해 `docker run -d -it --name go184 <DIR>/golang:1.18.4` 명령어로 확인합니다.

<br>

#### Go get
- `-d` : 설치는 하지 않고 소스 파일만 다운로드합니다.

- `-u` : 패키지 및 해당 종속성을 업데이트합니다.

- `-t` : 패키지에 대한 테스트를 빌드하는 데 필요한 패키지도 다운로드합니다.

- `-v` : 진행 및 디버그 출력

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
[root]# docker build -t <DIR(option)>/<NAME>:<VERSION> .
```

<br>

### 2-4. push

```
[root]# docker push <DIR(option)>/<NAME>:<VERSION>
```

<br>
<br>

### 2-5. Apply

이제 서비스들의 yaml 파일에 생성한 이미지의 name을 적어주고 실행합니다.
```
[root]# kubectl apply -f <ymal file name>
```

<br>

#### 2-6-1. kubectl pod의 logs 오류 `syntax error: unexpected word (expecting "do")`
- 위의 오류들은 앞에 특정 shell script파일이 windows에서 작업을 하여 linux에서 개행 오류로 인한 것입니다.
```
# vi -b yourscriptfile.sh 
```
위의 명령어로 `^M` 문자가 있다면 삭제를 해주세요