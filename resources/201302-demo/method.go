package main

import "fmt"

type Celsius float32
type Fahrenheit float32

func (c Celsius) String() string {
    return fmt.Sprintf("%.2fC", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%.2fF", f)
}

func (c Celsius) Fahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func main() {
    c := Celsius(100)
    fmt.Println(c, "==", c.Fahrenheit())

}
