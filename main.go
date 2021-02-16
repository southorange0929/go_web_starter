package main

import (
	"go_web_starter/config"
	"go_web_starter/dao"
	"log"
)

func main() {
	configs := config.Init()
	log.Print(configs)
	dao.Init()
	data := dao.GetPerson()
	log.Print(data)
}
