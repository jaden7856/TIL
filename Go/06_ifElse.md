# 규칙

### 1. True/False

**Go언어의 조건문의 조건식은 반드시 Boolean 형으로 표현돼야 합니다.** Go언어에서 bool 형은 false와 true만 지원하는 것을 지난 강의에서 배웠습니다. C언어와 C++는 조건식에 1, 0과 같은 숫자를 쓸 수 있는 것과 대조적입니다.



### 2. 조건식의 괄호는 생략 가능

대부분의 언어들은 조건문에서 조건식을 쓸 때 if 옆에 괄호를 입력했을 것입니다.

- 예를 들어, "if(k==0)" 이렇게 말입니다.

하지만, **Go언어에서는 "if k==0"과 같이 괄호를 생략해서 입력해도 됩니다.** Go에서는 생략해서 실행하는 것을 권장하고 Atom과 같은 에디터에서 패키지를 설치하고 실행한다면 괄호를 자동으로 생략해줍니다.



### 3. 조건문의 중괄호는 필수

- **Go언어에서는 반드시 중괄호를 입력해야합니다.**



### 4. 괄호의 시작과 else문은 같은 줄에

아래와 같은 코드로 작성해 주세요

```go
if num == 1 {
		fmt.Print("hello\n")
	} else if num == 2 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
```



### 5. 조건식에 간단한 문장(Optional Statement) 실행 가능

"if val := num*2 ; val==2" 와 같이 조건식 앞에 변수를 선언하고 식을 입력할 수 있다.

주의해야할 점은 **조건식 전에 정의된 변수는 해당 조건문 블록에서만 사용할 수 있다**는 것이다.

```go
package main

import "fmt"

func main() {
	var num int

	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if val := num * 2; val == 2 {
		fmt.Print("hello\n")
	} else if val := num * 3; val == 6 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
}
```

