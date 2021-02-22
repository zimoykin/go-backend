package app

import (
	"log"

	"github.com/zimoykin/go-backend/config"
	"github.com/zimoykin/go-backend/database"
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
