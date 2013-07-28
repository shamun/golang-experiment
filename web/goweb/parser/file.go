package main

import (
    "fmt"
    "os"
)

func main() {
    // create/remove directory
    os.Mkdir("test", 0777)
    os.MkdirAll("test/test1/test2", 0777)
    err := os.Remove("test")
    if err != nil {
        fmt.Println(err)
    }
    os.RemoveAll("test")

    // read/write file
    userfile := "test.txt"
    fout, err := os.Create(userfile)
    if err != nil {
        fmt.Println(userfile, err)
        return
    }
    defer fout.Close()
    for i := 0; i < 10; i++ {
        fout.WriteString("Just a test!\n")
        fout.Write([]byte("Just a test!\n"))
    }

    fl, err := os.Open(userfile)
    if err != nil {
        fmt.Println(userfile, err)
        return
    }
    defer fl.Close()
    buf := make([]byte, 1024)
    for {
        n, _ := fl.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
    }

    err = os.Remove(userfile)
    if err != nil {
        fmt.Println(userfile, err)
        return
    }
}
