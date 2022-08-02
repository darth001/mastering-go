package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  "net/http"
  "os"
  "time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Listen: %s\n", r.URL.Path)
  fmt.Printf("Listen: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
  t := time.Now().Format(time.RFC1123)
  body := "The current time is:"
  fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", body)
  fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>", t)
  fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
  fmt.Printf("Served time for: %s\n", r.Host)
}

func getData(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Serving: %s\n", r.URL.Path)
  fmt.Printf("Served: %s\n", r.Host)

  connStr := "user=postgres password=postgres dbname=s2 sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", err)
    return
  }

  rows, err := db.Query("select * from users")
  if err != nil {
    fmt.Fprintf(w, "<h2 align=\"center\">%s</h3>\n", err)
    return
  }
  defer rows.Close()
  
  for rows.Next() {
    var id int
    var firstName string
    var lastName string
    err = rows.Scan(&id, &firstName, &lastName)
    if err != nil {
      fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", err)
      return
    }
    fmt.Fprintf(w, "<h3 align=\"center\">%d, %s, %s</h3>\n", id, firstName, lastName)
  }

  err = rows.Err()
  if err != nil {
    fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", err)
    return
  }
}

func main() {
  PORT := ":8001"
  arguments := os.Args
  if len(arguments) != 1 {
    PORT = ":" + arguments[1]
  }
  fmt.Println("port number:", PORT)
  
  http.HandleFunc("/time", timeHandler)
  http.HandleFunc("/getData", getData)
  http.HandleFunc("/", rootHandler)

  err := http.ListenAndServe(PORT, nil)
  if err != nil {
    fmt.Println(err)
    return
  }
}
