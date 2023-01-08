package main

import (
  "github.com/DoanRii/go-arkcon/arkcon"
  "fmt"
  "bufio"
  "os"
  "strings"
)


func main() {
  ark, err := arkrcon.NewARKRconConnection("162.156.93.87:31011", "f7jhdJs8xcekXCKP")
  if err != nil {
    fmt.Println(err)
    return
  }

  ark.Slomo(1)
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("RCON")
  fmt.Println("---------------------")

  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare("hi", text) == 0 {
      fmt.Println("hello, Yourself")
    }

  }

}