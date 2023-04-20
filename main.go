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
	// host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")
	// fmt.Println(host)
	// fmt.Println(port)
}
