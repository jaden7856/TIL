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


- Runtime Controller Package
  - Controller가 관리할 CR의 변경을 감지하고 SDK Controller Package의 Reconcile Loop가 동작하도록 합니다.
- Runtime Manager Package
  - Kubernetes Client 및 Cache를 초기화합니다.
- SDK Controller Package
  - 실제 Controller 로직을 수행하는 Reconcile Loop와 Runtime Manager Package로 부터 전달받은 Kubernetes Client가 포함되어 있습니다.

<br>

### 동작 흐름

- **#Kubernetes**

![img.png](00_Intro-operator-sdk.assets/1.png)

<br>

- **Operator**

![img.png](00_Intro-operator-sdk.assets/2.png)

<br>
<br>

## WHAT IS CR(Custom Resource)?
- **Custom Resource**는 Custom Object의 모음이며 API 확장을 위한 기본 리소스로 사용할 객체를 정의하고 추상화한 구조적 데이터와 같습니다.


- **Custom Resource**에서의 객체 정의는 완전히 새로운 객체가 아닌 **이미 존재하는 Deployments, Services와 같은 기본 객체를 목적에 맞게 조합하고 
추상화해서 새로운 이름으로 명시할 수 있다는 의미**입니다. 


- **Custom Resource Definition**은 Costom Resource가 데이터로 어떤 항목이 정의되어야 하는지 등을 저장하는 선언적 메타데이터 객체일 뿐입니다.
(XML Schema 와 XML의 관게를 생각하면 이해하기 좋음.)


- **Custom Resource**는 `kubectl` 을 통해서 사용 가능.

<br>
<br>

## Kubernetes 용어

### **Kubernetes Object**

**"Persistent entities in the Kubernetes system"** 로 정의되어 있으며 Kubernetes에 저장된 실체들 이라고 생각하면 됩니다.

이 객체들로 다음과 같은 정보를 나타낼 수 있습니다.
- 애플리케이션이 배정된 노드들
- 애플리케이션이 사용할 수 있는 리소스들
- 애플리케이션 동작에 대한 정책 (어떻게 재 시작할지, 업데이트할지 등)

**Kubernetes 객체들은 생성했을 때 생성된 실체가 존재하는 것이 아니라 `Status`를 의미합니다.** 예를 들어 `kubectl`을 통해 Pod 1개를 생성하는 
요청을 보내면 Kubernetes는 1개의 Pod가 필요한 Status를 기록합니다. 그리고 Kubernetes는 현재의 Status와 기록된 Status를 비교해서 원하는 
Status를 맞추도록 동작하게 되는 원리입니다.

<br>

### Object Spec and Status

모든 Kubernetes 객체들은 공통적으로 두 개의 필드를 가지게 됩니다. `create api`를 하게 되면 `/api` 폴더에 `_types.go` 파일이 생성이 됩니다.
그 안의 코드중 `Spec`과 `Status` 객체 설명입니다.
- **`Spec`** - 객체가 가질 Status에 대한 명세 정보
- **`Status`** - 실체 클러스터에서 객체가 가진 상태 정보 (Kubernetes가 계속 검증하고 반영)

<br>

### Kubernetes API

Kubernetes의 객체를 이용해서 CRUD 작업을 하기 위해서는 Kubernetes API를 통해야 합니다. 즉, 사용자가 `kubectl`을 사용해서 객체 생성 명령을 실행하면 
`kubectl`은 Kubernetes API로 요청을 하고 Kubernetes는 해당 객체를 생성하게 됩니다. 

이 설명은 위에서 설명했던 [동작 흐름](#동작-흐름)

ex) 물론 `kubectl`이 아닌 Kubernetes API 클라이언트 라이브러리를 통해서 작업도 가능

<br>
<br>

### 참고
- https://sdk.operatorframework.io/docs/overview/
- https://frozenpond.tistory.com/111
- https://ccambo.blogspot.com/2020/12/kubernetes-operator-kubernetes-operator.html