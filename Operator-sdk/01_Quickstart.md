# Quickstart for Go-based Operators

### Install
 
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

### Create a new project

```shell
mkdir -p $HOME/projects/test-operator
cd $HOME/projects/test-operator

export GO111MODULE=on

# domain 을 `example.com` 으로 설정을 하면 API groups 가 `<group>.example.com`로 된다.
operator-sdk init --domain <example.com> --repo github.com/<example>/test-operator
```

`--domain`은 생성될 API group 의 접두사로 사용됩니다. API group 은 Kubernetes API의 일부를 그룹화하는 메커니즘입니다.
중요한 것은 RBAC를 사용하여 리소스 유형에 대한 액세스를 제어할 수 있는 방법을 결정하기 때문에 이해하기 쉽게 의미 있는 group 으로 
리소스 유형을 그룹화하도록 도메인 이름을 지정해야 합니다. 자세한 정보는 [Kubernetes 문서](https://kubernetes.io/docs/reference/using-api/#api-groups) 
및 [Kubebuilder 문서](https://book.kubebuilder.io/cronjob-tutorial/gvks.html)를 확인하세요.

**Note** 로컬 환경이 Apple Silicon(`darwin/arm64`)인 경우 `go/v4-alpha` init 명령어에 플래그 `--plugins=go/v4-alpha`를 추가하여 사용하세요.

<br>

Manager는 모든 컨트롤러가 리소스를 감시하는 네임스페이스를 제한할 수 있습니다.

기본적으로 이것은 연산자가 실행 중인 네임스페이스가 됩니다.
```go
mgr, err := ctrl.NewManager(cfg, manager.Options{Namespace: namespace})
```

모든 네임스페이스를 보려면 네임스페이스 옵션을 비워 둡니다.

```go
mgr, err := ctrl.NewManager(cfg, manager.Options{Namespace: ""})
```

<br>
<br>

### Create a new API and Controller

```shell
operator-sdk create api --group <example> --version v1alpha1 --kind Test --resource --controller
```

<br>

### Define the API 

시작하기에 앞서, 배포할 `Test` type 을 정의하여 API 를 나타냅니다.

```go
type TestSpec struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:ExclusiveMaximum=false
	Size int32 `json:"size,omitempty"`
}

// MemcachedStatus defines the observed state of Memcached
type TestStatus struct {
	// Represents the observations of a Memcached's current state.
	// Memcached.status.conditions.type are: "Available", "Progressing", and "Degraded"
	// Memcached.status.conditions.status are one of True, False, Unknown.
	// Memcached.status.conditions.reason the value should be a CamelCase string and producers of specific
	// condition types may define expected values and meanings for this field, and whether the values
	// are considered a guaranteed API.
	// Memcached.status.conditions.Message is a human readable message indicating details about the transition.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}
```
컨트롤러가 나머지 CR 개체를 변경하지 않고 CR 상태를 업데이트할 수 있도록 `+kubebuilder:subresource:status` [마커](https://book.kubebuilder.io/reference/generating-crd.html#status)를 
추가 하여 CRD 매니페스트에 [상태 하위 리소스 를 추가합니다.](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#status-subresource)

<br>

`*_types.go`파일을 수정한 후 다음 명령을 실행하여 해당 리소스 유형에 대해 생성된 코드를 업데이트합니다.

```shell
make generate
```

<br>

### Generating CRD manifests

```shell
make manifests
```

<br>
<br>

### Implement the Controller
