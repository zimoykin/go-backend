package app

import (
	"log"

	"github.com/zimoykin/golnag1/config"
	"github.com/zimoykin/golnag1/database"
)

type Application struct {
	Host   string
	Port   string
	Config *config.Configuration
	db     *database.Configuration
}

func (s *Application) Start() {

	log.Println("started")

}
