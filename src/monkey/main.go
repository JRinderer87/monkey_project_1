package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main(){
	user, err := user.Current()

	if err!=nil{
		panic(err)
	}

	fmt.Printf("Hello %s! This is the monkey programming language\n!",user.Username)

	fmt.Printf("Feel free to type in a command\n")

	repl.Start(os.Stdin,os.Stdout)


}