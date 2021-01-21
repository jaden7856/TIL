# 변수의 선언과 초기화

Go에서의 변수 선언 방식은 `var <name> 변수형`입니다. 그리고 변수를 선언한 곳에서 **바로 초기값을 설정**할 수 있습니다.

### Short Assignment Statement

-  **`:=`**
  - 이를 사용하면 별다른 **형 선언 없이** 타입 추론이 가능합니다.
- 주의해야할 점은 이 용법은 **함수(func) 내에서만 사용이 가능**합니다

- 따라서 함수 밖에서(전역 변수)는 꼭 `var` 키워드를 선언해줘야합니다. 



아래 코드를 보면 `a`는 선언과 동시에 1로 초기화 됩니다. 그리고 `c`와 `d`는 오류없이 `int`와 `string`이라는 자료형으로 자동 지정됩니다.

```go
var a int = 1
var b string = "Hello"
    
c := 1
d := "Hello"
```



**Go언어에서는 선언만 하고 쓰지 않았다면 에러를 발생하며 컴파일에 실패합니다.**

이는 변수, 패키지, 함수 등 모든 선언에서 동일하게 적용됩니다.

따라서 꼭 쓰이는 변수만 선언해야하며 값을 지울때는 선언한 모든 부분을 지워야 합니다.



```go
package main

import "fmt"

var globalA = 5 //함수 밖에서는 'var' 키워드를 입력해야함.
				// 꼭 형을 명시하지 않아도 됨
func main() {
    var a string = "goorm"	//groom
    fmt.Println(a)

    var b int = 10	// 10
    fmt.Println(b)

    var e int	// 0
    fmt.Println(e)

    i, j, k := 1, 2, 3		// 1, 2, 3
    fmt.Println(i, j, k)
    
    var str1, str2 string = "Hello", "World"	// Hello World
    fmt.Println(str1, str2)
	
	fmt.Println(globalA)	// 5
}
```

주석은 C언어와 동일한  `//`는 한 줄을 주석처리하고, `/*  */`는 `/*`과 `*/`사이에 들어간 내용을 라인에 상관없이 전부 주석처리합니다.