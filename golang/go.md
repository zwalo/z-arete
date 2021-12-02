# Golang

### 기본 문법   

1.데이터 타입

|자료형|저장범위|설명|
|------|---|---|
|uint||32비트 시스템-uint 32,64비트 시스템-uint64|
|uint8|0 ~ 255|부호 없는 8비트 정수형|
|uint16|0 ~ 65,535|부호 없는 16비트 정수형|
|uint32|0 ~ 4,294,967,295|부호 없는 32비트 정수형|
|uint64|0 ~ 18,446,744,073,709,551,615|부호 없는 64비트 정수형|
|int||32비트 시스템-int 32,64비트 시스템-int64|
|int8|-128 ~ 127|부호 있는 8비트 정수형|
|int16|-32,768 ~ 32,767|부호 있는 16비트 정수형|
|int32|-2,147,483,648 ~ 2,147,483,647|부호 있는 32비트 정수형|
|int64|-9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807|부호 있는 64비트 정수형|
|float32||IEEE-754 32비트 부동소수점, 7자리 정밀도|
|float64||IEEE-754 64비트 부동소수점, 12자리 정밀도|
|complex64||float32 크기의 실수부와 허수부로 구성된 복소수|
|complex128||float64 크기의 실수부와 허수부로 구성된 복소수|
|uintptr||uint와 같은 크기를 갖는 포인터형|
|bool||	참, 거짓을 표현하기 위한 8비트 자료형|
|byte||	8비트 자료형|
|rune||유니코드 저장을 위한 자료형, 크기는 int32와 동일|
|string||문자,문자열을 저장하기 위한 자료형|
   
2.변수 선언   

    기본 : var 변수명 타입
    상수 선언 : const 변수명 타입   
    
```go
     var num int32 
     const Pi float32 = 3.14   
```  
   
 ② 같은 타입을 가지는 변수를 여러 개 선언   
   ```go
    var num2, num3 int    
   ```

   ③ 여러 변수에 한 번에 값을 초기화 : 선언과 동시에 값을 초기화하면 타입을 명시할 필요가 없음
   ```go
    var num4, num5, str1 = 4, 5, "example"    
   ```
   ④ := 를 쓰면 var,const와 타입을 쓰지않고 선언과 동시에 값을 초기화(함수 안에서만 가능)  
   errorvar := str1   
   
   ⑤ 다른 타입을 가지는 변수를 여러 개 선언   
   ```go
    var (   
      i int   
      b bool   
      s string   
    )
```
```go
package main

import "fmt"

const ( 
	c1 = 10 //첫 번째 값은 무조건 초기화해야 함
	c2 // 선언되지 않은 값은 그 전에 선언된 값을 그대로 할당받음
	c3
	c4 = "goorm" //다른 형 선언 가능
	c5
	c6 = iota //c8까지 index값 저장
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func main() {
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)
    //10,10,10,goorm,goorm,5,6,7,earth,10,End
}    
```

3.zero value   
  선언만 하고 초기화하지 않은 변수는 zero value로 초기화   
  
  *선언 후 사용하지 않는 변수가 있다면 go는 에러를 낸다.

  숫자 타입 : 0   
  boolean 타입 : false   
  string 타입 : ""

  ```go   
  var (
    i int // zero value = 0
    f float64 // zero value = 0
    b bool // zero value = false
    s string // zero value = ""
)
```
   

4.함수 선언과 리턴
- 함수 선언은 func, 리턴 타입과 매개변수 타입은 뒤에 선언
```go
func add1(x int, y int) int {
    return x + y
}
```
- 매개변수 타입이 같을 경우 한번만 선언 가능
```go
func add1(x, y int) int {
    return x + y
}
```
- 여러값을 리턴 가능   
① return 뒤에 리턴 타입을 적어주는 방법   

```go
func divide1(dividend, divisor int) (int, int) {
    var quotient = (int)(dividend/divisor)
    var remainder = dividend%divisor
    return quotient, remainder
}
```
② return뒤에 리턴할 변수를 선언하는 방법
```go
func divide2(dividend, divisor int) (quotient, remainder int) {
    quotient = (int)(dividend/divisor)
    remainder = dividend%divisor
    return //return이라고만 적으면 미리 return값으로 정해 놓은 quotient와 remainder를 return합니다.
}
```

```go
import "fmt"

func main() {
    //①로 한 번에 여러개의 결과를 return받는 부분
    var quotient, remainder int
    quotient, remainder = divide1(10, 3)
    fmt.Println("①의 결과:", quotient, remainder)
    
    //②로 한 번에 여러개의 결과를 return받는 부분
    quotient, remainder = divide2(10, 3)
    fmt.Println("②의 결과:", quotient, remainder)
}
```