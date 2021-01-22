# Python에서 MariaDB를 연동하여 사용하기
**MySQL Client에 접속**

- `mysql -u python -p`를 치고 자신의 비밀번호 입력
- `show databases;`  
- `use <user name>;`



## pymysql과 sqlalchemy 연동

``` import pymysql
import pymysql
import sqlalchemy

# pymysql과 sqlalchemy 연동
pymysql.install_as_MySQLdb()
from sqlalchemy import create_engine

# Engine 객체 생성
# 주의점!)localhost:3307은 자신이 MariaDB설치할때 선택했던 port번호로 적는다.
		 python:python@localhost는 자신의 아이디로 적는다.
engine = create_engine('mysql+mysqldb://python:python@localhost:3307/python_db',\
                      encoding='utf-8')

# Engine을 사용해서 db에 연결
con = engine.connect()
```

 ```df.to_sql(name='<name>', con=engine, if_exists='replace', index=False)```

- DataFrame `to_sql()` 함수로 dataframe객체를 table로 저장

