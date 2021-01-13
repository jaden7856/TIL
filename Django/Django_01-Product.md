# Django 프로젝트 생성

- `terminal> django-admin startproject <name> .`
-  django/conf/project_template 구성으로 생성 됨 
  - manage.py : 웹사이트 관리를 도와주는 역할을 하는 파일
  - settings.py : 웹사이트 설정이 있는 파일
  - urls.py : urlresolver가 사용하는 요청 패턴(URL규칙) 목록을 포함하고 있는 파
  - wsgi.py : Web Server Gateway Interface이며 Python의 표준 Gateway Interface 입니다. 
  - asgi.py : Asynchronous Server Gateway Interface WSGI와 비슷한 구조를 가지며, 비동기 통신을 지원한다.



# Django 프로젝트 설정 변경

- settings.py에서 LANGUAGE_CODE, TIME_ZONE 함수에 값 변경
  
  - `LANGUAGE_CODE = 'ko'`
  - `TIME_ZONE = 'Asia/Seoul'`
  
- settings.py에서 정적 파일 경로를 추가함 `STATIC_URL`항목 바로 아래에 `STATIC_ROOT`을 추가

  ```python
  import os 
  STATIC_URL = '/static/' 
  STATIC_ROOT = os.path.join(BASE_DIR, STATIC_URL)
  ```



# Django 프로젝트 DB 생성과 Server 시작 

### 1. database 생성

- `terminal> python manage.py migrate`



### 2. Server 시작

- `terminal>  python manage.py runserver`
  - http://localhost:8000/admin/ 으로 접속



# Superuser 생성 및 관리자 화면

- `terminal> python manage.py createsuperuser`
  -  password를 체크 하므로 너무 짧거나 갂단한 글자로 입력하면 안됩니다.

