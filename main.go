package main

import (
	. "./utils/constants"
	. "./app"
)

func main() {

	a := App{}
	// You need to set your Username and Password here
	a.Initialize(DB_USER, DB_PASS, DB_NAME)

	a.Run(API_HOST + ":" + API_PORT)
}







