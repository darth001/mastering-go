package main

import (
  "fmt"
  "net/http"
  "net/http/pprof"
  "os"
  "time"
)

func someHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Serving %s\n", r.URL.Path)
  fmt.Printf("Served %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
  t := time.Now().Format(time.RFC1123)
  body := "The current time is:"
  fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", body)
  fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
  fmt.Fprintf(w, "Serving %s\n", r.URL.Path)
  fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
  PORT := ":8001"
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("Use default port:", PORT)
  } else {
    PORT = ":" + arguments[1]
    fmt.Println("Use specified port:", PORT)
  }

  r := http.NewServeMux()
  r.HandleFunc("/time", timeHandler)
  r.HandleFunc("/", someHandler)

  r.HandleFunc("/debug/pprof/", pprof.Index)
  r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
  r.HandleFunc("/debug/pprof/profile", pprof.Profile)
<<<<<<< HEAD
  r.HandleFunc("/debug/pprof/symbox", pprof.Symbol)
=======
  r.HandleFunc("/debug/pprof/symbox", pprof.Symbox)
>>>>>>> 2e0e20b875d167604db7c27a4f73b352b73849fb
  r.HandleFunc("/debug/pprof/trace", pprof.Trace)

  err := http.ListenAndServe(PORT, r)
  if err != nil {
    fmt.Println(err)
    return
  }
}
