### 1. 링크드 리스트 (Linked List) 구조
* 연결 리스트라고도 함
* 배열은 순차적으로 연결된 공간에 데이터를 나열하는 데이터 구조
* 링크드 리스트는 떨어진 곳에 존재하는 데이터를 화살표로 연결해서 관리하는 데이터 구조
* <font color='#BF360C'>본래 C언어에서는 주요한 데이터 구조이지만, 파이썬은 리스트 타입이 링크드 리스트의 기능을 모두 지원</font>



* 링크드 리스트 기본 구조와 용어
  - 노드(Node): 데이터 저장 단위 (데이터값, 포인터) 로 구성
  - 포인터(pointer): 각 노드 안에서, 다음이나 이전의 노드와의 연결 정보를 가지고 있는 공간


* 일반적인 링크드 리스트 형태
<img src="https://www.fun-coding.org/00_Images/linkedlist.png" />
(출처: wikipedia, https://en.wikipedia.org/wiki/Linked_list)





### 2. 간단한 링크드 리스트 예

#### Node 구현
- 보통 파이썬에서 링크드 리스트 구현시, 파이썬 클래스를 활용함
  - 파이썬 객체지향 문법 이해 필요
  - 참고: https://www.fun-coding.org/PL&OOP1-3.html

```python
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None
        
        
class Node:
    def __init__(self, data, next=None):
        self.data = data
        self.next = next
```

위의 코드에서 두개의 클래스간에 다른점은 `__init__`메소드에 `next`값의 default를 정해준것이 다른 점입니다.



#### Node와 Node 연결하기 (포인터 활용)

```python
node1 = Node(1)
node2 = Node(2)
node1.next = node2
head = node1
```

Node연결을 위해 node1의 `next`에 node2의 `data`값을 넣어주고 node1의 `data`값을 `head`라는 변수에 저장하겠습니다.



#### 링크드 리스트로 데이터 추가하기

```python
class Node:
    def __init__(self, data, next=None):
        self.data = data
        self.next = next

def add(data):
    node = head
    while node.next:
        node = node.next
    node.next = Node(data) 
```

`add`라는 메서드를 생성하여 추가하는 코드를 연습하겠습니다.

우선 node에 제일 처음값 head를 넣어주고 맨끝의 노드에 찾아가기 위해서 `node.next`에 에 값이 **있으면** `node`에 `node.next`의 값을 넣어주고 **없으면** `data`의 값을 넣어줍니다.



```python
node1 = Node(1)
head = node1
for index in range(2, 10):
    add(index)
```

테스트를 위해 확인을 하면 `node1`에 `1`이라는 `data`값을 넣어주고 `head`에 저장합니다. 그 다음 `2`부터 `9`까지 값을 넣어주기위해서 `add`라는 함수를 넣어주도록 하겠습니다.





#### 링크드 리스트 데이터 출력하기(검색하기)

```python
node = head
while node.next:
    print(node.data)
    node = node.next
print (node.data)
# 1
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
```





### 3. 링크드 리스트의 장단점 (전통적인 C언어에서의 배열과 링크드 리스트)
* 장점
  - 미리 데이터 공간을 미리 할당하지 않아도 됨
    - 배열은 **미리 데이터 공간을 할당** 해야 함
* 단점
  - 연결을 위한 별도 데이터 공간이 필요하므로, 저장공간 효율이 높지 않음
  - 연결 정보를 찾는 시간이 필요하므로 접근 속도가 느림
  - 중간 데이터 삭제시, 앞뒤 데이터의 연결을 재구성해야 하는 부가적인 작업 필요



### 4. 링크드 리스트의 복잡한 기능1 (링크드 리스트 데이터 사이에 데이터를 추가)
- 링크드 리스트는 유지 관리에 부가적인 구현이 필요함

<img src="https://www.fun-coding.org/00_Images/linkedlistadd.png" />
(출처: wikipedia, https://en.wikipedia.org/wiki/Linked_list)



```python
node = head
while node.next:
    print(node.data)
    node = node.next
print (node.data)
# 1
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
```

위의 코드에서 새로운값을 코드 중간에 추가해보도록 하겠습니다.



```python
node3 = Node(1.5)
```

`1.5`라는 값을 `node3`변수에 저장을 합니다. 그러면 제가 원하는 위치는 `1`과 `2`사이에 넣기위해서 `1`의 위치를 찾기위해 코드를 작성하겠습니다.



```python
node = head
search = True
while search:
    if node.data == 1:
        search = False
    else:
        node = node.next

node_next = node.next
node.next = node3
node3.next = node_next
```

맨처음 값은 `head`에 저장후 `search`라는 `True`값을 넣어준 다음 `node.data`가 만약 `1`이면 `search`를 `False`값을 주고 아니면 `node.next`값을 찾습니다.

`1`이라는 `data`값을 찾으면 빠져나와 `node.next`값을 다른 변수에 저장후, 우리가 원하는 `1.5`값을 `node.next`값에 넣어주겠습니다.

그러면 `1.5`다음에 `2`가 와야하기때문에 다른 변수에 저장했던 것을 `node3.next`에 저장하도록 하겠습니다.



```python
node = head
while node.next:
    print(node.data)
    node = node.next
print (node.data)
# 1
# 1.5
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
```

제가 원했던 `1` 다음에 `1.5` 그리고 `2`가 오도록 완성하였습니다.





### 5. 파이썬 객체지향 프로그래밍으로 링크드 리스트 구현하기

```python
class Node:
    def __init__(self, data, next=None):
        self.data = data
        self.next = next
    
class NodeMgmt:
    # 맨 처음 노드값을 알고있어야 하기때문에 head에 추가
    def __init__(self, data):
        self.head = Node(data)
    
    # 맨 끝에 값을 추가하는 메서드
    def add(self, data):
        if self.head == '':
            self.head = Node(data)
        else:
            node = self.head
            while node.next:
                node = node.next
            node.next = Node(data)
    
    # 값 전체를 출력시키는 메서드    
    def desc(self):
        node = self.head
        while node:
            print (node.data)
            node = node.next
```

```python
linkedlist1 = NodeMgmt(0)
linkedlist1.desc()
# 0
```

```python
for data in range(1, 10):
    linkedlist1.add(data)
linkedlist1.desc()
# 0
# 1
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
```





### 6. 링크드 리스트의 복잡한 기능2 (특정 노드를 삭제)
* 다음 코드는 위의 코드에서 delete 메서드만 추가한 것이므로 해당 메서드만 확인하면 됨

```python
class Node:
    def __init__(self, data, next=None):
        self.data = data
        self.next = next

# 위의 기존 코드에서 `delete`메서드만 추가하겠습니다.
class NodeMgmt:
    		:
	        :
    
    # 삭제 메서드
    def delete(self, data):
        # 처음 head값이 없으면
        if self.head == '':
            print ("해당 값을 가진 노드가 없습니다.")
            return
        
        # head값이 있을때 data변수가 같으면
        if self.head.data == data:
            # self.head 객체를 삭제하기 위해, 임시로 temp에 담아서 객체를 삭제한 이유는
            # self.head 객체를 삭제하면, 밑의 코드에서 실행이 안되기 때문!
            temp = self.head
            # head의 주소가 다음 노드의 head로 변경이 되어야 하기때문에 바꾸는 코드
            self.head = self.head.next
            # temp변수를 삭제
            del temp
        else:
            node = self.head
            # 다음 노드가 있으면
            while node.next:
                # 다음 노드의 데이터가 같은 데이터면
                if node.next.data == data:
                    temp = node.next
                    # 삭제할 노드의 다음 노드를 그 이전 노드와 연결
                    node.next = node.next.next
                    del temp
                    return
                
                # 다음의 노드의 데이터가 다른 데이터면
                else:
                    # 다음 노드로 간다
                    node = node.next
```





#### 테스트를 위해 1개 노드를 만들어 봄

```python
linkedlist1 = NodeMgmt(0)
linkedlist1.desc()	# 0
```





#### head 가 살아있음을 확인

```python
linkedlist1.head
# <__main__.Node at 0x1099fc6a0>
```





#### head 를 지워봄(위에서 언급한 경우의 수1)

```python
linkedlist1.delete(0)
```





#### 다음 코드 실행시 아무것도 안나온다는 것은 linkedlist1.head 가 정상적으로 삭제되었음을 의미

```python
linkedlist1.head
```





#### 다시 하나의 노드를 만들어봄

```python
linkedlist1 = NodeMgmt(0)
linkedlist1.desc()	# 0
```





#### 이번엔 여러 노드를 더 추가해봄

```python
for data in range(1, 10):
    linkedlist1.add(data)
linkedlist1.desc()
# 0
# 1
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
```





#### 노드 중에 한개를 삭제후 확인 (위에서 언급한 경우의 수2)

```python
linkedlist1.delete(4)
linkedlist1.desc()
# 0
# 1
# 2
# 3
# 5
# 6
# 7
# 8
# 9
```





### 7. 다양한 링크드 리스트 구조 
* 더블 링크드 리스트(Doubly linked list) 기본 구조 
  - 이중 연결 리스트라고도 함
  
  - 장점: 양방향으로 연결되어 있어서 노드 탐색이 양쪽으로 모두 가능

  
  
    <img src="https://www.fun-coding.org/00_Images/doublelinkedlist.png" />
    (출처: wikipedia, https://en.wikipedia.org/wiki/Linked_list)

```python
class Node:
    def __init__(self, data, prev=None, next=None):
        self.prev = prev
        self.data = data
        self.next = next

class NodeMgmt:
    def __init__(self, data):
        self.head = Node(data)
        self.tail = self.head

    def insert(self, data):
        if self.head == None:
            self.head = Node(data)
            self.tail = self.head
        else:
            node = self.head
            while node.next:
                node = node.next
            new = Node(data)
            node.next = new
            new.prev = node
            self.tail = new

    def desc(self):
        node = self.head
        while node:
            print (node.data)
            node = node.next
```

```python
double_linked_list = NodeMgmt(0)
for data in range(1, 10):
    double_linked_list.insert(data)
double_linked_list.desc()
# 0
# 1
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
```



<strong><font color="blue" size="3em">연습3: 위 코드에서 노드 데이터가 특정 숫자인 노드 앞에 데이터를 추가하는 함수를 만들고, 테스트해보기</font></strong><br>

- 더블 링크드 리스트의 tail 에서부터 뒤로 이동하며, 특정 숫자인 노드를 찾는 방식으로 함수를 구현하기<br>
- 테스트: 임의로 0 ~ 9까지 데이터를 링크드 리스트에 넣어보고, 데이터 값이 2인 노드 앞에 1.5 데이터 값을 가진 노드를 추가해보기

```python
class Node:
    def __init__(self, data, prev=None, next=None):
        self.prev = prev
        self.data = data
        self.next = next

class NodeMgmt:
    def __init__(self, data):
        self.head = Node(data)
        self.tail = self.head

    def insert(self, data):
        if self.head == None:
            self.head = Node(data)
            self.tail = self.head
        else:
            node = self.head
            while node.next:
                node = node.next
            new = Node(data)
            node.next = new
            new.prev = node
            self.tail = new

    def desc(self):
        node = self.head
        while node:
            print (node.data)
            node = node.next
    
    def search_from_head(self, data):
        if self.head == None:
            return False
    
        node = self.head
        while node:
            if node.data == data:
                return node
            else:
                node = node.next
        return False
    
    def search_from_tail(self, data):
        if self.head == None:
            return False
    
        node = self.tail
        while node:
            if node.data == data:
                return node
            else:
                node = node.prev
        return False
    
    def insert_before(self, data, before_data):
        if self.head == None:
            self.head = Node(data)
            return True
        else:
            node = self.tail
            while node.data != before_data:
                node = node.prev
                if node == None:
                    return False
            new = Node(data)
            before_new = node.prev
            before_new.next = new
            new.prev = before_new
            new.next = node
            node.prev = new
            return True
```

```python
double_linked_list = NodeMgmt(0)
for data in range(1, 10):
    double_linked_list.insert(data)
double_linked_list.desc()
```

```python
node_3 = double_linked_list.search_from_tail(3)
node_3.data	# 3
```

```python
double_linked_list.insert_before(1.5, 2)
double_linked_list.desc()
# 0
# 1
# 1.5
# 2
# 3
# 4
# 5
# 6
# 7
# 8
# 9
```

