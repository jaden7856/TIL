# Post New ( 글 추가 ) 페이지 작성하기

#### 1. `forms.py` 추가

   1) ModelForm을 생성해 자동으로 Model에 결과물을 저장할 수 있다.

   2) Form을 하나 만들어서 `Post` 모델에 적용한다.

   3) blog 디렉토리 안에 `forms.py`라는 파일을 작성한다.

```html
<!-- blog/forms.py -->
from django import forms
from .models import Post

class PostForm(forms.ModelForm):
	class Meta:
		model = Post
		fields = ('title', 'text',)

# forms.ModelForm은 django에 이 폼이 ModelForm이라는 것을 알려주는 구문이다.
class Meta 구문은 Form을 맊들기 위해서 어떤 model이 쓰여야 하는지 django에 알려주는 구문
이 폼에 필드는 title과 text맊 보여지게 된다. author는 로그인 하고 있는 사람이고,
created_date는 글이 등록되는 시간이다
```



#### 2.  `urls.py`에 Post New(글 추가) url 추가

   1) `post/new`란 URL이 `post` 문자를 포함해야 한다는 것을 말합니다

   2) /은 다음에 / 가 한 번 더 와야 한다는 의미입니다.

   3) $는 "마지막"을 말합니다. 그 뒤로 더는 문자가 오면 안 됩니다.

   4) ^은 "시작"을 뜻합니다.

```html
<!-- blog/urls.py -->
from django.urls import path
from . import views

urlpatterns = [
	path('', views.post_list, name='post_list'),
	path('post/<int:pk>/', views.post_detail, name='post_detail'),
    path('post/new/', views.post_new, name='post_new'),
]
```



#### 3.  base.html에 Post New(글 추가) 페이지 링크 추가

   1) blog/templates/blog/base.html 파일을 열어서, page-header 라는 `div class`에 등록 폼 `link`를 하나 추가한다.

   2) 새로운 view는 post_new입니다.

   3) 부트스트랩 테마에 있는 `glyphicon glyphicon-plus` 클래스로 더하기 기호가 보이게 됩니다.

```html
<!-- blog/templates/blog/base.html -->
<div class="page-header">
    <a href="{% url 'post_new' %}" class="top-menu">
        <span class="glyphicon glyphicon-plus"></span></a>
    <h1><a href="/">Django's Blog</a></h1>
</div>
```



#### 4. `views.py`에 `post_new()` 함수 추가

   1) 새 Post 폼을 추가하기 위해 `PostForm()` 함수를 호출하도록 하여 템플릿에 넘깁니다.

```html
<!-- blog/views.py -->
from .forms import PostForm

def post_new(request):
	form = PostForm()
	return render(request, 'blog/post_edit.html', {'form': form}) 
```



#### 5. post_edit.html 페이지 추가

   1) {% csrf_token %}를 추가하세요. 이 작업은 폼 보안을 위해 중요합니다.

   2) HTML Form의 POST요청에서 CSRF 토큰을 체크하며, 이때 CSRF토큰이 필요합니다. `csrf_token tag`를 통해 CSRF 토큰을 발급 받을 수 있습니다.

```html
<!-- blog/templates/blog/post_edit.html -->
{% extends 'blog/base.html' %}

{% block content %}
    <h1>New post</h1>
    <form method="POST" class="post-form">
        {% csrf_token %}
        {{ form.as_p }}
        <button type="submit" class="save btn btn-default">Save</button>
    </form>
{% endblock %}
```



#### 6-1) Form 저장하기

- 첫번째 : 처음 페이지에 접속 했을 때, 새 글을 쓸 수 있게 Form이 비어 있습니다. 이때의 Http method는 `GET` 
- 두번째 : Form에 입력된 데이터를 view 페이지로 가지고 올 때입니다. 이때의 Http method는 `POST`

```html
<!-- blog/views.py -->
from .forms import PostForm

def post_new(request):
    if request.method == "POST":
    	form = PostModelForm(request.POST)
    else:
    	form = PostModelForm()
    return render(request, 'blog/post_edit.html', {'form': form})
```



#### 6-2) Form 저장하기

- 폼에 들어있는 값들이 올바른 지를 확인하기 위해 `form.is_valid()`을 사용합니다.

- 첫번째 : `form.save()`로 폼을 저장하는 작업, `commit=False`란 데이터를 바로 Post 모델에 저장하지 않는다는 뜻입니다. 
- 두번째 : `author`와 `published_date`를 추가하는 작업,` post.save()`는 변경사항을 유지하고 새 블로그 글이 만들어 집니다.

```html
<!-- blog/views.py -->
if form.is_valid():
    post = form.save(commit=False)
    post.author = request.user
    post.published_date = timezone.now()
    post.save()
```



#### 6-3)  Form 저장하기

1) 새 블로그 글을 작성한 다음에 post_detail 페이지로 이동 합니다.
	
2) `post_detail`은 이동 해야 할 view의 name이고, `post_detail view`는 `pk=post.pk`를 사용해서 view에게 값을 넘겨줍니다.
	
3) post는 새로 생성한 블로그 글입니다.

```html
<!-- blog/views.py -->
from django.shortcuts import redirect

return redirect('post_detail', pk=post.pk)
```



#### 6-4) 완성된 post_new 함수

```html
<!-- blog/views.py -->
from django.contrib.auth.models import User
from django.shortcuts import redirect

def post_new(request):
    if request.method == "POST":
        form = PostForm(request.POST)
        if form.is_valid():
            post = form.save(commit=False)
            post.author = User.objects.get(username=request.user.username)
            post.published_date = timezone.now()
            post.save()
            return redirect('post_detail', pk=post.pk)
    else:
	    form = PostModelForm()
    return render(request, 'blog/post_edit.html', {'form': form})
```



# Post Edit ( 글 수정 ) 페이지 작성하기

#### 1. urls.py에 Post Edit(글 수정) url 추가

-  post/1/edit란 URL이 post 문자를 포함 해야 한다는 것을 말합니다

```html
<!-- blog/urls.py -->
urlpatterns = [
    # 생략
    path('post/<int:pk>/edit/', views.post_edit, name='post_edit'),
]
```



#### 2.  post_detail.html에 Post Edit(글 수정) 페이지 링크 추가

- ① blog/templates/blog/post_detail.html 파일을 열어서, link를 하나 추가한다.
- 부트스트랩 테마에 있는 `glyphicon glyphicon-pencil` 클래스로 아이콘이 보이게 됩니다.

```html
<div class="post">
    {% if post.published_date %}
    	<div class="date">
        	{{ post.published_date }}
    	</div>
    {% endif %}
    <a class="btn btn-default" href="{% url 'post_edit' pk=post.pk %}">
    <span class="glyphicon glyphicon-pencil"></span></a>
    <h1>{{ post.title }}</h1>
    <p>{{ post.text|linebreaksbr }}</p>
</div>
```



#### 3.  views.py에 `post_edit()` 함수 추가

- 첫 번째: url로부터 추가로 pk 매개변수를 받아서 처리합니다. 
- 두 번째: `get_object_or_404(Post, pk=pk)`를 호출하여 수정하고자 하는 글의 Post 모델 인스턴스(instance)로 가져온 데이터를 폼을 만들 때와 폼을 저장할 때 사용하게 됩니다.

```html
@login_required
def post_edit(request, pk):
	post = get_object_or_404(Post, pk=pk)
    if request.method == "POST":
    	form = PostModelForm(request.POST, instance=post)
    	if form.is_valid():
    		post = form.save(commit=False)
    		post.author = User.objects.get(username=request.user.username)
    		post.published_date = timezone.now()
    		post.save()
    		return redirect('post_detail', pk=post.pk)
    else:
    	form = PostModelForm(instance=post)
    
	return render(request, 'blog/post_edit.html', {'form': form})

```



# Post remove ( 글 삭제 ) 페이지 작성하기

#### 1. urls.py에 Post remove(글 삭제) url 추가

```html
<!-- blog/urls.py -->
from django.urls import path
from . import views

urlpatterns = [
	# 생략
	path('post/<int:pk>/remove/', views.post_remove, name='post_remove'),
]
```



#### 2.  post_detail.html에 Post Remove(글 삭제) 페이지 링크 추가

- blog/templates/blog/post_detail.html 파일을 열어서, link를 하나 추가한다.
- 부트스트랩 테마에 있는 glyphicon glyphicon-remove 클래스로 아이콘이 보이게 됩니다.

```html
<div class="post">
    {% if post.published_date %}
    	<div class="date">
        	{{ post.published_date }}
    	</div>
    {% endif %}
    <a class="btn btn-default" href="{% url 'post_remove' pk=post.pk %}">
    	<span class="glyphicon glyphicon-remove"></span>
    </a>
    <h1>{{ post.title }}</h1>
    <p>{{ post.text|linebreaksbr }}</p>
</div>
```



#### 3.  views.py에 `post_remove()` 함수 추가

- 첫 번째: url로부터 추가로 pk 매개변수를 받아서 처리 합니다. 
- 두 번째: `get_object_or_404(Post, pk=pk)`를 호출하여 삭제 하고자 하는 글의 Post 모델 인스턴스(instance)로 가져 와서 삭제 처리를 한다.

```html
<!-- blog/views.py -->
@login_required
def post_remove(request, pk):
    post = get_object_or_404(Post, pk=pk)
    post.delete()
    return redirect('post_list')
```

