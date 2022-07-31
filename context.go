package main

import (
  "context"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "strconv"
  "sync"
  "time"
)

var (
  url string
  delay int = 5
  w sync.WaitGroup
)

type response struct {
  r *http.Response
  err error
}

func connect(c context.Context) error {
  defer w.Done()

  res := make(chan response, 1)

  tr := &http.Transport{}
  httpClient := &http.Client{Transport: tr}

  req, _ := http.NewRequest("GET", url, nil)

  go func() {
    r, err := httpClient.Do(req)
    if err != nil {
      fmt.Println(err)
      res <- response{nil, err}
      return
    } else {
      pack := response{r, err}
      res <- pack
    }
  }()

  select {
  case <-c.Done():
    tr.CancelRequest(req)
    <-res
    fmt.Println("Request cancelled.")
    return c.Err()
  case ok := <-res:
    err := ok.err
    resp := ok.r
    if err != nil {
      fmt.Println("Error select:", err)
      return err
    }
    defer resp.Body.Close()

    realHttpData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Println("Error select:", err)
      return err
    }
    fmt.Printf("server response: %s\n", realHttpData)
  }

  return nil
}

func main() {
  if len(os.Args) == 1 {
    fmt.Println("url and delay required")
    return
  }

  url = os.Args[1]

  if len(os.Args) == 3 {
    t, err := strconv.Atoi(os.Args[2])
    if err != nil {
      fmt.Println(err)
      return
    }
    delay = t
    fmt.Println("delay:", delay)
  }

  c := context.Background()
  c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
  defer cancel()

  fmt.Printf("connect to: %s\n", url)
  w.Add(1)
  go connect(c)
  w.Wait()
  fmt.Println("exit.")
}
