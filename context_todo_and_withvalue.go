package main

import (
  "context"
  "fmt"
)

type key string

func searchKey(ctx context.Context, k key) {
  v := ctx.Value(k)
  if v != nil {
    fmt.Println("found value:", v)
    return
  } else {
    fmt.Println("key not found:", k)
  }
}

func main() {
  akey := key("some-secret-value")
  ctx := context.WithValue(context.Background(), akey, "some-secret-value")
  searchKey(ctx, akey)
  searchKey(ctx, key("key-not-found"))
  emptyCtx := context.TODO()
  searchKey(emptyCtx, key("notfound"))
}
