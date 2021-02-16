package main

import (
	"fmt"
	"go_web_starter/config"
)

func main(){
	configs := config.Init()
	fmt.Print(configs)
}