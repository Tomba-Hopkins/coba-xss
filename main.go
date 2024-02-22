package main

import (
	"coba-xss/handlers"
	"log"
	"net/http"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.DashboardHanlder)
	mux.HandleFunc("/parameter", handlers.ParameterHandler)
	
	fileserver := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	log.Println("Web started at port 8080")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// http.ListenAndServe("localhost:8080", mux) *atau gini


}