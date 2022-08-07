package main

import (
  "fmt"
  "net"
  "os"
)

func lookupIP(address string) ([]string, error) {
  hosts, err := net.LookupAddr(address)
  if err != nil {
    return nil, err
  }
  return hosts, nil
}

func lookupHostname(ip string) ([]string, error) {
  ips, err := net.LookupHost(ip)
  if err != nil {
    return nil, err
  }
  return ips, nil
}

func main() {
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("ip or hostname required")
    return
  }

  ip := net.ParseIP(arguments[1])
  if ip == nil {
    ips, err := lookupHostname(arguments[1])
    if err == nil {
      for _, sip := range ips {
        fmt.Println(sip)
      }
    }
  } else {
    hosts, err := lookupIP(arguments[1])
    if err == nil {
      for _, hostname := range hosts {
        fmt.Println(hostname)
      }
    }
  }
}
