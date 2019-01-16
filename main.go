package main

import (
	"smartcardWebServer/pkg/cmd"
	"os"
)

func main(){
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}












