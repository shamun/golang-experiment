package main

import (
    //"compress/gzip"
    "encoding/base64"
    "io"
    "os"
    "strings"
)

func main() {
    var r io.Reader
    data := "R29sYW5nIGlzIGF3ZXNvbWUhISBTbyBkbyBub2RlanMhIQ=="
    r = strings.NewReader(data)
    r = base64.NewDecoder(base64.StdEncoding, r)
    //r, _ = gzip.NewReader(r)
    io.Copy(os.Stdout, r)
}
