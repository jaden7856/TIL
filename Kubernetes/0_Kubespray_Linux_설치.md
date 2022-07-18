## Kubespray

- [Kubespray](#kubespray)
- [1. 개요](#1-개요)
- \2. 준비
  - [2-1. node](#2-1-node)
- \3. kubernetes 설치
  - [3-1. host 등록](#3-1-host-등록)
  - [3-2. ssh key 배포](#3-2-ssh-key-배포)
  - [3-3. ip forward 활성화](#3-3-ip-forward-활성화)
  - [3-4. firewalld 중지](#3-4-firewalld-중지)
  - [3-5. SELINUX permissive 변경](#3-5-selinux-permissive-변경)
  - [3-6. swap memory 사용 중지](#3-6-swap-memory-사용-중지)
  - [3-7. yum repo 등록](#3-7-yum-repo-등록)
  - [3-8. util install-1](#3-8-util-install-1)
  - [3-9. util install-2](#3-9-util-install-2)
- \4. kubespray
  - [4-1. kubespray git repo](#4-1-kubespray-git-repo)
  - [4-2. install dependencies](#4-2-install-dependencies)
  - [4-3. update ansible inventory file](#4-3-update-ansible-inventory-file)
  - [4-4. review and change parameters](#4-4-review-and-change-parameters)
- \5. kubernetes deploy
  - [5-1. ansible deploy](#5-1-ansible-deploy)
  - [5-2. kubenetes 상태 확인](#5-2-kubenetes-상태-확인)
- [6. 출처](#6-출처)

<br>
<br>

## 1. 개요

- HA 구성을 위한 3개 노드에 설치

<br>

## 2. 준비

### 2-1. node

- kubespray : 배포용 서버 (192.168.x.x)
- k8s1 : k8s1 (192.168.x.x)
- k8s2 : k8s2 (192.168.x.x)
- k8s3 : k8s3 (192.168.x.x)

<br>
<br>

## 3. kubernetes 설치

### 3-1. host 등록

- `kubespray node`

```
[root]# cat /etc/hosts
192.168.x.x k8s1
192.168.x.x k8s2
192.168.x.x k8s3
192.168.x.x kubespray
```
<br>

### 3-2. ssh key 배포

- `kubespray node`

```
[root]# ssh-keygen
... 그냥 엔터 

[root]# ssh-copy-id k8s1, k8s2, k8s3
```

<br>

### 3-3. ip forward 활성화

- `k8s node`

```
[root]# echo 1 > /proc/sys/net/ipv4/ip_forward
```

<br>

### 3-4. firewalld 중지

- `k8s node`

```
### kubenetes network에 대한 전반적인 이해 이전 임시 방편
[root]# systemctl stop firewalld && systemctl disable firewalld
```

<br>

### 3-5. SELINUX permissive 변경

- `k8s node`

```
[root]# sed -i 's/^SELINUX=.*/SELINUX=disabled/g' /etc/sysconfig/selinux && cat /etc/sysconfig/selinux; setenforce 0

# This file controls the state of SELinux on the system.
# SELINUX= can take one of these three values:
#     enforcing - SELinux security policy is enforced.
#     permissive - SELinux prints warnings instead of enforcing.
#     disabled - No SELinux policy is loaded.
SELINUX=disabled
# SELINUXTYPE= can take one of three values:
#     targeted - Targeted processes are protected,
#     minimum - Modification of targeted policy. Only selected processes are protected.
#     mls - Multi Level Security protection.
SELINUXTYPE=targeted

[root]# setenforce 0
```

<br>

### 3-6. swap memory 사용 중지

- `k8s node`

```
### kubenetes 성능 이슈로 인한 설치 요건
[root]# swapoff -a
```

<br>

### 3-7. yum repo 등록

- `kubespray node`

```
[root]# cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
exclude=kube*
EOF
```

<br>

### 3-8. util install-1

- `all node`

```
### 패키지 최신 업데이트 (선택)
[root]# yum update -y
[root]# reboot
```

- `kubespray node`

```
[### Python38 저장소 설치
[root]# yum install -y centos-release-scl

### Python38 설치
[root]# yum install -y rh-python38 rh-python38-python-pip which wget git
[root]# scl enable rh-python38 bash
[root]# vi ~/.bash_profile
...
scl enable rh-python38 bash
...
[root]# python --version
[root]# python3 --version

### 설치
[root]# python3 -m pip install --upgrade pip
[root]# pip install ansible                   ### ansible 6.0.0
[root]# pip install netaddr jinja2            ### netaddr 0.8.0 / jinja2 3.1.2

```

<br>

### 3-9. util install-2 (Ansible5 이하에서 설치시)

- `kubespray node`

```
[root]# python3 -m pip install --upgrade pip
[root]# pip3 install ansible                    ### ansible 4.10.0
[root]# pip3 install netaddr jinja2            ### netaddr 0.8.0 / jinja2 3.0.3

# 아래는 이전에 2.10.0으로 해야한다고 되어있었는데 최신 버전으로 설치해도 문제없었음. 참고만 하세요.
[root]# pip3.6 install ansible==2.10.0         ### kubespray 배포 시 "Check 2.9.0 <= Ansible version < 2.11.0"
```

<br>

<br>

## 4. kubespray

### 4-1. kubespray git repo

- `kubespray node`

```
[root]# git clone https://github.com/kubernetes-sigs/kubespray
```

<br>

### 4-2. install dependencies

- `kubespray node:/{kubespray_git_download_path}/`

```
[root]# pip3 install -r requirements.txt
```

<br>

### 4-3. update ansible inventory file

- `kubespray node:/{kubespray_git_download_path}/`

```
[root]# cp -r inventory/sample inventory/mycluster
[root]# KUBE_CONTROL_HOSTS=3 CONFIG_FILE=inventory/mycluster/hosts.yaml python3 contrib/inventory_builder/inventory.py k8s1,192.168.x.1 k8s2,192.168.x.2 k8s3,192.168.x.3

### KUBE_CONTROL_HOSTS : kube_control_plane 개수. default : 2
## 생성 결과 확인

[root]# cat inventory/mycluster/hosts.yaml
all:
  hosts:
    k8s1:
      ansible_host: 192.168.x.1
      ip: 192.168.x.1
      access_ip: 192.168.x.1
    k8s2:
      ansible_host: 192.168.x.2
      ip: 192.168.1.2
      access_ip: 192.168.x.2
    k8s3:
      ansible_host: 192.168.x.3
      ip: 192.168.x.3
      access_ip: 192.168.x.3
  children:
    kube_control_plane:
      hosts:
        k8s1:
        k8s2:
        k8s3:
    kube_node:
      hosts:
        k8s1:
        k8s2:
        k8s3:
    etcd:
      hosts:
        k8s1:
        k8s2:
        k8s3:
    k8s_cluster:
      children:
        kube_control_plane:
        kube_node:
    calico_rr:
      hosts: {}
```

<br>

### 4-4. review and change parameters (필요시)

- `kubespray node:/{kubespray_git_download_path}/`

```
[root]# vi inventory/mycluster/group_vars/k8s_cluster/addons.yml

# Helm deployment    ## helm chart 사용 위해
helm_enabled: true

# Metrics Server deployment      ## 리소스 모니터링
metrics_server_enabled: false
metrics_server_kubelet_insecure_tls: false

# Nginx ingress controller deployment     ## 파악 못함
ingress_nginx_enabled: true
```
```
[root]# vi inventory/mycluster/group_vars/k8s_cluster/k8s-cluster.yml

# Choose network plugin (cilium, calico, weave or flannel. Use cni for generic cni plugin)
# Can also be set to 'cloud', which lets the cloud provider setup appropriate routing
kube_network_plugin: calico      

# Setting multi_networking to true will install Multus: https://github.com/intel/multus-cni
kube_network_plugin_multus: true   

# Kube-proxy proxyMode configuration.
# Can be ipvs, iptables
kube_proxy_mode: ipvs      ## weave proxy mode

# Kubernetes cluster name, also will be used as DNS domain
cluster_name: kube     ## option
```

<br>

<br>

## 5. kubernetes deploy

### 5-1. ansible deploy

- `kubespray node:/{kubespray_git_download_path}/`

```
[root]# ansible-playbook -i inventory/mycluster/hosts.yaml --become --become-user=root cluster.yml

PLAY [localhost] *****************************************************************************************************************************************************************************************
Wednesday 14 July 2021  21:53:32 -0400 (0:00:00.043)       0:00:00.043 ********

TASK [Check 2.9.0 <= Ansible version < 2.11.0] ***********************************************************************************************************************************************************
ok: [localhost] => {
    "changed": false,
    "msg": "All assertions passed"
}
Wednesday 14 July 2021  21:53:32 -0400 (0:00:00.073)       0:00:00.116 ********

TASK [Check Ansible version > 2.10.11 when using ansible 2.10] *******************************************************************************************************************************************
ok: [localhost] => {
    "changed": false,
    "msg": "All assertions passed"
}
Wednesday 14 July 2021  21:53:32 -0400 (0:00:00.060)       0:00:00.177 ********

TASK [Check that python netaddr is installed] ************************************************************************************************************************************************************
ok: [localhost] => {
    "changed": false,
    "msg": "All assertions passed"
}
Wednesday 14 July 2021  21:53:33 -0400 (0:00:00.142)       0:00:00.320 ********

TASK [Check that jinja is not too old (install via pip)] *************************************************************************************************************************************************
ok: [localhost] => {
    "changed": false,
    "msg": "All assertions passed"
}

...
...
... (중략)

PLAY RECAP ***********************************************************************************************************************************************************************************************
k8s1                       : ok=632  changed=131  unreachable=0    failed=0    skipped=1150 rescued=0    ignored=1
k8s2                       : ok=564  changed=117  unreachable=0    failed=0    skipped=1010 rescued=0    ignored=0
k8s3                       : ok=566  changed=118  unreachable=0    failed=0    skipped=1008 rescued=0    ignored=0
localhost                  : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

Wednesday 14 July 2021  22:13:03 -0400 (0:00:00.093)       0:19:30.277 ********
===============================================================================
container-engine/docker : ensure docker packages are installed ----------------------------------------------------------------------------------------------------------------------------------- 67.87s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 54.41s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 47.85s
download : download_file | Download item --------------------------------------------------------------------------------------------------------------------------------------------------------- 46.84s
kubernetes/control-plane : Joining control plane node to the cluster. ---------------------------------------------------------------------------------------------------------------------------- 35.11s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 33.60s
download : download_file | Download item --------------------------------------------------------------------------------------------------------------------------------------------------------- 32.17s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 28.72s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 26.75s
kubernetes/control-plane : kubeadm | Initialize first master ------------------------------------------------------------------------------------------------------------------------------------- 26.21s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 25.11s
kubernetes/preinstall : Install packages requirements -------------------------------------------------------------------------------------------------------------------------------------------- 24.79s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 20.30s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 17.86s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 16.02s
download : download_container | Download image if required --------------------------------------------------------------------------------------------------------------------------------------- 15.86s
download : download_file | Download item --------------------------------------------------------------------------------------------------------------------------------------------------------- 14.77s
etcd : Gen_certs | Write etcd member and admin certs to other etcd nodes ------------------------------------------------------------------------------------------------------------------------- 14.75s
etcd : Gen_certs | Write etcd member and admin certs to other etcd nodes ------------------------------------------------------------------------------------------------------------------------- 14.29s
download : download_file | Download item --------------------------------------------------------------------------------------------------------------------------------------------------------- 13.14s
```

<br>

### 5-2. kubenetes 상태 확인

- `k8s node`

```
[root]# kubectl get pods --all-namespaces
NAMESPACE       NAME                              READY   STATUS    RESTARTS   AGE
ingress-nginx   ingress-nginx-controller-5qlrz    1/1     Running   0          4m13s
ingress-nginx   ingress-nginx-controller-dw77v    1/1     Running   0          4m13s
ingress-nginx   ingress-nginx-controller-jxp64    1/1     Running   0          4m13s
kube-system     coredns-8474476ff8-d4bgs          1/1     Running   0          3m49s
kube-system     coredns-8474476ff8-m6hj6          1/1     Running   0          3m43s
kube-system     dns-autoscaler-7df78bfcfb-thzf8   1/1     Running   0          3m45s
kube-system     kube-apiserver-k8s1               1/1     Running   0          6m33s
kube-system     kube-apiserver-k8s2               1/1     Running   0          6m4s
kube-system     kube-apiserver-k8s3               1/1     Running   0          5m49s
kube-system     kube-controller-manager-k8s1      1/1     Running   1          6m24s
kube-system     kube-controller-manager-k8s2      1/1     Running   0          6m4s
kube-system     kube-controller-manager-k8s3      1/1     Running   0          5m49s
kube-system     kube-multus-ds-amd64-sqgdb        1/1     Running   0          4m31s
kube-system     kube-multus-ds-amd64-wfss9        1/1     Running   0          4m31s
kube-system     kube-multus-ds-amd64-whhmc        1/1     Running   0          4m31s
kube-system     kube-proxy-6zc8k                  1/1     Running   0          4m54s
kube-system     kube-proxy-gf2rw                  1/1     Running   0          4m55s
kube-system     kube-proxy-vhgt4                  1/1     Running   0          4m54s
kube-system     kube-scheduler-k8s1               1/1     Running   0          6m33s
kube-system     kube-scheduler-k8s2               1/1     Running   0          6m4s
kube-system     kube-scheduler-k8s3               1/1     Running   1          5m50s
kube-system     metrics-server-86f4df7ff-lkhp6    1/2     Running   3          3m16s
kube-system     nodelocaldns-fv25k                1/1     Running   0          3m43s
kube-system     nodelocaldns-p97nm                1/1     Running   0          3m43s
kube-system     nodelocaldns-twr2f                1/1     Running   0          3m43s
kube-system     calico-net-9l9cb                  2/2     Running   1          4m40s
kube-system     calico-net-jgj8h                  2/2     Running   1          4m40s
kube-system     calico-net-swwn9                  2/2     Running   1          4m40s
```

```
[root]# kubectl get nodes --all-namespaces
NAME   STATUS   ROLES                  AGE     VERSION
k8s1   Ready    control-plane,master   7m45s   v1.21.2
k8s2   Ready    control-plane,master   7m14s   v1.21.2
k8s3   Ready    control-plane,master   7m      v1.21.2
```

```
[root]# kubectl get deployment --all-namespaces
NAMESPACE     NAME             READY   UP-TO-DATE   AVAILABLE   AGE
kube-system   coredns          2/2     2            2           5m15s
kube-system   dns-autoscaler   1/1     1            1           5m13s
kube-system   metrics-server   1/1     1            1           4m42s
```

<br>

<br>

## 6. 출처

- https://kubernetes.io/docs/setup/production-environment/tools/kubespray/
- https://github.com/kubernetes-sigs/kubespray
- https://waspro.tistory.com/558
- https://daehancni.tistory.com/8