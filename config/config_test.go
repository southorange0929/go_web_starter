package config

import "testing"

func TestConfig(t *testing.T) {
	appData := app{
		Mode: "debug",
		Port: "8080",
		Https: false,
	}

	Init()
	conf := Config
	if conf == nil {
		t.Errorf("Config init failed.")
	}
	application := Config.AppConfig
	if application != appData {
		t.Errorf("Config init failed")
	}
}
