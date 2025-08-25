package config

import (
	"github.com/cynx-io/cynx-core/src/configuration"
)

var Config *AppConfig

type AppConfig struct {
	Perintis struct {
		Email struct {
			From                   string `mapstructure:"from"`
			SubjectPreorderSuccess string `mapstructure:"subjectPreorderSuccess"`
		} `mapstructure:"email"`
		Social struct {
			Instagram string `mapstructure:"instagram"`
			Tiktok    string `mapstructure:"tiktok"`
			Facebook  string `mapstructure:"facebook"`
			Twitter   string `mapstructure:"twitter"`
			Website   string `mapstructure:"website"`
			Discord   string `mapstructure:"discord"`
		} `mapstructure:"social"`
	} `mapstructure:"perintis"`
	Aws struct {
		Region          string `mapstructure:"region"`
		AccessKeyID     string `mapstructure:"accessKeyId"`
		SecretAccessKey string `mapstructure:"secretAccessKey"`
	} `mapstructure:"aws"`
	Elastic struct {
		Url   string `json:"url"`
		Level string `json:"level"`
	} `json:"elastic"`
	Hermes struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"hermes"`
	Plutus struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"plutus"`
	App struct {
		Name    string `mapstructure:"name"`
		Address string `mapstructure:"address"`
		Key     string `mapstructure:"key"`
		Port    int    `mapstructure:"port"`
		Debug   bool   `mapstructure:"debug"`
	} `mapstructure:"app"`
	Database struct {
		Host        string `mapstructure:"host"`
		Database    string `mapstructure:"database"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		Dialect     string `mapstructure:"dialect"`
		AutoMigrate bool   `mapstructure:"autoMigrate"`
		Pool        struct {
			Max     int `mapstructure:"max"`
			Min     int `mapstructure:"min"`
			Acquire int `mapstructure:"acquire"`
			Idle    int `mapstructure:"idle"`
		} `mapstructure:"pool"`
		Port int `mapstructure:"port"`
	} `mapstructure:"database"`
}

func Init() {

	Config = &AppConfig{}
	err := configuration.InitConfig("config.json", Config)
	if err != nil {
		panic("failed to initialize config: " + err.Error())
	}
}
