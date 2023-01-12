package main

import (
	"database/sql"
	"fmt"

	//"fmt"
	"log"

	"net/http"

	pHandler "github.com/Prodget/handler/product"
	pService "github.com/Prodget/service/product"
	m "github.com/Prodget/store/products"
	"github.com/gorilla/mux"

	//pStore "github.com/Prodget/store"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
)

func main() {
	conn := "root:#Dhiran77@tcp(localhost:3306)/gofr"
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Print("error while opening mysql ")
	}

	str := m.New(db) //okay?
	srvs := pService.New(str)
	hdlr := pHandler.HttpH{Srv: srvs}

	r := mux.NewRouter()
	//r.HandleFunc("/",hdlr.GetById)

	r.HandleFunc("/user/{id}", hdlr.GetById).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", hdlr.PutById).Methods(http.MethodPut)
	// r.HandleFunc("/user/{id}", hdlr.GetById).Methods(http.MethodGet)
	// r.HandleFunc("/user/{id}", hdlr.GetById).Methods(http.MethodGet)

	fmt.Println("Server starting at 8000")
	err = http.ListenAndServe(":8000", r)
	//println(err.Error())
}
