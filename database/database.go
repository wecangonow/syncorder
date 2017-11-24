package database


import (
	scribble "github.com/nanobox-io/golang-scribble"
	"sync"
	"syncorder/config"
)

var db *scribble.Driver
var once sync.Once

// start loggeando
func GetInstance() *scribble.Driver {
	once.Do(func() {
		db, _ = scribble.New(config.AppConfig.FilePath.DbFilePath, nil)
	})
	return db

}

