# 로그인 처리하기 

#### 1.  @login_required 데코레이터

- 로그인 사용자만 포스트를 접근 할 수 있도록 `post_new`, `post_edit`, `post_remove`의 View들을 보호하려고 한다면
- Django에서 제공하는 `django.contrib.auth.decorators` 모듈 안의 `login_required` 데코레이터 를 사용하면 된다

```html
<!-- blog/views.py -->
from django.contrib.auth.decorators import login_required

@login_required
def post_new(request):
	[...]

@login_required
def post_remove(request):
	[...]

@login_required
def post_edit(request):
	[...]
```



#### 2.  urls.py 에 login url 추가

- blog/urls.py가 아니라 myjango/urls.py에 로그인 url 추가

```html
<!-- mydjango/urls.py -->
from django.contrib import admin
from django.urls import path, include
from django.contrib.auth import views as auth_views

urlpatterns = [
    # 생략
    path('accounts/login/', auth_views.LoginView.as_view(
		  template_name="registration/login.html"),name="login"),
]
```



#### 3. 로그인 페이지 템플릿 추가

- blog/templates/registration 디렉토리를 생성하고, login.html 파일 작성

```ht
{% extends "blog/base.html" %}

{% block content %}
    {% if form.errors %}
    	<p>이름과 비밀번호가 읷치하지 않습니다. 다시 시도해주세요.</p>
    {% endif %}
    
    <form method="post" action="{% url 'login' %}">
    	{% csrf_token %}
        <table class="table table-bordered table-hover">
        <tr>
        	<td>{{ form.username.label_tag }}</td><td>{{ form.username }}</td>
        </tr>
        <tr>
        	<td>{{ form.password.label_tag }}</td><td>{{ form.password }}</td>
        </tr>
        </table>
        
        <input type="submit" value="login" class="btn btn-primary btn-lg">
        <input type="hidden" name="next" value="{{ next }}">
    </form>
{% endblock %}
```



#### 4.  settings.py 에 설정 추가

- 로그인 하면 최상위 index 레벨에서 로그인이 된다.

```html
<!-- mydjango/settings.py -->
LOGIN_REDIRECT_URL = '/'
```



#### 5. 로그인 여부 체크하기

- 인증이 되었을 때는 추가/수정 버튼을 보여주고, 인증이 되지 않았을 때는 로그읶 버튼을 보여줌

```html
<!-- blog/templates/blog/base.html -->
<div class="page-header">
    {% if user.is_authenticated %}
    <a href="{% url 'post_new' %}" class="top-menu">
        <span class="glyphicon glyphicon-plus"></span>
    </a>
    {% else %}
    <a href="{% url 'login' %}" class="top-menu">
        <span class="glyphicon glyphicon-lock"></span>
    </a>
    {% endif %}
    <h1><a href="/">Django's Blog</a></h1>
</div>
```



- 로그인 사용자만 글을 수정, 삭제 할 수 있도록 체크하기
  - {% if %} 태그를 추가해 관리자로 로그인한 사용자들 맊 글 수정,삭제 링크가 보일 수 있게 만든다

```html
<!-- blog/templates/blog/post_detail.html -->
{% if user.is_authenticated %}
    <a class="btn btn-default" href="{% url 'post_edit' pk=post.pk %}">
        <span class="glyphicon glyphicon-pencil"></span>
	</a>
    <a class="btn btn-default" href="{% url 'post_remove' pk=post.pk %}">
        <span class="glyphicon glyphicon-remove"></span>
	</a>
{% endif %}

```



# 로그아웃 처리하기

####  1. base.html 수정하기

- “Hello <사용자이름>” 구문을 추가하여 인증된 사용자라는 것을 알려주고, logout link를 추가함

```html
<!-- blog/templates/blog/base.html -->
<div class="page-header">
    {% if user.is_authenticated %}
    <a href="{% url 'post_new' %}" class="top-menu">
    	<span class="glyphicon glyphicon-plus"></span>
    </a>
    <p class="top-menu">Hello {{ user.username }}<small>
     (<a href="{% url 'logout' %}?next={{request.path}}">Logout</a>)</small>
    </p>
    {% else %}
    <a href="{% url 'login' %}" class="top-menu">
    	<span class="glyphicon glyphicon-lock"></span>
    </a>
    {% endif %}
    <h1><a href="/">Django's Blog</a></h1>
</div>
```



#### 2. urls.py 에 logout url 추가

- blog/urls.py가 아니라 myjango/url.py에 로그아웃 url 추가

```html
urlpatterns = [
	# 생략
    path('accounts/logout/', auth_views.LogoutView.as_view(),
		 {'next': None}, name='logout'),
]

```

