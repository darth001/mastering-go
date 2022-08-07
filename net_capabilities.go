package main

import (
  "fmt"
  "net"
)

func main() {
  interfaces, err := net.Interfaces()
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, i := range interfaces {
    fmt.Printf("%v\n", i.Name)
    fmt.Printf("      flags: %s\n",i.Flags.String())
    fmt.Printf("        mtu: %d\n", i.MTU)
    fmt.Printf("mac address: %s\n", i.HardwareAddr.String())
  }
}
