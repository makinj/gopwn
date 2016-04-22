package remote

import (
  "strconv"
  "fmt"
  "net"
  "io"
)

type Remote struct {
  connection net.Conn
}

func remote(address string, port int) *Remote {
  remoteAddressWithPort := address
  remoteAddressWithPort += ":"
  remoteAddressWithPort += strconv.Itoa(port)

  fmt.Printf(remoteAddressWithPort)

  conn, err := net.Dial("tcp", remoteAddressWithPort)
  if err != nil{
    return nil
  }


  return &Remote{
    connection: conn,
  }
}

func (r *Remote) send(s string) {

}

/*
 * reads n bytes from the connection in `r`
 */

func (r *Remote) recv(n int) string {
  t := make([]byte, n)
  bytesRead, err := r.connection.Read(t)

  if err != nil {
    if err != io.EOF {
      fmt.Println("read error:", err)
    }
  }

  fmt.Println("got ", bytesRead, " bytes")

  return string(t)
}
