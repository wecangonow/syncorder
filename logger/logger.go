package logger

import (
	"log"
	"os"
	"sync"
	"syncorder/config"
)

type Logger struct {
	filename string
	*log.Logger
}

var logger *Logger
var once sync.Once

// start loggeando
func GetInstance() *Logger {
	once.Do(func() {
		logger = createLogger(config.AppConfig.FilePath.LogPath)
	})
	return logger

}

func createLogger(fname string) *Logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &Logger{
		filename: fname,
		Logger:   log.New(file, "My app Name ", log.Lshortfile),
	}

}
