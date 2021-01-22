# 오직 for

**Go언어에서는 while문을 제공하지 않아 for문만 사용할 수 있습니다.**  for문의 쓰임새를 확장시킴으로써 무한루프와 같은 기능도 for문으로 간결하게 구현할 수 있게 만들었기 때문입니다. 

하지만 if문과 같이 코드 작성에 깔끔함을 추구하기 때문에 for문 역시 if문과 같이 블록 시작 브레이스({)를 for문 을 선언한 같은 줄에 입력해야합니다. 

```go
package main
 
import "fmt"
 
func main() {
	sum := 0
	
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1부터 10까지 정수 합계:", sum)
}
```



### 조건식만 쓰는 for 루프

```go
package main

import "fmt"

func main() {
	n := 2
	
	for n < 100 {
		fmt.Printf("count %d\n", n)
		
		n *= 2
	}
}
```



### 무한루프

- Go언어에서는 `for {`와 같은 형식으로 입력(모든 식을 생략)하는 것만으로 무한루프가 됩니다.

- 무한루프를 빠져나오기 위해서는 맥과 윈도우 동일하게 ctrl+c 를 입력하면 됩니다.

```go
package main

import "fmt"

func main() {
	for {
		fmt.Printf("무한루프입니다.\n")
	}
}
```



### for range문

 Go언어에서는 배열을 `var arr [3]int = [3]int{1, 2, 3}`와 같은 형식으로 선언합니다.

`for range`문은 **"for 인덱스, 요소값 := range 컬렉션이름"** 같이 for 루프를 구성합니다. range 키워드 다음의 컬렉션으로부터 하나씩 요소를 리턴해서 그 요소의 위치인덱스와 값을 for 키워드 다음의 2개의 변수에 각각 할당합니다. 즉, `for range`문은 컬렉션의 모든 요소에 접근해 차례로 리턴할 때 사용합니다.

```go
package main

import "fmt"

func main() {
	var arr [6]int = [6]int{1, 2, 3}

	for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
	}
}

> Output
arr[0]의 값은 1입니다.
arr[1]의 값은 2입니다.
arr[2]의 값은 3입니다.
```

`for range`문은 굳이 인덱스와 요소값을 모두 받아오지 않아도 됩니다. **인덱스와 요소값 둘 중에 하나를 생략해서 사용할 수 있습니다.**

**인덱스를 생략하기 위해서는 `for _, num :=`, 요소값을 생략하기 위해서는 `for num :=`로만 입력하면 됩니다.**



컬랙션의 맵을 활용하면 인덱스가 꼭 정수가 아니더라도 다양한 형태로 선언할 수 있기 때문에 for range문을 다양한 형태로 활용할 수 있습니다.

```go
package main

import "fmt"

func main() {
	var fruits map[string]string = map[string]string{
		"apple":  "red",
		"banana": "yellow",
	}

	for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다.\n", fruit, color)
	}
}

> Output
apple의 색깔은 red입니다.
banana의 색깔은 yellow입니다.
```

