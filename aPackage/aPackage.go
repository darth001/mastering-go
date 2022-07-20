// The source code of a Go package, which can contain multiple files and multiple directories,
// can be found within a single directory that is named after the package name, with the obvious
// exception of the main package, which can be located anywhere.
package aPackage

import (
  "errors"
  "fmt"
  "math/rand"
  "time"
)

const MyConstant = 1334
const privateConstant = 21

func init() {
  rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
  // A slice of message format
  formats := []string {
    "Hi %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
  }
  
  return formats[rand.Intn(len(formats))]
}

func A() {
  fmt.Println("This is function A")
}

func B() {
  fmt.Println("privateConstant:", privateConstant)
}

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("empty name")
  }
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

func Hellos(names []string) (map[string]string, error) {
  messages := make(map[string]string)
  for _, name := range names {
    message, err := Hello(name)
    if err != nil {
      return nil, err
    }
    messages[name] = message
  }
  return messages, nil
}

