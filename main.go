package main

import(
	"github.com/gorilla/mux"
	"net/http"
	"API/driver"
	c "API/controller"
	"database/sql")

type api struct {
	Db *sql.DB
}

func main(){
	
	r:=mux.NewRouter()

	db :=driver.ConnectDB()

	a := &api{Db:db}

	r.HandleFunc("/create",c.CreateHandler(a.Db)).Methods("POST")
	r.HandleFunc("/list/limit={limit}&offset={offset}",c.ListHandler(a.Db)).Methods("GET")
	r.HandleFunc("/update/id={id}",c.UpdateHandler(a.Db)).Methods("PUT")
	r.HandleFunc("/delete/id={id}",c.DeleteHandler(a.Db)).Methods("DELETE")

	http.ListenAndServe(":8080",r)
}