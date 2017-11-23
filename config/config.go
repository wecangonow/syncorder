package config

import (
	"flag"
	"github.com/go-ini/ini"
	"log"
)

type Config struct {
	Shopify
	Fetchx
	FilePath
}

type FilePath struct {
	LogPath string `ini:"log_path"`
}

type Shopify struct {
	AppKey       string `ini:"app_key"`
	AppPassword  string `ini:"app_password"`
	SharedSecret string `ini:"shared_secret"`
}
type Fetchx struct {
	Authorization string `ini:"authorization"`
}

var (
	AppConfig = new(Config)
)

func Init() {
	configFile := flag.String("c", "./config.ini", "configuration file")

	flag.Parse()

	cfg, _ := ini.Load(*configFile)

	cfg.MapTo(AppConfig)

	log.Printf("%v", AppConfig)

}
