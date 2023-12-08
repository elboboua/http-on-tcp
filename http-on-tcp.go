package main

import (
  "fmt"
  "net"
)

func HandleRequest(conn net.Conn) {
    defer conn.Close()
    fmt.Printf("Received connection '%s'\n", conn.RemoteAddr().String())
    request := make([]byte,256)
    _, err := conn.Read(request)
    fmt.Println(string(request))

    _, err = conn.Write([]byte("HTTP/1.1 200 OK \r\n\r\n"))
    if err != nil {
      fmt.Println(err)
    }
}

func main() {
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Listening on port 8080")

  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println(err)
    }
    go HandleRequest(conn)
  }
}
