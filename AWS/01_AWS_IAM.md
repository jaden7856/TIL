

# Introduction to AWS Identity and Access Management (IAM)

### AWS IAM

- 자격을 인증하고 권한을 부여
- AWS 고객이 AWS 계정 및 AWS 내에서 사용 가능한 API 및 서비스에 대한 사용자의 액세스 및 권한을 관리할 수 있도록 하는 서비스
- 사용자, 보안 자격 증명(예: API Access Key)을 관리하고 사용자가 AWS 리소스에 액세스할 수 있도록 허용할 수 있음



### IAM 정책

![img](01_AWS_IAM.assets/xY98axjYWBFt2g8vrDXEpq-IRQuwUzil_gn6kCmnnVUIaAKIhzGUkr1PbTYDbID9Mh6OBB_MSCSRb7kR8eFxCKDWl1R5V_uGhkoJ8czuMomXHZhHNdTFUgQQsTCaYkY2M4gvb6Rh)

**S3-Support Group에 적용되어 있는 정책**

```
{
  "Version": "2012-10-17",
  "Statement": [					⇐ 정책 문서
    {
      "Effect": "Allow",			⇐ 효과 = 리소스에 대한 작업의 허용 여부를 명시
      "Action": [					⇐ 작업
        "s3:Get*",
        "s3:List*"
      ],
      "Resource": "*"				⇐ 자원
    }
  ]
}
```



**2-Support Group에 적용되어 있는 정책**

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "ec2:Describe*",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": "elasticloadbalancing:Describe*",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "cloudwatch:ListMetrics",
        "cloudwatch:GetMetricStatistics",
        "cloudwatch:Describe*"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": "autoscaling:Describe*",
      "Resource": "*"
    }
  ]
}
```



**2-Admin Group에 적용되어 있는 정책**

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*",
        "ec2:StartInstances",
        "ec2:StopInstances"
      ],
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "elasticloadbalancing:Describe*",
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": [
        "cloudwatch:ListMetrics",
        "cloudwatch:GetMetricStatistics",
        "cloudwatch:Describe*"
      ],
      "Resource": "*",
      "Effect": "Allow"
    },
    {
      "Action": "autoscaling:Describe*",
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
```



## 추가 LAB

### 1. `S3-Support Group`에 EC2 인스턴스를 실행하고, 중지할 수 있도록 권한을 추가합니다.

   1) `EC-Admin` 그룹의 `ec2-admin` (인라인 정책) 내용을 복사

   2) `S3-Support` 그룹에 인라인 정책으로 추가

![img](01_AWS_IAM.assets/AYvLrTz8zNWcWYJ3w3qidd_bbk90sPQLCm-VQekVCwtDmYs5RvDypKR0FB-4ZQlHxFDwqH6fQCMNm-SKQqsKi01M4JXi90MJuxSr9F7eMtgCzWRCLJHPASDX63x4tKzPH-Fa0H6R)

![img](01_AWS_IAM.assets/tuBCn0qzwQDtk7bqScS79fJ_nioAONOKfVL6gDDVMnNrcmEqMqi8dw9keKI7FFMFFBlc5hO_LALkseIk860uQC2FYgJkaV834N_NKxFmos0Jc5laznVrA7z2NUNaWtfgjwYwSon6)



### 2. `user-1` 사용자가 EC2 인스턴스를 조회할 수 있도록 권한을 추가

![img](01_AWS_IAM.assets/iSFvclD4QPzPy1mffdxYs6JYj_NLGjnTkArbPNBqctuB5TJeDHhXQq_bgfagY96wGguSeIMWL6F6IotE6Tth1y1Du04S8TUrtRRVM7F1ije1WO1JmWBon-4QmxNxlw1B-zUzvZDS)



![img](01_AWS_IAM.assets/1AqyRDZfB4wKlXG1RqejeBYpkfJl8HYmsQRQLq5aopUT8WEjJ8E5x12ZHlacrES1PTtD_IXghbwbsqMM7pVWK6P6L7aR9bHiKm3bfbUVY--m7guoeUlxhL7fwfiplK5Ko-KEG4qC)



### 3. EC2 인스턴스를 실행하고 중지할 수 있는 관리형 정책 `ec2-manager` 생성 후 기존 `ec2-admin` 인라인 정책 대신에 적용

![img](01_AWS_IAM.assets/4uuk06_BVx_FNgyvhJSO5wXL7pkC21sFcE7uhsvdXJ8tAnzKluyNniCILEqcgbXU9Nm4s8Q_CEN2_A4lUl7gI1eTmhwVlbwiE0Hz4Oa-WaglBHBb5cnF1r5ZS29xtDJ5qstbnA-o)

![](01_AWS_IAM.assets/ZvSjLB9Q9o7JOc3Wp0lAw593IxDeN8G-KQ1adrIon-Vg4k-2hfVVtMJQWkVvKum2eZy7wHs1WDdDT5PQxpWW7nMaOhoXsy5h6ejVEr4zpk2_OyL1ICejoa4NnIxTrhII5GOMJpK2)



- 정책을 생성합니다.

![](01_AWS_IAM.assets/PGiiH4TunYhwGzc2PYx9dOiZQ5psiuyhV8XQvtrURZ_Skx3nH1mcFD7x9PAGq7aMmqcAPlGUzqtb_Vp5XNN1VGgIk3yv_JE3jfhjf_xRNvI3iO347eDdGLp7mzdWtGiK9LADq8af)



- 정책을 설정합니다.

![img](01_AWS_IAM.assets/5UEYoBHltiEpgIfQc7tA5YceWDZWE8vXpBS0HeIy6emGsoOcTunCunrM_Ob3FnsejAZo9_xYYcz12IOs-YzUR32y-1FBFY-n_TNyalKQP-oCi7C8HYEVB5y62nGvbeSNu8xSFXAF)



### `ec2-manager` 정책에서 중지 권한 삭제 후 `user-1`(S3-Support 그룹), `user-3`(EC2-Admin 그룹) 계정으로 EC2 인스턴스 중지 가능 여부 확인

![img](01_AWS_IAM.assets/risHLBuJM9iWt6S4zOGJ8FsB-zveAyyDToz8MhcV1NcWDnu_0OfsWtqLMzbBXn9IWNxMxFAWkfI61TyUo4ZAK0rj-wIdb0XrVepuxIkQRsXOepFRt_01hpEaCxUwLXPJPctZr_lE)



- 중지 권한이 없기때문에 에러가 발생합니다.

![img](01_AWS_IAM.assets/0VnQ_mlaeaBqfHdtcJv3biH5ssoRI6k0hpyJU9qCMBboBHbgPSaG0_FEoe18yEeXFfBBKQb8txvoQcRAbjg7oIqLVxMEbEPcC-5q5RA8yVEPIYJgO-oPnBncY-GklEwI2HG6NXs7)