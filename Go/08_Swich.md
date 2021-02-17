# Swich문에 의한 선택적 실행

`if ~else` `if`문은 num이 1일 때 이 조건과 맞는 조건문을 하나씩 검토해서 출력하는 느낌입니다. '여기는 2일때.. 여기는 3일때.. 여기는 4일때.. 여기가 1일때 실행하는 곳이구나!'라는 느낌입니다. 하지만 `switch`문은 num이 1이면 라벨이 1인 곳을 딱 찾아내서 수행 구문을 실행시키는 느낌입니다. '음.. 1일때 실행해야 하는 곳이.. 바로 여기구나!'라는 느낌입니다.

switch문은 기본적으로 변수를 가져와 `switch` 옆에 '태그'로 사용합니다. 변수는 어느 자료형이든 쓸 수 있습니다. 태그의 값에 따라 `case`의 '라벨'과 일치하는 것을 찾고 일치하는 `case`의 구문을 수행합니다. Go언어에서는 `switch` 옆에 태그뿐만이 아니라 '표현식'을 쓰는 경우가 있습니다. 그리고 case 옆에도 라벨뿐만이 아니라 참/거짓을 판별할 수 있는 표현식을 쓰는 경우도 있습니다. 그리고 **태그나 표현식이 어느 조건에도 맞지 않는다면 default문을 사용해 해당 구문을 수행합니다.**

그리고 굳이 `if`문처럼 블록 시작 브레이스({)를 같은 줄에 쓰지 않아도 실행이 됩니다. 그리고 **`break`를 따로 입력하지 않아도 해당되는 `case`만 수행**합니다.

```go
package main
 
import "fmt"
 
func main() {
	var num int
	fmt.Print("정수 입력:")
	fmt.Scanln(&num)
	
	switch num {
	case 0:
		fmt.Println("영")
	case 1:
		fmt.Println("일")
	case 2:
		fmt.Println("이")
	case 3:
		fmt.Println("삼")
	case 4:
		fmt.Println("사")
	default:
		fmt.Println("모르겠어요.")
	}
}
```



## 쓰임새가 비교적 넓은 Go언어에서의 switch문

- switch에 전달되는 인자로 태그 사용
- switch에 전달되는 인자로 표현식 사용
- switch에 전달되는 인자 없이 case에 표현식 사용(참/거짓 판별)

```go
//switch에 전달되는 인자로 태그 사용
package main

import "fmt"

func main() {
	var fruit string
	
	fmt.Print("apple, banana, grape중에 하나를 입력하시오:")
	fmt.Scanln(&fruit)
	
	if (fruit != "apple") && (fruit != "banana") && (fruit != "grape") {
		fmt.Println("잘못 입력했습니다.")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("RED")
	case "banana":
		fmt.Println("YELLOW")
	case "grape":
		fmt.Println("PURPLE")
	}
}
```

위 예시 코드는 일부러 if문 조건식을 사용했습니다. "apple", "banana", "grape" 세 개 이외의 값을 입력하면 "잘못 입력했습니다." 라는 구문이 출력되고 프로그램이 종료됩니다. 여기서 알아야 할 점이 있습니다.

- `defualt`문을 사용하지 않으면 `if`문을 사용해 따로 예외 처리를 해야하기 때문에 코드가 길어집니다.
- "return"을 실행하면 해당 함수가 종료됩니다. main 함수 안에서 return은 main 함수를 종료한다는 것을 의미하기 때문에 프로그램이 종료됩니다.

```go
//switch에 전달되는 인자로 표현식 사용
package main

import "fmt"

func main() {
	var num int
	var result string
	
	fmt.Print("10, 20, 30중 하나를 입력하시오:")
	fmt.Scanln(&num)

	switch num / 10 { //표현식
	case 1:
		result = "A"
	case 2:
		result = "B"
	case 3:
		result = "C"
	default:
		fmt.Println("모르겠어요.")
		return
	}
	
	fmt.Println(result)
}
```

위 예시는 default문을 사용해 예외 처리를 했습니다. 그리고 return을 입력하지 않았다면 "모르겠어요."를 출력한 뒤 아래 줄인 `fmt.Println(result)`를 실행합니다. 따라서 아무 값도 초기화되지 않은 `result`는 빈칸으로 출력됩니다. 불 필요한 실행을 막기 위해 잘못된 입력이 되면 `return`으로 프로그램을 종료한 것입니다.

다른 언어에서는 switch문의 쓰임에 제한이 많아 활용도 제한적이었지만, Go언어에서는 쓰임새가 확장되어 조건문을 독점하다시피했던 `if`문의 지분을 많이 차지할 수 있습니다. 따라서 `if ~else` `if`문을 사용할지 `switch`문을 사용할지는 취향에 따른 문제입니다. **보통 조건이 많지 않다면 `if ~else` `if`, 조건이 많다면 `switch`문을 사용합니다.**