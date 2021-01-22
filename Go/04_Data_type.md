# 자료형의 종류와 특징

**어떤 데이터를 저장할 지 표현하는 것이 '자료형'입니다.** Go언어의 특징으로는 ':=' 용법을 활용한 자료형 추론이 가능합니다. 예를 들어 정수값은 int로 실수 값은 float32로 자동 할당됩니다.

`import "unsafe"`를 입력하면 **"unsafe.Sizeof(변수)"** 형태를 사용하여 선언한 자료형의 size를 알 수 있습니다.

| 자료형           | 선언 | 크기(byte)             |
| ---------------- | ---- | ---------------------- |
| 정수형(음수포함) | int  | n비트 시스템에서 n비트 |
| int8             | 1    |                        |
| int16            | 2    |                        |
| int32            | 4    |                        |
| int64            | 8    |                        |
| 정수형(0, 양수)  | uint | n비트 시스템에서 n비트 |
| uint8            | 1    |                        |
| uint16           | 2    |                        |
| uint32           | 4    |                        |
| uint64           | 8    |                        |
| uintptr          | 8    |                        |



# 문자열 타입

다른 언어에서 표현되는 null과 같이 Go언어에서 사용되는 nil이 아닐 수 있습니다. **string으로 선언한 문자열 타입은 immutable 타입으로서 값을 수정할 수 없습니다.**

예를 들어, `var str string = "hello"`와 같이 선언하고 `str[2] = 'a'`로 수정이 불가능합니다.

| 자료형 | 선언   | 크기(byte) |
| ------ | ------ | ---------- |
| 문자열 | string | 16         |



# 문자열의 표현

```go
package main

import "fmt"

func main() {
	// Raw String Literal. 복수라인.
	var rawLiteral string = `바로 실행해보면서 배우는 \n Golang`

	// Interpreted String Literal
	var interLiteral string = "바로 실행해보면서 배우는 \nGolang"

	plusString := "구름 " + "EDU\n" + "Golang"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	fmt.Println()
	fmt.Println(plusString)
}

> Output
바로 실행해보면서 배우는 \n Golang

바로 실행해보면서 배우는
Golang

구름 EDU
Golang
```



# 자료형의 변환

- 자동 형 변환(묵시적 형 변환)

```go
package main

import "fmt"

func main() {
	var num1, num2 int = 3, 4
	
	var result float32 = num1 / num2	
	
	fmt.Printf("%f", result)
}

> Output
오류
```



- 강제 형 변환(명시적 형 변환)

```go
package main

import "fmt"

func main() {
	var num int = 10
	var changef float32 = float32(num) //int형을 float32형으로 변환
	changei := int8(num)               //int형을 int8형으로 변환

	var str string = "goorm"
	changestr := []byte(str) //바이트 배열
	str2 := string(changestr) //바이트 배열을 다시 문자열로 변환

	fmt.Println(num)
	fmt.Println(changef, changei)

	fmt.Println(str)
	fmt.Println(changestr)
	fmt.Println(str2)
}

> Output
[103 111 111 114 109]
goorm
```

**Go언어에서는 형 변환을 할 때 변환을 명시적으로 지정해주어야합니다. **예를 들어 float32에서 uint로 변환할 때, 암묵적 변환은 일어나지 않으므로 **"uint(변수이름)"**과 같이 반드시 변환을 지정해줘야합니다. 만약 명시적인 지정이 없다면 런타임 에러가 발생합니다.