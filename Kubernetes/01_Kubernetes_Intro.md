# Kubernetes란?

쿠버네트스는 컨테이너 운영을 자동화하기 위한 컨테이너 오케스트레이션 도구로, 구글의 주도로 개발됐다. **많은 수의 컨테이너를 협조적으로 연동시키기 위한 통합 시스템**이며 이 컨테이너를 다루기 위한 API 및 명령행 도구 등이 함께 제공된다.

컨테이너를 이용한 애플리케이션 배포 외에도 다양한 운영 관리 업무를 자동화할 수 있다. 도커 호스트관리, 서버 리소스의 여유를 고려한 컨테이너 배치, 스케일링, 여러 개의 컨테이너 그룹에 대한 로드 밸런싱, 헬스 체크 등의 기능을 갖추고 있다.

**쿠버네트스의 가장 큰 특징은 다양한 부품을 조합해 유연한 애플리케이션을 구축할 수 있다는 점이다.**



# Kubernetes 명령어

터미널 창에서 `kubectl` 명령을 실행하려면 다음의 구문을 사용한다.

```
kubectl [command] [TYPE] [NAME] [flags]
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

  - 타입 및 이름으로 리소스를 지정하려면 다음을 참고한다.

    - 리소스가 모두 동일한 타입인 경우 리소스를 그룹화하려면 다음을 사용한다. 

      - ```
        TYPE1 name1 name2 name<#>
        ```

      - 예: `kubectl get pod example-pod1 example-pod2`

    - 여러 리소스 타입을 개별적으로 지정하려면 다음을 사용한다.

      -  ```
        TYPE1/name1 TYPE1/name2 TYPE2/name3 TYPE<#>/name<#>
         ```

      - 예: `kubectl get pod/example-pod1 replicationcontroller/example-rc1`

  - 하나 이상의 파일로 리소스를 지정하려면 다음을 사용한다. 

    - ```
      -f file1 -f file2 -f file<#>
      ```

    - YAML이 특히 구성 파일에 대해 더 사용자 친화적이므로, [JSON 대신 YAML을 사용한다](https://kubernetes.io/ko/docs/concepts/configuration/overview/#일반적인-구성-팁).
      예: `kubectl get -f ./pod.yaml`

- `flags`: 선택적 플래그를 지정한다. 예를 들어, `-s` 또는 `--server` 플래그를 사용하여 쿠버네티스 API 서버의 주소와 포트를 지정할 수 있다.

> **주의:** 커맨드 라인에서 지정하는 플래그는 기본값과 해당 환경 변수를 무시한다.



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

