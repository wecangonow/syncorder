package config

import (
	"flag"
	"time"

	"github.com/go-ini/ini"
)

type Config struct {
	HttpPort    string        `ini:"http_port"`
	TickSeconds time.Duration `ini:"tick_seconds"`
	Shopify
	Fetchx
	FilePath
}

type FilePath struct {
	LogPath    string `ini:"log_path"`
	DbFilePath string `ini:"db_file_path"`
}

type Shopify struct {
	AppKey       string `ini:"app_key"`
	AppPassword  string `ini:"app_password"`
	SharedSecret string `ini:"shared_secret"`
	ApiUrl       string `ini:"api_basic_url"`
}
type Fetchx struct {
	Authorization string `ini:"authorization"`
	TrackingURL   string `ini:"tracking_url"`
	ApiUrl        string `ini:"api_basic_url"`
}

var (
	AppConfig = new(Config)
)

func Init() {
	configFile := flag.String("c", "./config.ini", "configuration file")

	flag.Parse()

	cfg, _ := ini.Load(*configFile)

	cfg.MapTo(AppConfig)
}
