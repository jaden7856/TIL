# Docker에서 Kubernetes 설치

- pply & Restart 클릭

![image-20210215143542154](00_Kubernetes_Download.assets/image-20210215143542154.png)



- `install` 클릭

![image-20210215143721904](00_Kubernetes_Download.assets/image-20210215143721904.png)



#### 만약 설치오류가 떠서 되지 않는다면 minikube를 설치해보자! 

> 위의 설치가 된다면 하지않아도 된다.

- 제어판 -> 프로그램 및 기능 -> Windows 기능 켜기/끄기 -> Hyper-V를 체크! (Windows 10 Pro 화면입니다.)

![image-20210215151840223](00_Kubernetes_Download.assets/image-20210215151840223.png)



- 원하는 장소에 minikube라는 이름의 폴더를 생성후, 다음 명령으로 최신 릴리스를 다운로드한 파일을  minikube 폴더에 옮긴다.

```
curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
```



-  [minikube-windows-amd64.exe](https://github.com/kubernetes/minikube/releases/download/v1.17.1/minikube-windows-amd64.exe) 를 다운로드 해서 minikube 폴더에 저장후 경로를 복사

-  시작에서 검색하여 `시스템 환경 변수 편집`을 클릭

![image-20210215153502115](00_Kubernetes_Download.assets/image-20210215153502115.png)

`환경 변수`를 클릭



![image-20210215153539289](00_Kubernetes_Download.assets/image-20210215153539289.png)

Path에서 편집 클릭



![image-20210215153609462](00_Kubernetes_Download.assets/image-20210215153609462.png)

새로 만들기를클릭하여 복사했던 경로를 추가한다.



- 다음 명령 프롬프트 창을 열어서

  - `$  minikube version`을 해서 확인
  - `$ minikube start`
