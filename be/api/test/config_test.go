package test

import (
	"fmt"
	"testing"

	"gapp.sys/config"
)

func TestConfig(t *testing.T) {
	cfgService := config.NewConfigService()
	appDir, err := cfgService.GetAppDir()
	if err != nil {
		t.Fatal("get app dir error", err)
	}
	fmt.Println(appDir)
}
func TestLoadYamlConfigFromFile(t *testing.T) {
	cfgService := config.NewConfigService()
	cfg, err := cfgService.LoadYamlConfigFromFile("./../config.yaml", "./../config.env")
	if err != nil {
		t.Fatal("load config error", err)
	}
	fmt.Println(cfg)
}
