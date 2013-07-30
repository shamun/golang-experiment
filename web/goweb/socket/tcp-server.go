package main

import (
  "fmt"
  "net"
  "os"
  "time"
)

func checkError(err error) {
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(0)
  }
}

func handleClient(conn net.Conn) {
  defer conn.Close()
  daytime := time.Now().String()
  conn.Write([]byte(daytime))
}

func main() {
  service := ":7777"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkError(err)
  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)
  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    go handleClient(conn)
  }
}
