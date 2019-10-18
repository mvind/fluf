package main

import (
	"fmt"
	"os/exec"
	//	"github.com/senseyeio/roger"
)

/*
TODO / IDEAS:
We need a way to install Rserve if its not currently installed
Get pid number to kill daemon thread
1 socket pr instance, note eval function fucks with Rserve with async eval calls
*/

func main() {
	//var query string
	//query = "FX USD JPN" // test

	// Setup Rserve
	//r := "R CMD Rserve"
	//pid := []string{"ax", "|", "grep", "Rserve"}

	//cmd := exec.Command("/bin/ps", "ax", "|")
	cmd := exec.Command("/tmp/Rserv", "R CMD Rserve")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))

	// fmt.Println("GETTING: " + query)
	// rClient, err := roger.NewRClient("127.0.0.1", 6311)
	// if err != nil {
	// 	fmt.Println("Failed to connect")
	// 	return
	// }

	// value, err := rClient.Eval("source(\"C:/Users/morte_000/github/fluf/r_backend/test123.R\")")
	// if err != nil {
	// 	fmt.Printf("Command failed: " + err.Error())
	// }
	// // git setup
	// //n := len(value)
	// fmt.Println(value)

}
