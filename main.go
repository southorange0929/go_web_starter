package main

import (
	"go_web_starter/cmd"
	"go_web_starter/util"
)

func main() {
	util.LogInit()
	cmd.WebServerStart()
}
