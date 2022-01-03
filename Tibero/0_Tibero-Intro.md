# Tibero Database 란?

> 티맥스소프트사에서 제작한 한국산 DBMS이다. 오라클에 대한 대체품으로 개발하여 SQL 등이 오라클과 거의 유사하다.



### 주요기능

- 티베로 액티브 클러스터(TAC)
  - TAC는 Oracle의 RAC와 유사한 기능으로 미션 크리티컬 업무의 클러스터링 환경에서 성능 및 시스템 가용성 (HA을 위한 기능)
- 병령쿼리 지원
- 테이블 파티션 기능
- 데이터 압축 기능
- 데이터 암호화 기능
- 업무별 엔진그룹 분할 적용 기능
- 데이터베이스 링크 (DB Link)
- 온라인 오브젝트 재구성
- 접근 제어 및 감사
- 통계정보를 통한 성능튜닝
- 백업 및 복구
- 다양한 유틸리티
- 재해복구 시스템



### 장점

- 멀티 쓰래드, 멀티프로세스 기반의 고성능 구조
- 다양한 백업 및 고가용성 기반의 안정성 향상
- 표준 SQL 준수 및 표준 인터페이스를 통한 개발 편의성 제공
- 다양한 운영관리 유틸리티를 제공하여 마이그레이션 및 운영 편의성 증진



### 성능

- Thread Architecture를 채택하여 다중 사용자 성능 극대화
- Row Level Locking 기술을 통해 보다 많은 사용자 수용
- 다중버전 동시성 제어(MVCC: Multi-Version Concurrency Control) 기술을 통한 다중 사용자 접속 동시 처리 성능 향상
- Parallel Query를 통한 대용량 데이터 처리
- 다양한 인덱스 및 파티셔닝 기법 지원(Range, List, Hash, Composite partition)
- 비용 기반 질의처리 최적화(CBO), 다양한 Hint 지원



### 호환성

- 표준 SQL 및 표준 인터페이스 지원
- 이기종 DBMS의 PL/SQL, Embedded SQL, Data Type, Function 호환
- Database Link를 통한 다양한 이기종 DBMS와의 연동
- 자동화된 Migration 툴 지원



### 안전성

- 공유 디스크 기반 Active Cluster를 통한 이중화 제공
- 다양한 백업 및 복구 기능 지원
- 다양한 고가용성 아키텍처를 통한 이증화 선택 유연성 제공
- TAC(Tibero Active Cluster), TSC(Tibero Standby Cluster)
- Parallel Query를 통한 대용량 데이터 처리
- 다양한 인덱스 및 파티셔닝 기법 지원(Range, List, Hash, Composite partition)
- 비용 기반 질의처리 최적화(CBO), 다양한 Hint 지원



### 구조적

![img](https://media.vlpt.us/images/sysop/post/4068f2e1-2a8a-44ed-ac03-ba629a75e997/image.png)
![img](https://media.vlpt.us/images/sysop/post/6113f907-b7ef-44d8-a23b-9e81645a395d/image.png)
![img](https://media.vlpt.us/images/sysop/post/708e8971-89ed-4870-9961-32d96a5e251b/image.png)