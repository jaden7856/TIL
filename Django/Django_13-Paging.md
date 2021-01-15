# Paginator 설치

- Django 프레임워크에서 지원하는 `paginator` 를 사용합니다.

  - ```python
    from django.core.paginator import Paginator
    ```



# Paginator 적용

#### 1. views.py 에 `post_list`함수를 수정한다.

```python
# blog/views.py

def post_list(request):
    post_list_paging = Post.objects.filter(published_date__lte=timezone.now()).
    				   order_by('published_date')
    paginator = Paginator(post_list_paging, 2)
    page_no = request.GET.get('page')
    
    try:
        posts = paginator.page(page_no)
    except PageNotAnInteger:
        posts = paginator.page(1)
    except EmptyPage:
        posts = paginator.page(paginator.num_pages) # num_pages : 전체 페이지 갯수
    
    return render(request, 'blog/post_list.html', {'posts': posts})
```



#### 2. post_list에 적용

- post_list에서 마지막 {% endblock %} 전에 추가한다.

```html
<!-- blog/post_list.html -->	
	{% endfor %}

    {% include 'blog/post_pagination.html' with page=posts %}

{% endblock %}
```



#### 2. pagination.html 생성

- [Bootstrap](https://getbootstrap.com/docs/5.0/components/pagination/)에 들어가 원하는 html을 복사한다.
  - 다른 곳에서 원하는 것을 찾아 써도 똑같다.
- 'blog/post_pagination.html'을 만들고 거기에 붙여넣기를 한다.

```html
<nav aria-label="Page navigation example">
    <ul class="pagination justify-content-center">
        <li class="page-item disabled">
            <a class="page-link" href="#">Previous</a>
        </li>
        <li class="page-item"><a class="page-link" href="#">1</a></li>
        <li class="page-item"><a class="page-link" href="#">2</a></li>
        <li class="page-item"><a class="page-link" href="#">3</a></li>
        <li class="page-item">
            <a class="page-link" href="#">Next</a>
        </li>
    </ul>
</nav>
```



- 위의 코드를 previous와 next를 활성화한 코드이다.
  - `href=""` 안에 `?`가 빠지지않게 주의한다.

```html
<nav aria-label="Page navigation example">
    <ul class="pagination justify-content-center">
        <li class="page-item disabled">
            {% if page.has_previous %}
                <a class="page-link" href="?page={{page.previous_page_number}}">
                    Previous
            	</a>
            {% endif %}
        </li>
        <li class="page-item">
            <a class="page-link">Page {{page.number}} of {{page.paginator.num_pages}}
            </a>
        </li>
        <li class="page-item">
            {% if page.has_next %}
                <a class="page-link" href="?page={{page.next_page_number}}">Next</a>
            {% endif %}
        </li>
    </ul>
</nav>
```



