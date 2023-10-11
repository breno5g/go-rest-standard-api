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

func Get(id int) entities.Task {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	var task entities.Task

	err := db.QueryRow("SELECT * FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return entities.Task{}
	}

	return task
}

func Delete(id int) {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM tasks WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
}

func Update(id int, title, description string) {
	db := infra.ConnectWithDatabase()
	defer db.Close()

	update, err := db.Prepare("UPDATE tasks SET title = $1, description = $2 WHERE id = $3")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(title, description, id)
}
