# 가상환경 생성

- 각 프로젝트마다 최적화된 라이브러리 저장을위해 생성
- `$ conda info --envs` - 현재 가지고있는 가상환경 정보 출력
- `$ conda create -n <name>` - 생성
- `$ conda activate <name>` - <name>에 접속



# 설치

- 가상환경에 필요한 라이브러리 설치
  `$ conda install scrapy`
  - 한개를 설치하면 나머지 다른 필수적인 라이브러리를 자동 설치해준다.



# 프로젝트 생성

- 프로젝트 생성
  - `$ scrapy startproject <project_name>`

- Bot 생성
  - `$ scrapy genspider <bot_name> "URL"`
    - "URL"은 자신이 크롤링할 사이트를 넣는다
      - ex) "movie.naver.com/movie/point/af/list.nhn"



### 프로젝트 열기

- `$ code .`



# settings.py

- txt파일 생성 유무
  - `ROBOTSTXT_OBEY = True` -> False

- 크롤링한 데이터 파일 저장 포맷
  - `FEED_FORAMT="csv"` ex) csv, json, xml
- 파일 이름과 경로 지정
  - `FEED_URL="my_movie_reviews.csv"`
- 한글 지정
  - `FEED_EXPORT_ENCODING="utf-8-sig"`



# items.py

```python
class MymovieItem(scrapy.Item):
    # 스크래핑할 데이터의 이름 지정
    title = scrapy.Field()
    writer = scrapy.Field()
```

- 모든 데이터는 HTML을 거치면서 숫자데이터도 String이 되기때문에 모든 작업을 .Field()로 한다.



# <bot_name>.py

```python
#desc -> 40데이터 (공백포함) -> 10데이터 (공백 제거)
def remove_space(desc:list) -> list:
    result = []
    # 공백제거
    for i in range(len(descs)):
        if len(descs[i].strip()) > 0:
            result.append(descs[i].strip())

    return result

class MymovieBotsSpider(scrapy.Spider):
    name = '<bot_name>'
    allowed_domains = ['naver.com']
    start_urls = ['http://movie.naver.com/movie/point/af/list.nhn']

    def parse(self, response):
        titles = response.xpath('//*[@id="old_content"]/table/tbody/tr/td[2]/a[1]/text()').extract()
        conv_titles = remove_space(titles)
        writers = response.css(".author::text").extract()
        
        for row in zip(conv_titles, writers):
            item = MymovieItem()
            item['title'] = row[0]
            item['writer'] = row[1]

            yield item
```

- `//*[@id="old_content"]/table/tbody/tr[1]/td[2]/a[1]`
  - 위 `xpath`는 크롤링할 title의 `xpath`를 copy한 것이고 우리가 쓸것은 `tr[1] `에서
  - 여러 title 데이터는 `tr[2]`, `tr[3]` ... `tr[10]` 이기때문에 `tr`만 쓰고
  - 그 곳에서 text만 필요하기때문에 `/text()`를 추가한 후에 
  - `.extract()`로 selector가 아닌 실제값을 추출하게 한다.

- `response.css(".author::text")`
  - css를 쓸 경우는 `class`나 `id`가 원하는 데이터에서만 쓰고, 바로 text를 추출할 수 있을경우에 쓴다.
  - `.`은 `class`, `#`은 `id`이다.

- `remove_space`는 text를 출력할때 앞뒤로 space(\t\n)같은 것들도 같이 출력이 될때 list에서 공백을 제거





# 실행

- `$ cd <project_name>`
- `$ scrapy crawl <bot_name>`



### 실행 상태

- response code (status code)
  - 2xx -> ok
    - 200 - Success
    - 201 - Created
  
  - 3xx -> x(System에서 사용)
  
  - 4xx -> Error(Client 문제)
  
    - 400 - Bad Request
    - 401 - Unauthorized
    - 404 - Resource Not Found
  
  - 5xx -> Error(Server 문제)
  
    