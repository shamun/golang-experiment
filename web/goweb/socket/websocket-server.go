package main

import (
  "fmt"
  "code.google.com/p/go.net/websocket"
  "log"
  "net/http"
)

func Echo(ws *websocket.Conn) {
  var err error

  for {
    var reply string

    if err = websocket.Message.Receive(ws, &reply); err != nil {
      fmt.Println("Can't receive")
      break
    }

    fmt.Println("Receive from client: " + reply)

    msg := "Received: " + reply
    fmt.Println("Sending to client: " + msg)

    if err = websocket.Message.Send(ws, msg); err != nil {
      fmt.Println("Cannot send")
      break
    }
  }
}

func main() {
  http.Handle("/", websocket.Handler(Echo))

  if err := http.ListenAndServe(":1234", nil); err != nil {
    log.Fatal(err)
  }
}
