package config

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Address     string
	Port        string
	DB_Driver   string
	DB_Name     string
	DB_Address  string
	DB_Port     string
	DB_Username string
	DB_Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = InitConfig()
	}

	return appConfig
}

func InitConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Address = getEnv("ADDRESS", "http://localhost")
	defaultConfig.Port = getEnv("PORT", "3000")
	defaultConfig.DB_Driver = getEnv("DB_DRIVER", "mysql")
	defaultConfig.DB_Name = getEnv("DB_NAME", "todolist")
	defaultConfig.DB_Address = getEnv("DB_ADDRESS", "localhost")
	defaultConfig.DB_Port = getEnv("DB_PORT", "3307")
	defaultConfig.DB_Username = getEnv("DB_USERNAME", "root")
	defaultConfig.DB_Password = getEnv("DB_PASSWORD", "")

	fmt.Println(defaultConfig)
	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
