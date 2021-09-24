## 대표적인 데이터 구조6: 해쉬 테이블 (Hash Table)

### 1. 해쉬 구조
* Hash Table: 키(Key)에 데이터(Value)를 저장하는 데이터 구조
  - Key를 통해 바로 데이터를 받아올 수 있으므로, 속도가 획기적으로 빨라짐
  - 파이썬 딕셔너리(Dictionary) 타입이 해쉬 테이블의 예: Key를 가지고 바로 데이터(Value)를 꺼냄
  - 보통 배열로 미리 Hash Table 사이즈만큼 생성 후에 사용 (공간과 탐색 시간을 맞바꾸는 기법)
  - <font color='#BF360C'>단, 파이썬에서는 해쉬를 별도 구현할 이유가 없음 - 딕셔너리 타입을 사용하면 됨</font>



### 2. 알아둘 용어
* 해쉬(Hash): 임의 값을 고정 길이로 변환하는 것
* 해쉬 테이블(Hash Table): 키 값의 연산에 의해 직접 접근이 가능한 데이터 구조
* 해싱 함수(Hashing Function): Key에 대해 산술 연산을 이용해 데이터 위치를 찾을 수 있는 함수
* 해쉬 값(Hash Value) 또는 해쉬 주소(Hash Address): Key를 해싱 함수로 연산해서, 해쉬 값을 알아내고, 이를 기반으로 해쉬 테이블에서 해당 Key에 대한 데이터 위치를 일관성있게 찾을 수 있음
* 슬롯(Slot): 한 개의 데이터를 저장할 수 있는 공간
* 저장할 데이터에 대해 Key를 추출할 수 있는 별도 함수도 존재할 수 있음
<img src="https://www.fun-coding.org/00_Images/hash.png" width=400 />



#### 3.1. hash table 만들기
* 참고: 파이썬 list comprehension - https://www.fun-coding.org/PL&OOP5-2.html

```python
hash_table = list([i for i in range(10)])
hash_table
# [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
```



#### 3.2. 이번엔 초간단 해쉬 함수를 만들어봅니다.
- 다양한 해쉬 함수 고안 기법이 있으며, 가장 간단한 방식이 Division 법 (나누기를 통한 나머지 값을 사용하는 기법)

```python
def hash_func(key):
    return key % 5
```



#### 3.3. 해쉬 테이블에 저장해보겠습니다.
- 데이터에 따라 필요시 key 생성 방법 정의가 필요함

```python
data1 = 'Andy'
data2 = 'Dave'
data3 = 'Trump'
data4 = 'Anthor'
## ord(): 문자의 ASCII(아스키)코드 리턴
print (ord(data1[0]), ord(data2[0]), ord(data3[0]))
print (ord(data1[0]), hash_func(ord(data1[0])))
print (ord(data1[0]), ord(data4[0]))
# 65 68 84
# 65 0
# 65 65
```



- 3.3.2. 해쉬 테이블에 값 저장 예
  - data:value 와 같이 data 와 value를 넣으면, 해당 data에 대한 key를 찾아서, 해당 key에 대응하는 해쉬주소에 value를 저장하는 예

```python
def storage_data(data, value):
    key = ord(data[0])
    hash_address = hash_func(key)
    hash_table[hash_address] = value
```



#### 3.4. 해쉬 테이블에서 특정 주소의 데이터를 가져오는 함수도 만들어봅니다.

```python
storage_data('Andy', '01055553333')
storage_data('Dave', '01044443333')
storage_data('Trump', '01022223333')
```



#### 3.5. 실제 데이터를 저장하고, 읽어보겠습니다.

```python
def get_data(data):
    key = ord(data[0])
    hash_address = hash_func(key)
    return hash_table[hash_address]
```

```python
get_data('Andy')	# '01055553333'
```



### 4. 자료 구조 해쉬 테이블의 장단점과 주요 용도
- 장점
  - 데이터 저장/읽기 속도가 빠르다. (검색 속도가 빠르다.)
  - 해쉬는 키에 대한 데이터가 있는지(중복) 확인이 쉬움
- 단점 
  - 일반적으로 저장공간이 좀더 많이 필요하다.
  - **여러 키에 해당하는 주소가 동일할 경우 충돌을 해결하기 위한 별도 자료구조가 필요함**
- 주요 용도
  - 검색이 많이 필요한 경우
  - 저장, 삭제, 읽기가 빈번한 경우
  - 캐쉬 구현시 (중복 확인이 쉽기 때문)



### 5. 프로그래밍 연습

<div class="alert alert-block alert-warning">
<strong><font color="blue" size="3em">연습1: 리스트 변수를 활용해서 해쉬 테이블 구현해보기</font></strong><br>
1. 해쉬 함수: key % 8<br>
2. 해쉬 키 생성: hash(data)
</div>

```python
hash_table = list([0 for i in range(8)])

def get_key(data):
    return hash(data)

def hash_function(key):
    return key % 8

def save_data(data, value):
    hash_address = hash_function(get_key(data))
    hash_table[hash_address] = value
    
def read_data(data):
    hash_address = hash_function(get_key(data))
    return hash_table[hash_address]
```

```python
save_data('Dave', '0102030200')
save_data('Andy', '01033232200')
read_data('Dave')
# '0102030200'
```

```python
hash_table	# ['0102030200', 0, 0, 0, 0, 0, 0, '01033232200']
```

