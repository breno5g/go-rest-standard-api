package model

import (
	"github.com/breno5g/rest-standard-go-api/entities"
	"github.com/breno5g/rest-standard-go-api/infra"
)

func List() []entities.Task {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err.Error())
	}

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description)
		if err != nil {
			panic(err.Error())
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func Create(title, description string) {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO tasks(title, description) VALUES($1, $2)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(title, description)
}
