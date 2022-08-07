package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("domain name required")
    return
  }

  domain := os.Args[1]
  nss, err := net.LookupNS(domain)
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, ns := range nss {
    fmt.Println(ns.Host)
  }
}
