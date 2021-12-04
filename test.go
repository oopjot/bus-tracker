package main

import (
  "fmt"
  "time"
)

func worker (channel chan string) {
  fmt.Println("Working")
  i := 0
  for {
    time.Sleep(time.Second)
    channel <- fmt.Sprintf("iteracja %d", i)
    i += 1
  }
}

func main() {
  fmt.Println("Kurwa")
  messages := make(chan string)

  go worker(messages)
  go func() {
    for msg := range messages {
      fmt.Println(msg)
    }
  }()

  var a string
  fmt.Scanf("itma", &a)
  fmt.Println(a)
}
