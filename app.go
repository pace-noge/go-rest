package gorest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var Apps []App

type App interface {
	Name() string
	Description() string
	Setup() error
	Register(r *chi.Mux)
}

type BaseApp struct {
	Title        string
	Info         string
	RouteHandler func(r *chi.Mux)
	SetupHandler func() error
	Middlewares  []func(http.Handler) http.Handler
}

func (app *BaseApp) Name() string {
	return app.Title
}

func (app *BaseApp) Description() string {
	return app.Info
}

func (app *BaseApp) Setup() error {
	fmt.Printf("Configuring the %v app\n", app.Name())
	err := app.SetupHandler()
	if err != nil {
		return err
	}
	return nil
}

func (app *BaseApp) Register(r *chi.Mux) {
	app.RouteHandler(r)
}

func SetupApps() {
	fmt.Println("Configuring apps...")
	for _, app := range Apps {
		err := app.Setup()
		if err != nil {
			log.Fatal("unable to configure apss ", err.Error())
		}
	}
}

func RegisterApps(r *chi.Mux) {
	registerInternalUrls(r)
	fmt.Println("Registering apps...")
	for _, app := range Apps {
		app.Register(r)
	}
}
