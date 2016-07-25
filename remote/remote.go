package remote

import (
  "strconv"
  "fmt"
  "net"
  "io"
)

type RemoteObj struct {
  connection net.Conn
}

func Remote(address string, port int) *RemoteObj {
  remoteAddressWithPort := address
  remoteAddressWithPort += ":"
  remoteAddressWithPort += strconv.Itoa(port)

  conn, err := net.Dial("tcp", remoteAddressWithPort)
  if err != nil{
    return nil
  }

  return &RemoteObj {
    connection: conn,
  }
}

func (r *RemoteObj) Send(s string) {
  _, err := r.connection.Write([]byte(s))

  if err != nil {
    fmt.Println("write error:", err)
  }

  return
}

/*
 * reads n bytes from the connection in `r`
 */

func (r *RemoteObj) Recv(n int) string {
  t := make([]byte, n)
  _, err := r.connection.Read(t)

  if err != nil {
    if err != io.EOF {
      fmt.Println("read error:", err)
    }
  }

  return string(t)
}
