package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := make([]byte, 0, 100)
    str = strconv.AppendInt(str, 4567, 10)
    str = strconv.AppendBool(str, false)
    str = strconv.AppendQuote(str, "adcdef")
    str = strconv.AppendQuoteRune(str, '中')
    fmt.Println(string(str))

    fl := strconv.FormatFloat(123.23, 'g', 12, 64)
    fmt.Println(fl)

    ui, err := strconv.ParseUint("12345", 10, 64)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(ui + 10)
}
