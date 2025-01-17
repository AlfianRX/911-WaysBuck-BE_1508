package main

import (
	"fmt"
	"net/http"
	"waysbuck/database"
	"waysbuck/pkg/mysql"
	"waysbuck/routes"

	"github.com/gorilla/mux"
)

func main() {

	//. initial DB
	mysql.DatabaseInit()

	//. run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running at port 5000")
	http.ListenAndServe("localhost:5000", r)
}
