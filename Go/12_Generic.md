# Generic
Golang `1.18` version 이후로 새롭게 추가된 기능인 Generic 프로그래밍에 대해 알아보자.

제너릭은 타입 파라미터를 통해서 하나의 함수나 타입이 여러 타입에 대해서 동작하도록 해주는 프로그래밍 기법입니다. 자바나 C++, C#과 같은 다른 언어에서 
이미 제공되던 기능으로 Go에서도 기능이 생겼습니다.

일단, 기본적인 코드를 보겠습니다.
```go
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}
```
위의 코드는 맵의 값을 더하고 합계를 반환하는 두 함수를 선언합니다.
- `SumFloats`은 string 키값에 대한 float64 value 의 맵을 가져 옵니다 
- `SumInts`은 string 키값에 대한 int64 value 의 맵을 가져 옵니다

<br>

```go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

<br>
<br>

```go
// ProtoToModel proto file 을 Gorm Model 로 변환하여 반환
func (x *Node) ProtoToModel() (*model.PolicyNode, error) {
   b, err := json.Marshal(x)
   if err != nil {
	   return nil, errors.Unknown(err)
   }
   
    var m model.PolicyNode
    if err := json.Unmarshal(b, &m); err != nil {
        return nil, errors.Unknown(err)
    }

    return &m, nil
}
   
// ModelToProto Gorm Mode 을 proto file 로 변환하여 반환
func (x *Node) ModelToProto(m *model.PolicyNode) error {
   b, err := json.Marshal(m)
   if err != nil {
	   return errors.Unknown(err)
   }

   _ = json.Unmarshal(b, x)
   
   return nil
}
```

```go
func GeModel[T1 any, T2 any](from *T1, to *T2) error {
    b, err := json.Marshal(from)
    if err != nil {
        return err
    }
	
    if err = json.Unmarshal(b, to); err != nil {
        return err
    }

    return nil
}
```