package main

import (
    "fmt"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
    // bug
    //src, err := os.Open(srcName)
    //if err != nil {
    //    return
    //}

    //dst, err := os.Create(dstName)
    //if err != nil {
    //    return
    //}

    //written, err = io.Copy(dst, src)
    //dst.Close()
    //src.Close()
    //return

    // use defer
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
