package main

import "fmt"

// Person は人間を表す構造体。
type Person struct {
    Name string
    Age  int
}

func main() {

    // 値として、pに代入
    p := Person{
        Name: "太郎",
        Age:  20,
    }

    fmt.Printf("最初のp :%+v\n", p)

    p2 := p
    p2.Name = "二郎"
    p2.Age = 21
    // pではなく
    fmt.Printf("p2で二郎に書き換えを行なったはずのp :%+v\n", p)

    // &pで*Person(Personのポインタ型)を生成する
    // p3はpのアドレスが格納されている状態になる
    p3 := &p
    p3.Name = "二郎"
    p3.Age = 21

    fmt.Printf("p3で二郎に書き換えを行なったp :%+v\n", p)
}
