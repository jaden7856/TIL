# WHAT IS OPERATOR SDK?
> Kubernetes 의 기본 애플리케이션을 효과적이고 자동화되고 확장 가능한 방식으로 관리하기 위한 오픈 소스 툴킷인 Operator Framework 의 구성 요소입니다.

<br>

## WHAT CAN I DO WITH OPERATOR SDK?

Operator SDK는 Operators를 빌드, 테스트 및 패키지하는 도구를 제공합니다. 처음에 SDK는 애플리케이션의 비즈니스 로직(예: 확장, 업그레이드 또는 백업 방법)을 
Kubernetes API와 결합하여 해당 작업을 실행하는 것을 용이하게 합니다.

Operator SDK는 `controller-runtime` 라이브러리를 사용하여 다음을 제공하여 연산자를 더 쉽게 작성하는 프레임워크입니다.

- 운영 로직을 보다 직관적으로 작성하기 위한 고급 API 및 추상화
- 새 프로젝트를 빠르게 부트스트랩하기 위한 스캐폴딩 및 코드 생성 도구
- 일반적인 연산자 사용 사례를 다루는 확장

<br>
<br>

## Workflow

SDK는 Go, Ansible 또는 Helm에서 연산자를 개발하기 위한 워크플로를 제공합니다. 우리는 이 중에서 Go를 사용할 것입니다.
다음 워크플로는 Go를 위한 것입니다.

- SDK 명령줄 인터페이스(CLI)를 사용하여 새 운영자 프로젝트 생성
- CRD(Custom Resource Definitions)를 추가하여 새 리소스 API 정의
- 자원을 감시하고 조정하는 컨트롤러 정의
- SDK 및 컨트롤러 런타임 API를 사용하여 컨트롤러에 대한 조정 로직 작성
- SDK CLI를 사용하여 운영자 배포 매니페스트 빌드 및 생성

<br>
<br>

## Install

앞서 말했듯이 이 자료에서는 Go를 이용하여 SDK를 사용할 것이기 때문에 Go로 SDK를 설치하겠습니다. 
**유의사항으로 `operator-sdk`릴리즈 버전에 따라 golang version 을 맞춰주세요 현재(22.11.03)에선 1.19 버전이 필요합니다.**
```
git clone https://github.com/operator-framework/operator-sdk
cd operator-sdk
git checkout master
make install

# 해당 명령어를 입력했을때 version이 출력되면 설치 완료입니다. 
operator-sdk version
```

<br>
<br>

# 