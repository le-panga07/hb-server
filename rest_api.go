package main

import (
	"database/sql"
	"hb-server/config"
	"hb-server/controller/logger"

	"hb-server/controller/homecontroller"
	"net/http"

	_ "hb-server/github.com/go-sql-driver/mysql"
	"hb-server/github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	router := mux.NewRouter()
	db, _ := config.GetMySQLDB()

	fs := http.StripPrefix("/hb-server/static/", http.FileServer(http.Dir("../hb-server/static/")))
	router.PathPrefix("/hb-server/static/").Handler(fs)

	router.HandleFunc("/", redirect).Methods("GET")
	router.HandleFunc("/home", homecontroller.Index(db)).Methods("GET")
	router.HandleFunc("/home/{id}", homecontroller.GetConfigMap(db)).Methods("GET")

	router.HandleFunc("/providerResponseLog", logger.Log(db, "providerResponseLog")).Methods("POST")

	router.HandleFunc("/auctionParticipantLog", logger.Log(db, "auctionParticipantLog")).Methods("POST")

	router.HandleFunc("/auctionWinnerLog", logger.Log(db, "auctionWinnerLog")).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func redirect(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/home", http.StatusFound)
}
