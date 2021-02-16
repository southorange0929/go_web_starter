package util

import "os"

/**
获取当前ENV环境变量
*/
func GetEnv() string {
	var env = "dev"
	osEnv := os.Getenv("ENV")
	if osEnv != "" {
		env = osEnv
	}
	return env
}
