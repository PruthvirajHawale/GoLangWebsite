package controller

import (
	"BASIC/model"
	"BASIC/view"
	"encoding/json"
	"net/http"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := view.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some Error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.ReadById(name)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			json.NewEncoder(w).Encode(data)
		}else if r.Method == http.MethodDelete{
			name := r.URL.Path[1:]
			 err := model.DeleteTodo(name)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte("Record Deleted Successfully"))
		}
	}
}
