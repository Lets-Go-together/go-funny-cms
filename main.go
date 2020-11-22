package main

import (
	"gocms/pkg/command"
	"log"
	"os"
)

func main() {
	// step1： 随便写点什么
	app := command.InitApp()
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
