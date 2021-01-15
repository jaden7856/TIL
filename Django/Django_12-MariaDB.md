# MariaDB와 Django 연동

#### 1. 사용자 계정 생성과 Database 생성

- Root 권한으로 사용자 계정 생성하기

```
mysql -u root –p
show databases;
use mysql;
create user python@localhost identified by ‘python';
grant all on *.* to python@localhost;
flush privileges;
exit;
```



- 사용자 계정으로 database 생성하기

```
mysql -u python -p
create database django_db;
show databases;
use django_db;
```



#### 2. Pymysql과 Mysqlclient 설치

- Django와 Mysql(MariaDB)을 연동하기 위해서는 `pymysql`과 `mysqlclient` 라는 패키지가 둘 다 필요합니다
  - `pip install mysqlclient`
  - `pip install pymysql`



#### 3. Django 프로젝트 설정 (settings.py) 수정하기

- mydjango/settings.py 파일에서 DATABASES 부분을 찾아서 아래와 같이 수정

```python
import pymysql

pymysql.version_info = (2, 0, 3, "final", 0)
pymysql.install_as_MySQLdb()

DATABASES = {
    # 'default': {
    # 'ENGINE': 'django.db.backends.sqlite3',
    # 'NAME': os.path.join(BASE_DIR, 'db.sqlite3'),
    # }
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': '<DB Name>', # DB명
        'USER': '<user>', # 데이터베이스 계정
        'PASSWORD':'<user password>', # 계정 비밀번호
        'HOST':'localhost', # 데이테베이스 IP
        'PORT':'<my_port_number>', # 데이터베이스 port
    }
}
```



- MariaDB에 마이그레이션 하고, 새로운 User를 생성한다.
  - `python manage.py migrate `
  - `python manage.py makemigrations <name>`
  - `python manage.py migrate <name>`
  - `python manage.py createsuperuser`