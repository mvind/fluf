package main

import(
  "fmt"
  // "path/filepath"
  "github.com/senseyeio/roger"

)

// We need to create to dynamicly start a R enviroment to where we can run
// run.Rserve() using the Rserve package ???
// And then we run this file

// Learn about type structures in golang
// See the returned object 

func main() {
  // p := "./p/p"
  // q := ".\\q\\q"
  //
  // fmt.Println(path.base("C:\\Users\\morte_000\\github\\"))
  // WIN package root
  // win_path := "C:\Users\morte_000\github\fluf\r_backend"
  //
  // fmt.Println(win_path)
  rClient, err := roger.NewRClient("127.0.0.1", 6311)
  if err != nil {
    fmt.Println("Failed to connect")
    return
  }

  value, err := rClient.Eval("source(\"C:/Users/morte_000/github/fluf/r_backend/test123.R\")")
  if err != nil {
    fmt.Printf("Command failed: " + err.Error())
  }

  
}
