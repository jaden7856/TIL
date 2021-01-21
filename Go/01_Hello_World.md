### 모든 언어의 시작 Hello World 를 만들어보자!

```go
package main

import "fmt"

func main() {
    var num1 int = 1
    var num2 int = 2
    
    fmt.Print("Hello World!", num1, num2, "\n")
    
    fmt.Println("Hello World!", num1, num2)
	
    fmt.Printf("num1의 값은:%d num2의 값은:%d\n", num1, num2)
}


# Output
> Hello World!1 2
Hello World! 1 2
num1의 값은:1 num2의 값은:2
```

### fmt

- 일반적으로 Go언어에서 콘솔 입출력을 위해서는 fmt 패키지를 `import` 해서 사용합니다.



### println, print

- Go언어에서는 꼭 `fmt` 패키지를 `import` 하지 않아도 기본적으로 콘솔 출력 함수인 `Println`과 `Print` 함수를 지원합니다. 두 함수의 차이점은 단순히 호출 후 개행을 하느냐 안 하느냐 입니다.

- 이 함수들은 함수 안에서의 연산 식을 결과 값으로 출력이 가능합니다. 예를 들어 `fmt.Println(3+5)`를 하면 8이 출력됩니다. 



# 실행 방법

- `$ go run <file_name>.go`



