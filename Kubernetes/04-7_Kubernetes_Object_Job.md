# 잡(Job)

- 하나 이상의 파드를 생성해 지정된 수(parallelism)의 파드가 정상 종료될 때까지 이를 관리하는 리소스

- 잡이 생성한 파드는 정상 종료 후에도 삭제되지 않고 남아 있어 로그나 실행 결과를 분석할 수 있음

- 배치 작업 형태에 적합



### 생성 및 조회

- **`[vagrant@master ~]$ vi simple-job.yaml`**

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pingpong
  labels:
    app: pingpong
spec:
  parallelism: 3                        ⇐ 동시에 실행하는 포드의 수 (= 병렬로 실행)
  template:                             ⇐ 포드를 정의
    metadata:
      labels:
        app: pingpong
    spec:
      containers:
        - name: pingpong
          image: gihyodocker/alpine:bash
          command: ["/bin/bash"]  
          args: 
          - "-c"
          - |
            echo [`date`] ping!
            sleep 10
            echo [`date`] pong!
      restartPolicy: Never
```

- `restartPolicy` : 포드 종료 후 재실행 여부를 설정합니다.
  - 종류 : `Always`, `Never`, `OnFailure`
- `Pod`는 `Always`가 기본
- `Job`은 `Always`로 설정 불가, `Never`, `OnFailure`만 설정 가능



```sh
[vagrant@master ~]$ kubectl apply -f simple-job.yaml
job.batch/pingpong created

[vagrant@master ~]$ kubectl logs -l app=pingpong
[Wed Feb 24 10:35:37 UTC 2021] ping!	← (1) 37초
[Wed Feb 24 10:35:47 UTC 2021] pong!
[Wed Feb 24 10:35:39 UTC 2021] ping!	← (2) 39초
[Wed Feb 24 10:35:49 UTC 2021] pong!
[Wed Feb 24 10:35:53 UTC 2021] ping!	← (3) 53초
[Wed Feb 24 10:36:03 UTC 2021] pong!

[vagrant@master ~]$ kubectl get pods -l app=pingpong -o wide
NAME             READY   STATUS      RESTARTS   AGE     IP                NODE    ...
pingpong-mgb58   0/1     Completed   0          3m52s   192.168.104.1     node2   ...
pingpong-vdm7r   0/1     Completed   0          3m52s   192.168.104.2     node2   ...
pingpong-w9hs6   0/1     Completed   0          3m52s   192.168.166.137   node1   ...
```





# 크론잡(CronJob)

- `Job`은 한 번만 실행되는 반면, `CronJob`은 스케줄을 지정해 정기적으로 `Pod`를 실행

- `cron` 등을 사용해 정기적으로 실행하는 작업에 적합
- **스케줄 정의 시간**

```sh
*　　　　　　*　　　　　　*　　　　　　*　　　　　　*
분(0-59)　　시간(0-23)　　일(1-31)　　월(1-12)　　　요일(0-7)
```



### 생성 및 조회

- **`[vagrant@master ~]$ vi simple-cronjob.yaml`**

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: pingpong
spec:
  schedule: "*/1 * * * *"                    ⇐ 파드를 실행할 스케줄을 정의 : 60분/1 - 60초
  jobTemplate:                               ⇐ 파드를 정의
    spec:
      template:
        metadata:
          labels:
            app: pingpong
        spec:
          containers:
            - name: pingpong
              image: gihyodocker/alpine:bash
              command: ["/bin/bash"]
              args: 
              - "-c"
              - |
                echo [`date`] ping!
                sleep 10 
                echo [`date`] pong!
          restartPolicy: OnFailure
```



```sh
[vagrant@master ~]$ kubectl apply -f simple-cronjob.yaml
cronjob.batch/pingpong created

[vagrant@master ~]$ kubectl get job -l app=pingpong
NAME                  COMPLETIONS   DURATION   AGE
pingpong-1614164340   1/1           13s        2m7s		⇐ 약 60초 간격으로 실행
pingpong-1614164400   1/1           12s        66s
pingpong-1614164460   0/1           6s         6s

[vagrant@master ~]$ kubectl logs -l app=pingpong
[Wed Feb 24 10:57:52 UTC 2021] ping!
[Wed Feb 24 10:58:02 UTC 2021] pong!
[Wed Feb 24 10:58:53 UTC 2021] ping!
[Wed Feb 24 10:59:03 UTC 2021] pong!
[Wed Feb 24 10:59:52 UTC 2021] ping!
[Wed Feb 24 11:00:02 UTC 2021] pong!
```

