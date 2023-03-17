package main

import (
	"BASIC/controller"
	"BASIC/model"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	http.ListenAndServe(":8000", mux)
}
