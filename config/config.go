package config

import (
	"go_web_starter/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	MySQL mysql `yaml: "mysql"`
}

type mysql struct {
	Host     string `yaml: "host"`
	Port     string `yaml: "port"`
	User     string `yaml: "user"`
	Password string `yaml: "password"`
	Database string `yaml: "database"`
}

var Config = &config{
	MySQL: mysql{Host: "127.0.0.1", Port: "3306", User: "root", Password: "admin", Database: "test"},
}

/**
根据当前环境变量ENV读取对应的配置文件
*/
func Init() *config {
	var conf = &config{}
	configFilePath := filepath.Join(os.Getenv("ROOTDIR"), "./conf/", util.GetEnv()+"_conf.yaml")
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Print(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Print(err)
	}
	return conf
}
