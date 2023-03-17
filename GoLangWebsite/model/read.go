package model

import (
	"BASIC/view"
	"fmt"
)

func ReadAll() ([]view.PostRequest, error) {
	rows, err := con.Query("select * from Todo")
	if err != nil {
		fmt.Println("Errror while fetching data")
		return nil, err
	}
	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}

func ReadById(name string) ([]view.PostRequest, error) {
	rows, err := con.Query("select * from Todo where name=$1", name)
	fmt.Println(name)
	if err != nil {
		fmt.Println("Errror while fetching data in ReadById")
		fmt.Print(err)
		return nil, err
	}
	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}
