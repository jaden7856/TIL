# AWS

- 특징
  - 확장성 → Auto Scaling 서비스 
  - 탄력성 → 수요가 떨어졌을 때 용량을 자동으로 줄이는 개념 ⇒ 비용 효율적
  - 비용 관리 → 자본지출(CAPEX)에서 운영비용(OPEX)으로 IT 지출 내용을 변경 ⇒ 장기적 수요를 예측 → 위험을 감수



# AWS 클라우드 서비스 범위

![img](00_AWS.assets/lhYOEIYbryDOE5dGOt8wrJpOddAc-KKDHyXDxsHwiEhRsybO5VndGLrzNtzCalr8igyUprjA0PJTX0YoXUksoqpIPcdWjLKVGNHGON_cDR66rKpz5HqflzHdnvj4N7tNAlw6ScGB)

### 컴퓨팅 - 물리 서버가 하는 역할을 복제한 클라우드 서비스

- EC2(Elastic Compute Cloud)
- Lambda
- Auto Scaling
- Elastic Load Balancing
- Elastic Beanstalk



### 네트워킹 - 어플리케이션 연결, 접근 제어, 원접 연결, … 

- VPC(Virtual Private Cloud)
- Direct Connect
- Route 53
- CloudFront



### 스토리지 - 빠른 액세스, 장기 백업과 같은 요구를 충족하는 스토리지 플랫폼

- S3(Simple Storage Service)
- Glacier
- EBS(Elastic Block Store)
- Storage Gateway



### 데이터베이스 - 관계형, NoSQL, 캐싱 등

- RDS(Relational Database Service)
- DynamoDB



### 어플리케이션 관리

- CloudWatch
- CloudFormation
- CloudTrail
- Config



### 보안과 자격 증명

- IAM(Identity and Access Management)
- KMS(Key Management Service)
- Directory Service



### 어플리케이션 통합

- SNS(Simple Notification Service)
- Simple WorkFlow(SWF)
- SQS(Simple Queue Service)
- API Gateway