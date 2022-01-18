# Tibero Database 란?

> 티맥스소프트사에서 제작한 한국산 DBMS이다. 오라클에 대한 대체품으로 개발하여 SQL 등이 오라클과 거의 유사하다.

> [참고 사항](https://technet.tmaxsoft.com/upload/download/online/tibero/pver-20140808-000002/tibero_admin/ch_03.html)



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





### 데이터 저장 구조

Tibero의 데이터를 저장하는 구조는 다음과 같이 두 가지 영역으로 나뉜다.

- 논리적 저장 영역

  - 데이터베이스의 스키마 객체를 저장하는 영역이다.

  - 논리적 저장 영역은 다음과 같은 포함 관계가 있다.

    ```
    데이터베이스 > 테이블 스페이스 > 세그먼트 > 익스텐트
    ```

- 물리적 저장 영역

  - 운영체제와 관련된 파일을 저장하는 영역이다.

  - 물리적 저장 영역은 다음과 같은 포함 관계가 있다.

    ```
    데이터 파일 > 운영체제의 데이터 블록
    ```





## 테이블 스페이스

**테이블 스페이스**는 논리적 저장 영역과 물리적 저장 영역에 공통적으로 포함된다. 논리적 저장 영역에는 Tibero의 모든 데이터가 저장되며, 물리적 저장 영역에는 데이터 파일이 하나 이상 저장된다.

테이블 스페이스는 논리적 저장 영역과 물리적 저장 영역을 연관시키기 위한 단위이다.

### 3.2.1. 테이블 스페이스 구성

테이블 스페이스는 크게 두 가지 구성으로 Tibero의 데이터를 저장한다.

#### 테이블 스페이스의 논리적 구성

다음은 테이블 스페이스의 논리적 구성을 나타내는 그림이다.



**[그림 3.1] 테이블 스페이스의 논리적 구성**

![테이블 스페이스의 논리적 구성](https://technet.tmaxsoft.com/upload/download/online/tibero/pver-20140808-000002/tibero_admin/resources/tbadmin_1.png)



테이블 스페이스는 [[그림 3.1\]](https://technet.tmaxsoft.com/upload/download/online/tibero/pver-20140808-000002/tibero_admin/ch_03.html#fig_tablespace_logical)과 같이 세그먼트(Segment), 익스텐트(Extent), 데이터 블록(Block)으로 구성된다.

|  구성요소   | 설명                                                         |
| :---------: | :----------------------------------------------------------- |
|  세그먼트   | 익스텐트의 집합이다.하나의 테이블, 인덱스 등에 대응되는 것으로 CREATE TABLE 등의 문장을 실행하면 생성된다. |
|  익스텐트   | 연속된 데이터 블록의 집합이다.세그먼트를 처음 만들거나 세그먼트의 저장 공간이 더 필요한 경우 Tibero는 테이블 스페이스에서 연속된 블록의 주소를 갖는 데이터 블록을 할당받아 세그먼트에 추가한다. |
| 데이터 블록 | 데이터베이스에서 사용하는 데이터의 최소 단위이다.Tibero는 데이터를 블록(Block) 단위로 저장하고 관리한다. |

### 참고

논리적 저장 영역을 관리하는 방법에 대한 자세한 내용은 [“제4장 스키마 객체 관리”](https://technet.tmaxsoft.com/upload/download/online/tibero/pver-20140808-000002/tibero_admin/ch_04.html)를 참고한다.

#### 테이블 스페이스의 물리적 구성

다음은 테이블 스페이스의 물리적 구성을 나타내는 그림이다.



**[그림 3.2] 테이블 스페이스의 물리적 구성**

![테이블 스페이스의 물리적 구성](https://technet.tmaxsoft.com/upload/download/online/tibero/pver-20140808-000002/tibero_admin/resources/tbadmin_2.png)



테이블 스페이스는 [[그림 3.2\]](https://technet.tmaxsoft.com/upload/download/online/tibero/pver-20140808-000002/tibero_admin/ch_03.html#fig_tablespace_physical)와 같이 물리적으로 여러 개의 데이터 파일로 구성된다. Tibero는 데이터 파일 외에도 컨트롤 파일과 로그 파일을 이용하여 데이터를 저장할 수 있다.

빈번하게 사용되는 두 테이블 스페이스(예: 테이블과 인덱스)는 물리적으로 서로 다른 디스크에 저장하는 것이 좋다. 왜냐하면 한 테이블 스페이스를 액세스하는 동안에 디스크의 헤드가 그 테이블 스페이스에 고정되어 있기 때문에 다른 테이블 스페이스를 액세스할 수 없다.

따라서 서로 다른 디스크에 각각의 테이블 스페이스를 저장하여 동시에 액세스하는 것이 데이터베이스 성능을 향상시키는 데 도움이 된다.





### 구조적

![img](https://media.vlpt.us/images/sysop/post/4068f2e1-2a8a-44ed-ac03-ba629a75e997/image.png)
![img](https://media.vlpt.us/images/sysop/post/6113f907-b7ef-44d8-a23b-9e81645a395d/image.png)
![img](https://media.vlpt.us/images/sysop/post/708e8971-89ed-4870-9961-32d96a5e251b/image.png)