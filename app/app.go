package app

import (
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"../routes"
	"log"
	"net/http"
	. "../repository"
	"fmt"
	"database/sql"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(user, password, dbname string) {

	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	DBCon , err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = routes.NewRouter()

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}