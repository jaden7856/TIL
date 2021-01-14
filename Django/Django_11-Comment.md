# Comment Model 속성들

- post(Post Model를 참조하는 Foreign Key)
- author(글쓴이)
- text(내용)
- created_date(작성일)
- approved_comment(승인여부)



# 댓글(Comment) 작성

#### 1.  Comment Class 작성

- models.py 에 Comment class 추가

```html
<!-- blog/models.py -->
from django.db import models
from django.utils import timezone

class Comment(models.Model):
	post = models.ForeignKey('blog.Post',on_delete=models.
							  CASCADE, related_name='comments')
    author = models.CharField(max_length=200)
    text = models.TextField()
    created_date = models.DateTimeField(default=timezone.now)
    approved_comment = models.BooleanField(default=False)
    
	def approve(self):
    	self.approved_comment = True
    	self.save()

    def __str__(self):
    	return self.text

```



#### 2. 테이블 생성 및 관리자 패널에 등록

- 마이그레이션 파일(migration file) 생성하기
  - terminal> python manage.py makemigrations blog
- 실제 데이터베이스에 Post, Comment Model 클래스를 반영하기
  - terminal> python manage.py migrate blog
- 관리자 페이지에서 만든 모델을 보기 위해 Comment 모델을 등록
  - blog/admin.py에 아래 코드를 추가

```html
from django.contrib import admin
from .models import Post,Comment

admin.site.register(Post, PostAdmin)
admin.site.register(Comment)
```



#### 3. Comment를 화면에 나타내기

- Comment를 화면에 나타나게 하기 위해 {% endblock %} tag 전에 아래 코드를 추가

```html
<!-- blog/templates/blog/post_detail.html -->
<hr>
{% for comment in post.comments.all %}
    <div class="comment">
        <div class="date">
            {{ comment.created_date }}
    	</div>
    	<strong>{{ comment.author }}</strong>
    	<p>{{ comment.text|linebreaks }}</p>
    </div>
{% empty %}
	<p>No comments here yet :(</p>
{% endfor %}
```

- blog/css/blog.css 에 아래의 코드 추가

```css
.comment {
	margin: 20px 0px 20px 20px;
}
```

- Post list 페이지에서 각 post 별 댓글 갯수를 출력하기

```html
<!-- blog/templates/blog/post_list.html -->
<a href="{% url 'post_detail' pk=post.pk %}">Comments:{{ post.comments.count }}</a>
```



#### 4. Comment를 등록하기

- blog/forms.py 파일 끝에 아래 코드를 추가하기

```python
from .models import Post, Comment

class CommentForm(forms.ModelForm):
    class Meta:
        model = Comment
        fields = ('author', 'text',)

```



- Comment 등록 url 추가

```python
# blog/urls.py
from django.urls import path

path('post/<int:pk>/comment/',
      views.add_comment_to_post, name='add_comment_to_post'),
```



- add_comment_to_post() 함수 추가

```python
# blog/views.py
from .forms import PostForm, CommentForm

def add_comment_to_post(request, pk):
    post = get_object_or_404(Post, pk=pk)
    if request.method == "POST":
        form = CommentForm(request.POST)
        if form.is_valid():
            comment = form.save(commit=False)
            comment.post = post
            comment.save()
            return redirect('post_detail', pk=post.pk)
        else:
            form = CommentForm()
            return render(request, 'blog/add_comment_to_post.html', {'form': form})
```



- Comment 등록 link 추가
  - {% for comment in post.comments.all %} 전에 아래 코드를 추가

```html
<!-- blog/templates/blog/post_detail.html -->
<a class="btn btn-default" href="{% url 'add_comment_to_post'
pk=post.pk %}">Add comment</a>
```



- Comment 등록 할 수 있는 템플릿 추가
  - blog/templates/blog/add_comment_to_post.html 추가

```html
{% extends 'blog/base.html' %}

{% block content %}
<table>
    <h1>New comment</h1>
    <form method="POST" class="post-form">{% csrf_token %}
        <table class="table table-bordered table-hover">
            {{ form.as_table }}
        </table>
        <button type="submit" class="save btn btn-default">Send</button>
    </form>
</table>
{% endblock %}
```



# 댓글 승인 , 삭제 하기

#### 1. post detail 페이지에 댓글 삭제, 승인 버튼을 추가합니다.

```html
<!-- blog/templates/blog/post_detail.html -->
{% for comment in post.comments.all %}
    {% if user.is_authenticated or comment.approved_comment %}
    <div class="comment">
        <div class="date">
            {{ comment.created_date }}
            {% if not comment.approved_comment %}
            <a class="btn btn-default" href="{% url 'comment_remove' 
                                             pk=comment.pk %}">
                <span class="glyphicon glyphicon-remove"></span></a>
            <a class="btn btn-default" href="{% url 'comment_approve' 
                                             pk=comment.pk %}">
                <span class="glyphicon glyphicon-ok"></span></a>
            {% endif %}
        </div>
        <strong>{{ comment.author }}</strong>
        <p>{{ comment.text|linebreaks }}</p>
    </div>
    {% endif %}
{% empty %}
	<p>No comments here yet :(</p>
{% endfor %}
```



#### 2. Comment 승인, 삭제 url 추가

```python
# blog/urls.py

path('comment/<int:pk>/approve/', views.comment_approve, name='comment_approve'),
path('comment/<int:pk>/remove/', views.comment_remove, name='comment_remove'),
```



- `comment_approve()`, `comment_remove()` 함수 추가

```python
# blog/views.py

@login_required
def comment_approve(request, pk):
    comment = get_object_or_404(Comment, pk=pk)
    comment.approve()
    return redirect('post_detail', pk=comment.post.pk)

@login_required
def comment_remove(request, pk):
    comment = get_object_or_404(Comment, pk=pk)
    post_pk = comment.post.pk
    comment.delete()
    return redirect('post_detail', pk=post_pk)
```



- 등록된 모든 댓글의 갯수가 대싞에 승인된 댓글의 갯수가 노출 되도록 수정

``` html
<!-- blog/templates/blog/post_list.html -->

<a href="{% url 'blog.views.post_detail' pk=post.pk %}">
    Comments: {{ post.approved_comments.count }}
</a>
```



- Post 모델에 approved_comments 메서드를 추가

```python
# blog/models.py

def approved_comments(self):
	return self.comments.filter(approved_comment=True)
```

