package main

import (
  "fmt"
  "math/rand"
  "net/http"
  "os"
  "time"
)

func random(min, max int) int {
  return rand.Intn(max - min) + min
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  delay := random(0, 15)
  time.Sleep(time.Duration(delay) * time.Second)
  fmt.Fprintf(w, "serving: %s\n", r.URL.Path)
  fmt.Fprintf(w, "delay: %d\n", delay)
  fmt.Printf("Served: %s\n", r.Host)
}

func main() {
  seed := time.Now().Unix()
  rand.Seed(seed)

  PORT := ":8001"
  args := os.Args
  if len(args) == 1 {
    fmt.Println("use default port", PORT)
  } else {
    PORT = ":" + args[1]
  }

  http.HandleFunc("/", rootHandler)
  err := http.ListenAndServe(PORT, nil)
  if err != nil {
    fmt.Println(err)
    return
  }
}
