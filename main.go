package main

import (
	"SDK/cli"
	"fmt"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
