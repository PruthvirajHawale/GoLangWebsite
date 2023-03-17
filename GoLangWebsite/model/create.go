package model

import "fmt"

func CreateTodo(name string, taskTodo string) error {
	insertQ, err := con.Query("insert into Todo values($1,$2)", name, taskTodo)
	if err != nil {
		fmt.Println("There was error in the Insert Query ........................................................")
		fmt.Println(err)
		fmt.Println("........................................................................")

	}
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}


func DeleteTodo(name string) error {
	insertQ, err := con.Query("delete from todo where name=$1", name)
	if err != nil {
		fmt.Println("There was error in the delete Query ........................................................")
		fmt.Println(err)
		fmt.Println("........................................................................")

	}
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}
