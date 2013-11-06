package main

import "fmt"

type Stringer interface {
    String() string
}

func main() {
    var s Stringer
    v := Vertex{3, 4}
    c := Celsius(100)

    s = v
    s = c
    fmt.Println(v, c, s)
}

type Vertex struct { X, Y float64 }
func (v Vertex) String() string {
    return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

type Celsius float32
func (c Celsius) String() string {
    return fmt.Sprintf("%.2fC", c)
}
