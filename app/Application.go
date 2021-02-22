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

func (app *Application) Start() {

	log.Println("started")

}

func (app *Application) init() Application {
	return Application{
		Host: "", Port: "",
	}
}
