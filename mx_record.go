package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("domain named required")
    return
  }

  domain := arguments[1]
  mxs, err := net.LookupMX(domain)
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, mx := range mxs {
    fmt.Println(mx.Host)
  }
}
