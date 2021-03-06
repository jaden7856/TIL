# 연산자 우선순위

| 순위   | 연산기호                                     | 연산자                        | 결합방향 |
| ------ | -------------------------------------------- | ----------------------------- | -------- |
| 1      | ()                                           | 함수호출                      | →        |
| []     | 인덱스                                       |                               |          |
| ->     | 간접지정                                     |                               |          |
| ++, -- | 증가 및 감소                                 |                               |          |
| 2      | +, -                                         | 부호 연산(음수와 양수의 표현) | ←        |
| !      | 논리 NOT                                     |                               |          |
| ~      | 비트 단위 NOT                                |                               |          |
| (type) | 타입 변환                                    |                               |          |
| *      | 간접 지정 연산                               |                               |          |
| &      | 주소연산                                     |                               |          |
| sizeof | 바이트 단위 크기 계산                        |                               |          |
| 3      | *,/,%                                        | 곱셈, 나눗셈 관련 연산        | →        |
| 4      | +,-                                          | 덧셈, 뺄셈                    | →        |
| 5      | <<, >>                                       | 비트 이동                     | →        |
| 6      | <, <=, >, =>                                 | 대소 비교                     | →        |
| 7      | ==, !=                                       | 동등 비교                     | →        |
| 8      | &                                            | 비트 AND                      | →        |
| 9      | ^                                            | 비트 XOR                      | →        |
| 10     | \|                                           | 비트 OR                       | →        |
| 11     | &&                                           | 논리 AND                      | →        |
| 12     | \|\|                                         | 논리 OR                       | →        |
| 13     | ? :                                          | 조건 연산                     | ←        |
| 14     | =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, \|= | 대입 연산                     | ←        |
| 15     | ,                                            | 콤마 연산                     | →        |



# 콘솔 입력 함수의 기본

- Scanln은 여러 값을 동시에 입력받을 수 있습니다. **빈칸(스페이스바)으로 값을 구분하고 엔터(개행)를 입력하면 입력이 종료됩니다. 입력받는 변수에 '&' 연산자를 붙여 입력받습니다.** 물론 입력받는 변수는 미리 선언되어야 합니다.

```go
package main

import "fmt"

func main() {
	var num1, num2, num3 int
	
	fmt.Printㄹ("정수 3개를 입력하세요 :")
	fmt.Scanln(&num1, &num2, &num3)
	fmt.Println(num1, num2, num3)
}

> Output
정수 3개를 입력하세요 : 1 2 3
1 2 3
```

