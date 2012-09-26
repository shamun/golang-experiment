package main

//import "fmt"
import (
    "math/cmplx"
    "fmt"
)

var (
    ToBe bool = false
    MaxInt uint64 = 1<<64 - 1
    z complex128 = cmplx.Sqrt(-5+12i)
)

// Structs
type Vertex struct {
    X int
    Y int
}

type VertexMap struct {
    Lat, Long float64
}

var m map[string]VertexMap

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4/9
    y = sum - x
    return
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)

    fmt.Println(split(17))

    // for loop
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
    sum = 1
    for sum < 50 {
        sum += sum
    }
    fmt.Println(sum)

    // Basic types
    const f = "%T(%v)\n"
    fmt.Println(f, ToBe, ToBe)
    fmt.Println(f, MaxInt, MaxInt)
    fmt.Println(f, z, z)

    //structs
    fmt.Println(Vertex{1, 2})
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v)

    //Pointer
    p := Vertex{1, 2}
    q := &p
    fmt.Println(q)
    q.X = 1e9
    fmt.Println(p)

    //Map
    m = make(map[string]VertexMap)
    m["Bell Labs"] = VertexMap{
        40.68433, 74.39967,
    }
    fmt.Println(m["Bell Labs"])
}
