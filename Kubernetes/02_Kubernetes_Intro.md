# Kubernetes란?

쿠버네트스는 컨테이너 운영을 자동화하기 위한 컨테이너 오케스트레이션 도구로, 구글의 주도로 개발됐다. **많은 수의 컨테이너를 협조적으로 연동시키기 위한 통합 시스템**이며 이 컨테이너를 다루기 위한 API 및 명령행 도구 등이 함께 제공된다.

컨테이너를 이용한 애플리케이션 배포 외에도 다양한 운영 관리 업무를 자동화할 수 있다. 도커 호스트관리, 서버 리소스의 여유를 고려한 컨테이너 배치, 스케일링, 여러 개의 컨테이너 그룹에 대한 로드 밸런싱, 헬스 체크 등의 기능을 갖추고 있다.

**쿠버네트스의 가장 큰 특징은 다양한 부품을 조합해 유연한 애플리케이션을 구축할 수 있다는 점이다.**

<br>
<br>

## Kubernetes 용어

### Kubernetes Object

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

이 설명은 위에서 설명했던 [동작 흐름](#동작-흐름) (물론 `kubectl`이 아닌 Kubernetes API 클라이언트 라이브러리를 통해서 작업도 가능)

---

<br>
<br>

# Kubernetes 명령어

터미널 창에서 `kubectl` 명령을 실행하려면 다음의 구문을 사용한다.

```
$ kubectl [command] [TYPE] [NAME] [flags]
```

다음은 `command`, `TYPE`, `NAME` 과 `flags` 에 대한 설명이다.

- `command`: 하나 이상의 리소스에서 수행하려는 동작을 지정한다. 예: **`create`, `get`, `describe`, `delete`**

  

- `TYPE`: [리소스 타입](https://kubernetes.io/ko/docs/reference/kubectl/overview/#리소스-타입)을 지정한다. 리소스 타입은 대소문자를 구분하지 않으며 단수형, 복수형 또는 약어 형식을 지정할 수 있다. 예를 들어, 다음의 명령은 동일한 출력 결과를 생성한다.

  ```shell
  kubectl get pod pod1
  kubectl get pods pod1
  kubectl get po pod1
  ```

- `NAME`: 리소스 이름을 지정한다. 이름은 대소문자를 구분한다. 이름을 생략하면, 모든 리소스에 대한 세부 사항이 표시된다. 예: `kubectl get pods`
 
여러 리소스에 대한 작업을 수행할 때, 타입 및 이름별로 각 리소스를 지정하거나 하나 이상의 파일을 지정할 수 있다.

<br>

타입 및 이름으로 리소스를 지정하려면 다음을 참고한다.

- 리소스가 모두 동일한 타입인 경우 리소스를 그룹화하려면 다음을 사용한다. 
  - `TYPE1 name1 name2 name<#>`
  - 예) `kubectl get pod example-pod1 example-pod2`


- 여러 리소스 타입을 개별적으로 지정하려면 다음을 사용한다.
  - `TYPE1/name1 TYPE1/name2 TYPE2/name3 TYPE<#>/name<#>`
  - 예) `kubectl get pod/example-pod1 replicationcontroller/example-rc1`

<br>

하나 이상의 파일로 리소스를 지정하려면 다음을 사용한다. 
```shell
-f file1 -f file2 -f file<#>
```
- YAML이 특히 구성 파일에 대해 더 사용자 친화적이므로, [JSON 대신 YAML을 사용한다](https://kubernetes.io/ko/docs/concepts/configuration/overview/#일반적인-구성-팁).
  예: `kubectl get -f ./pod.yaml`


- `flags`: 선택적 플래그를 지정한다. 예를 들어, `-s` 또는 `--server` 플래그를 사용하여 쿠버네티스 API 서버의 주소와 포트를 지정할 수 있다.

> **주의:** 커맨드 라인에서 지정하는 플래그는 기본값과 해당 환경 변수를 무시한다.

<br>
<br>

### Basic Commands(Intermediate)

- `# kubectl get [name]` -- `pods`, `nodes`, `services` 들의 리소스 리스트를 불러온다.
  - `-o wide` -- 상세 정보

- `# kubectl edit [type] [name]` -- 리소스의 정보를 수정

- `# kubectl delete [COMMAND] [NAME]` -- 삭제

<br>

### Basic Commands (Beginner)

- `# kubectl expose (-f FILENAME \| TYPE NAME) [--port=port] ...[option]` -- 서비스를 생성
  - ex) `kubectl expose pod nginx-test` 
- `# kubectl create -f FILENAME [flags]` -- 리소스 file 생성

- `# kubectl run [pod_name] -- image [image] --port=[port] ... [option]` -- pod 실행, 생성
  - ex) `# kubectl run nginx-test --image=nginx --port 80 --generator=run-pod/v1`

- `# systemctl start kubelet` -- kubelet 재시작

<br>

### Troubleshooting and Debugging Commands

- `# kubectl exec POD [-c CONTAINER] [-i][-t] [flags] [– COMMAND [args…]]` -- pod 내부의 컨테이너에 명령어 날리기

- `# kubectl describe [type] [name]` -- 리소스 상태 조회
- `# kubectl logs -f [pod_name]` -- 로그 확인

- `# kubectl port-forward [type] [name] [port:port]` -- 파드 이름과 같이 리소스 이름을 사용하여 일치하는 파드를 선택해 포트 포워딩하는 것을 허용

<br>

### Advanced Commands

- `# kubectl apply -f [file_name, URL]` -- 존에 존재하는 것을 제외한, 지정한 디렉터리 내 구성 파일에 의해 정의된 모든 오브젝트를 생성
  - `apply`와 `create`의 사용 차이
    - 처음 생성은 `create`로 실행하고, 그 다음 업데이트가 발생할경우 `apply`를 하면 된다.
- `--record` : `rollout`의 history를 기록합니다.
- `# kubectl replace -f [file_name]` -- 설정 파일 수정하거나, 설정 파일을 새로 만들어서 그 파일로 설정을 업데이트
- `# kubectl rollout history <TYPE> <NAME>` -- `deployment`나 `replicaset`, `pod`들을 수정할때 그 기록이 남아 조회를 할 수 있습니다.
  - `--revision=<NUMBER>` : 특정 `history`의 내용을 조회합니다.
  - `kubectl rollout undo <TYPE> <NAME> --to-revision=<revision>` -- 특정 `revision`으로 수정을 합니다.
- `# kubectl annotate deployment/echo kubernetes.io/change-cause="<바꿀 이름>"` -- `history`의 이름을 자신이 알아보기 쉽게 원하는 이름으로 변경

<br>

### Other Commands

- `# kubectl api-versions [flags]` -- 사용가능한 API version 조회

- `# kubectl api-resources` -- [API 그룹](https://kubernetes.io/ko/docs/concepts/overview/kubernetes-api/#api-그룹)과 함께 지원되는 모든 리소스 유형들, 그것들의 [네임스페이스](https://kubernetes.io/ko/docs/concepts/overview/working-with-objects/namespaces)와 [종류(Kind)](https://kubernetes.io/ko/docs/concepts/overview/working-with-objects/kubernetes-objects)를 나열
  - API 리소스를 탐색하기 위한 다른 작업

```bash
kubectl api-resources --namespaced=true      # 네임스페이스를 가지는 모든 리소스
kubectl api-resources --namespaced=false     # 네임스페이스를 가지지 않는 모든 리소스
kubectl api-resources -o name                # 모든 리소스의 단순한 (리소스 이름 만) 출력
kubectl api-resources -o wide                # 모든 리소스의 확장된 ("wide"로 알려진) 출력
kubectl api-resources --verbs=list,get       # "list"와 "get"의 요청 동사를 지원하는 모든 리소스 출력
kubectl api-resources --api-group=extensions # "extensions" API 그룹의 모든 리소스
```

<br>
<br>

## Kubernetes 주요 개념

| Resource or Object      | 용도                                                         |
| ----------------------- | ------------------------------------------------------------ |
| Node                    | 컨테이너가 배치되는 서버                                     |
| Namespace               | 쿠버네티스 클러스터 안의 가상 클러스터                       |
| Pod                     | 컨테이너의 집합 중 가장 작은 단위, 컨테이너의 실행 방법 정의 |
| Replica Set             | 같은 스펙을 갖는 파드를 여러 개 생성하고 관리하는 역할       |
| Deployment              | 레플리카 세트의 리비전을 관리                                |
| Service                 | 파드의 집합에 접근하기 위한 경로를 정의                      |
| Ingress                 | 서비스를 쿠버네티스 클러스터 외부로 노출                     |
| ConfigMap               | 설정 정보를 정의하고 파드에 전달                             |
| Persistent Volume       | 파드가 사용할 스토리지의 크기 및 종류를 정의                 |
| Persistent Volume Claim | 퍼시스턴트 볼륨을 동적으로 확보                              |

<br>
<br>

# 리소스[RESOURCE] 종류

**kubectl에 적용 가능한 쿠버네티스 리소스 종류와 단축어 리스트 입니다.**

| 리소스 종류                | 단축어 |
| :------------------------- | :----- |
| apiservices                |        |
| certificatesigningrequests | csr    |
| clusters                   |        |
| clusterrolebindings        |        |
| clusterroles               |        |
| componentstatuses          | cs     |
| configmaps                 | cm     |
| controllerrevisions        |        |
| cronjobs                   |        |
| customresourcedefinition   | crd    |
| daemonsets                 | ds     |
| deployments                | deploy |
| endpoints                  | ep     |
| events                     | ev     |
| horizontalpodautoscalers   | hpa    |
| ingresses                  | ing    |
| jobs                       |        |
| limitranges                | limits |
| namespaces                 | ns     |
| networkpolicies            | netpol |
| nodes                      | no     |
| persistentvolumeclaims     | pvc    |
| persistentvolumes          | pv     |
| poddisruptionbudget        | pdb    |
| podpreset                  |        |
| pods                       | po     |
| podsecuritypolicies        | psp    |
| podtemplates               |        |
| replicasets                | rs     |
| replicationcontrollers     | rc     |
| resourcequotas             | quota  |
| rolebindings               |        |
| roles                      |        |
| secrets                    |        |
| serviceaccounts            | sa     |
| services                   | svc    |
| statefulsets               |        |
| storageclasses             |        |