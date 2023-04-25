package main

import (
	"cli/src/command"
	"fmt"
)

func main() {

	rootCmd := command.InitParser()
	error := rootCmd.Execute()
	if error != nil {
		fmt.Println(error.Error())
	}
}
