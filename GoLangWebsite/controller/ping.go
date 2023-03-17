package controller

import (
	s "BASIC/view"
	"encoding/json"
	"fmt"
	"net/http"
)

func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := s.Response{
				Code: http.StatusOK,
				Body: "Pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
			fmt.Println("Data Inserted Succesfully")
		}
	}
}
