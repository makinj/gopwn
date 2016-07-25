package main

import (
  "fmt"
  "github.com/makinj/gopwn/remote"
)

func main() {
  basicConnection()
}

func basicConnection() {
  r := remote.Remote("reddit.com", 443)
  r.Send("GET / HTTP/1.0\r\n\r\n")
  fmt.Println(r.Recv(1024))
}