# Architecture :  ORM과 Queryset 



![image-20210112164120819](Django_ORM_QuerySets.assets/image-20210112164120819.png)



# QuerySet이란?

- 쿼리셋(QuerySet)은 DB로부터 데이터를 읽고, 필터링을 하거나, 정렬을 할 수 있습니다
- 쿼리셋을 사용하기 위해 먼저 python shell을 실행
  - 인터렉티브 콘솔(Interactive Console) 실행
    - `terminal>  python manage.py shell`



# 모든 객체 조회하기

```
In [1]: Post.objects.all()
Traceback (most recent call last):
	File "<console>", line 1, in <module>
NameError: name 'Post' is not defined
```

```python
In [1]: from blog.models import Post
```

```python
In [2]: Post.objects.all()
Out[2]: <QerySet [<Post: my post title>, <Post: another post title>]>
```



# 객체 생성하기

- 객체를 저장하기 위해 `create()`함수를 사용합니다.
- 작성자(author)로서 User(사용자) 모델의 인스턴스를 가져와 젂달 해줘야 합니다.

```python
In [1]: from djnago.contrib.auth.models import User

In [2]: User.objects.all()
Out[2]: <QuerySet [<User: admin>]>

In [3]: me = User.objects.get(username='<user-name>')
Out[3]: <User: <user-name>>

In [4]: Post.objects.create(author=me, title='Sample title', text='Test')
```



# 필터링 하기

-  특정 사용자가 작성한 글을 찾고자 할 때 

```python
In [1]: Post.objects.filter(author=me)
```



-  글의 제목(title)에 'title' 이라는 글자가 들어간 글을 찾고자 할 때

```python
In [1]: Post.objects.filter(title__contains='title')
```



- 게시일(published_date)로 과거에 작성한 글을 필터링하여 목록을 가져올 때

``` python
In [1]: from django.utils import timezone

In [2]: Post.objects.filter(published_date__lte=timezone.now())
```



- 게시(publish)하려는 Post의 인스턴스를 가져온다.

```python
In [1]: post = Post.objects.get(title="Sample title")
```



- 가져온 Post 인스턴스를 publish() 메서드를 이용하여 게시한다.

``` python
In [1]: post.publish()
```



# 정렬 하기

- 작성일(created_date) 기준으로 오름차순으로 정렬하기

``` python
In [1]: Post.objects.order_by('created_date')
```



- 작성일(created_date) 기준으로 내림차순으로 정렬하기 : – 을 붙이면 내림차순 정렬

```python
Post.objects.order_by('-created_date')
```



- 쿼리셋들을 함께 연결(chaining) 하기

```python
Post.objects.filter(published_date__lte=timezone.now()).order_by('published_date')
```

