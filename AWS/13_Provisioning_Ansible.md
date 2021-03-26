# Provisioning 이란?

**프로비저닝**(provisioning)은 사용자의 요구에 맞게 시스템 자원을 할당, 배치, 배포해 두었다가 필요 시 시스템을 즉시 사용할 수 있는 상태로 미리 준비해 두는 것을 말한다. 서버 자원 프로비저닝, OS 프로비저닝, 소프트웨어 프로비저닝, 스토리지 프로비저닝, 계정 프로비저닝 등이 있다. 수동으로 처리하는 '수동 프로비저닝'과 자동화 툴을 이용해 처리하는 '자동 프로비저닝'이 있다.

크게보면 CI(Continuous Integration)/CD(Continuous Deploy) 라고 말한다.



# Ansible

WindowsOS나 MacOS를 사용한다면 제일 먼저 Vagrant를 설치하여 가상서버 Linux를 사용하겠습니다.

원하는 폴더 밑에서 cmd창을 열고 **`vagrant init`**을 하고 Vagrantfile에 밑의 코드를 입력하겠습니다.

```yaml
Vagrant.configure("2") do |config|
  config.vm.define:"ansible-server" do |cfg|
    cfg.vm.box = "centos/7"
    cfg.vm.provider:virtualbox do |vb|
        vb.name="Ansible-Server"
        vb.customize ["modifyvm", :id, "--cpus", 2]
        vb.customize ["modifyvm", :id, "--memory", 2048]
    end
    cfg.vm.host_name="ansible-server"
    cfg.vm.synced_folder ".", "/vagrant", disabled: true
    cfg.vm.network "public_network", ip: "172.20.10.10"
    cfg.vm.network "forwarded_port", guest: 22, host: 19210, auto_correct: false, id: "ssh"
    cfg.vm.network "forwarded_port", guest: 8080, host: 58080
  end
end
```



실행을 하고  그냥 완료가 될수도 있고 선택지를 줄 수도 있습니다. 선택지를 준다면 컴퓨터마다 다르겠지만 대부분 `1)`을 선택하시면 됩니다. 정확한 판별을 하고싶다면 google검색을 하시면 됩니다.

```cmd
> vagrant up
```



이제 내부로 접속해 보겠습니다.  저희가 위의 코드중에 `host_name="ansible-server"`이라는 HOST_NAME을 설정하였기 때문에 접속 할때 그 이름으로 접속하시면 됩니다.  자신이 원하는 이름을 설정하셔도 무방합니다.

```cmd
> vagrant ssh ansible-server
```



그 다음 필요한 설치와 확인을 하겠습니다.

```cmd
$ sudo yum install -y net-tools
$ sudo yum install -y epel-release
$ sudo yum install -y ansible
$ ansible --version
```





### 자동화

위에 설치를 수동으로 여러개를 설치하면 힘들기 때문에 자동화하기 위해 다른 Vagrantfile을 수정하겠습니다.

우선 접속했던 ansible-server에서 나오셔서 vagrant를 중지하겠습니다.

```cmd
$ exit
> vagrant halt
```



기존의 Vagrantfile에서 `bootstrap.sh`이 적혀있는 코드를 추가하겠습니다.

```yaml
Vagrant.configure("2") do |config|
  config.vm.define:"ansible-server" do |cfg|
			:
			:
    cfg.vm.provision "shell", path: "bootstrap.sh"
  end
end
```



그 폴더에 `bootstrap.sh`라는 파일을 생성 후 위에서 저희가 설치했던 명령어들을 적어 보겠습니다.

```sh
#! /user/bin/env bash

yum install -y net-tools
yum install -y epel-release
yum install -y ansible
```





### Node

Ansible을 정확히 사용하기 위해 Server부분을 만들었으니 Node를 만들어 보겠습니다.

위치는 Server코드에서 end로 끝나는 바로 밑에 추가하시면 됩니다.

Node를 최소 1개에서 2개는 만들어보기위해 고쳐야 할 부분은 체크 표시하겠습니다.

```yaml
config.vm.define:"ansible-node-1" do |cfg|					# "ansible-node-2"
    cfg.vm.box = "centos/7"
    cfg.vm.provider:virtualbox do |vb|
        vb.name="Ansible-Node-1"							# "Ansible-Node-2"
        vb.customize ["modifyvm", :id, "--cpus", 1]
        vb.customize ["modifyvm", :id, "--memory", 1048]
    end
    cfg.vm.host_name="ansible-node-1"						# "ansible-node-2"
    cfg.vm.synced_folder ".", "/vagrant", disabled: false
    cfg.vm.network "public_network", ip: "172.20.10.11"		# ip: "172.20.10.12"
    cfg.vm.network "forwarded_port", guest: 22, host: 19211, auto_correct: false, id: "ssh"														  	  # host: 19212
    cfg.vm.network "forwarded_port", guest: 80, host: 10080		# host: 20080
    cfg.vm.provision "shell", path: "bash_ssh_conf_4_CentOS.sh"
end
```



그 폴더에 `bash_ssh_conf_4_CentOS.sh`라는 파일을 생성 후 위에서 저희가 설치했던 명령어들을 적어 보겠습니다.

```sh
#! /usr/bin/snv bash

now=$(date +"%m_%d_%Y")
cp /etc/ssh/sshd_config /etc/ssh/sshd_config_$now.backup
sed -i -e 's/PasswordAuthentication no/PasswordAuthentication yes/g' /etc/ssh/sshd_config
systemctl restart sshd
```



그리고 `node`를 하나 더 추가하셔서 CentOS에서 Ubuntu로 바꾸고 만들어 보겠습니다.

```yaml
config.vm.define:"ansible-node-3" do |cfg|
    cfg.vm.box = "ubuntu/trusty64"
    cfg.vm.provider:virtualbox do |vb|
        vb.name="Ansible-Node-3"
        vb.customize ["modifyvm", :id, "--cpus", 1]
        vb.customize ["modifyvm", :id, "--memory", 1048]
    end
    cfg.vm.host_name="ansible-node-3"
    cfg.vm.synced_folder ".", "/vagrant", disabled: true
    cfg.vm.network "public_network", ip: "172.20.10.13"
    cfg.vm.network "forwarded_port", guest: 22, host: 19213, auto_correct: false, id: "ssh"
    cfg.vm.network "forwarded_port", guest: 80, host: 30080
  end
```



실행 및 확인

- 만약 ssh client가 있다면 그걸로 접속하셔도 됩니다.

```cmd
> vagrant up
> vagrant ssh ansible-server
```

```cmd
$ ansible --version
```



서로 연결이 잘 되어있는지 확인해 보겠습니다.

```cmd
server
$ ping 172.20.10.11
$ ping 172.20.10.12
$ ping 172.20.10.13

node
$ ping 172.20.10.10
```

만약 ip접속말고 이름으로 하고싶다면 `hosts`로 접속해서 코드를 추가합니다.

```cmd
$ sudo vi /etc/hosts

172.20.10.10 ansible-server
172.20.10.11 ansible-node-1
172.20.10.12 ansible-node-2
172.20.10.13 ansible-node-3
```





hosts로 접근하여 밑의 코드를 추가하겠습니다. 만약에 Node의 숫자가 더 많다면 그에 맞게 ip들을 추가 해 주시면 됩니다.

```cmd
$ sudo vi /etc/ansible/hosts #맨 밑으로 이동

[nginx]				# 그룹이름
172.20.10.11
172.20.10.12
172.20.10.13

[webserver]
172.20.10.11

[backup]
172.20.10.12
172.20.10.13

[ubuntu]
172.20.10.13
```



접속 할 때 마다 password 입력을 하지 않기 위해 다음 작업을 하겠습니다.

```cmd
$ ssh-keygen
$ ssh-copy-id root@ansible-node-1
$ ssh-copy-id root@ansible-node-2
$ ssh-copy-id root@ansible-node-3
$ ssh-copy-id vagrant@ansible-node-1
$ ssh-copy-id vagrant@ansible-node-2
$ ssh-copy-id vagrant@ansible-node-3
```

비밀번호 입력시 자신이 설정했던 비밀번호 'vagrant'를 입력



##### Error

만약 위에 코드를 할 때 `denied`에러가 뜬다면 `bash_ssh_conf_4_CentOS.sh`파일이 제대로 적용이 되지 않았을 때 입니다. 그러면 수동으로 적용이 필요합니다.

```cmd
$ sudo vi /etc/ssh/sshd_config
```

접속해서 `PasswordAuthentication` 부분을 찾아 `no`를 `yes`로 바꾸고 저장합니다.

```cmd
$ sudo systemctl restart sshd
```

다시 재시작을 하고 `ssh-copy-id`를 다시 해보시기 바랍니다.



Ubuntu로 만든 `node-3`는 root@ansible-node-3로 할때 비밀번호가 틀릴겁니다. 그래서 고쳐주도록 하겠습니다.

```cmd
$ sudo vi /etc/ssh/sshd_config
```

들어가서 `PermitRootLogin without-password`부분을 찾아 `without-password`를 지우고 `yes`로 바꿉니다.

```cmd
sudo /etc/init.d/ssh restart
```



#### 서버에서 node의 정보를 인증

```cmd
$ ansible all -m user -a "user=test1 password=1234"
$ ansible all -m user -a "user=test1 password=1234" -k
172.20.10.13 | CHANGED => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python"
    },
    "changed": true,
    "comment": "",
    "create_home": true,
    "group": 1002,
    "home": "/home/test1",
    "name": "test1",
    "password": "NOT_LOGGING_PASSWORD",
    "shell": "",
    "state": "present",
    "system": false,
    "uid": 1002
}
		:
```



#### Node 전체에 파일 보내기

txt파일을 아무거나 한개 만들어서 모든 Node에 복사하여 보내보도록 하겠습니다.

```cmd
$ vi test_server.txt
Hello, there.
```



```cmd
$ ansible all -m copy -a "src=./test_server.txt dest=/home/vagrant"
```

- `src` -- server단에서 보낼 파일의 위치와 이름

- `dest` -- node단에서 받을 위치



Node home에서 `ll`로 확인

```cmd
$ vagrant@ansible-node-3:~$ ls
test_server.txt
```





## Test

```cmd
$ ansible all -m ping
172.20.10.13 | SUCCESS => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python"
    },
    "changed": false,
    "ping": "pong"
}
172.20.10.12 | SUCCESS => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python"
    },
    "changed": false,
    "ping": "pong"
}
172.20.10.11 | SUCCESS => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python"
    },
    "changed": false,
    "ping": "pong"
}
```

```cmd
$ ansible all -m shell -a "uptime"		# ansible <group-name>
172.20.10.13 | CHANGED | rc=0 >>
 06:53:22 up 7 min,  2 users,  load average: 0.01, 0.20, 0.16
172.20.10.12 | CHANGED | rc=0 >>
 05:56:07 up  2:53,  1 user,  load average: 0.08, 0.03, 0.05
172.20.10.11 | CHANGED | rc=0 >>
 05:56:07 up  2:55,  1 user,  load average: 0.24, 0.06, 0.06
```



##### 용량 확인

```cmd
$ ansible all -m shell -a "df -h"
172.20.10.13 | CHANGED | rc=0 >>
Filesystem      Size  Used Avail Use% Mounted on
udev            504M   12K  504M   1% /dev
tmpfs           102M  372K  102M   1% /run
/dev/sda1        40G  1.5G   37G   4% /
none            4.0K     0  4.0K   0% /sys/fs/cgroup
none            5.0M     0  5.0M   0% /run/lock
none            508M     0  508M   0% /run/shm
none            100M     0  100M   0% /run/user
172.20.10.12 | CHANGED | rc=0 >>
Filesystem      Size  Used Avail Use% Mounted on
devtmpfs        500M     0  500M   0% /dev
tmpfs           507M     0  507M   0% /dev/shm
tmpfs           507M  6.8M  500M   2% /run
tmpfs           507M     0  507M   0% /sys/fs/cgroup
/dev/sda1        40G  3.0G   38G   8% /
tmpfs           102M     0  102M   0% /run/user/1000
172.20.10.11 | CHANGED | rc=0 >>
Filesystem      Size  Used Avail Use% Mounted on
devtmpfs        500M     0  500M   0% /dev
tmpfs           507M     0  507M   0% /dev/shm
tmpfs           507M  6.8M  500M   2% /run
tmpfs           507M     0  507M   0% /sys/fs/cgroup
/dev/sda1        40G  3.0G   38G   8% /
tmpfs           102M     0  102M   0% /run/user/1000
```



##### 메모리 확인

```cmd
$ ansible all -m shell -a "free -h"
172.20.10.13 | CHANGED | rc=0 >>
             total       used       free     shared    buffers     cached
Mem:          1.0G       396M       618M       384K        13M       249M
-/+ buffers/cache:       133M       881M
Swap:           0B         0B         0B
172.20.10.11 | CHANGED | rc=0 >>
              total        used        free      shared  buff/cache   available
Mem:           1.0G         96M        823M        6.8M         92M        796M
Swap:          2.0G          0B        2.0G
172.20.10.12 | CHANGED | rc=0 >>
              total        used        free      shared  buff/cache   available
Mem:           1.0G        102M        783M        6.8M        127M        773M
Swap:          2.0G          0B        2.0G
```
