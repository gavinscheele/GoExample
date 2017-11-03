package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pubgslotsExample/src/repository"
	"github.com/pubgslotsExample/src/repository/psql"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	env            *appEnv
	router         *mux.Router
	userRepository repository.UserRepository
}

type appEnv struct {
	DbHost     string `envconfig:"db_host" require:"true"`
	DbName     string `envconfig:"db_name" required:"true"`
	DbUsername string `envconfig:"db_username" required:"true"`
	DbPassword string `envconfig:"db_password" required:"true"`
	Port       int    `required:"true"`
}

func NewApp() *App {
	var env appEnv
	envconfig.MustProcess("example", &env)

	repository, err := psql.NewPsqlRepository(env.DbHost, env.DbName, env.DbUsername, env.DbPassword)
	if err != nil {
		log.Fatal(err)
	}

	app := App{
		env:            &env,
		router:         mux.NewRouter(),
		userRepository: repository,
	}
	app.InitializeRoutes()
	return &app
}

func (app *App) Close() {
	app.userRepository.Close()
}

func (app *App) InitializeRoutes() {
	app.router.HandleFunc("/newuser/", app.CreateUser).Methods("GET")
	app.router.HandleFunc("/user/{id:[0-9]+}", app.GetUser).Methods("GET")
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", app.env.Port), app.router))
}
