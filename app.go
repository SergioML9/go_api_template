package ileave_server


import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"ileave_server/model"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}


func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(moonfishteam.com:24396)/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()

	c := model.Content{}
	
}

func (a *App) initializeRoutes() {
	//a.Router.HandleFunc("/content", )
}

func (a *App) Run(addr string) { }