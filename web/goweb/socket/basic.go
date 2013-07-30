package main

import (
  "net"
  "os"
  "fmt"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "Usage %s id-addr\n", os.Args[0])
    os.Exit(1)
  }

  name := os.Args[1]
  addr := net.ParseIP(name)
  if addr == nil {
    fmt.Println("Invalid")
  } else {
    fmt.Println("address is ", addr.String())
  }
  os.Exit(0)
}
