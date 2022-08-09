package models

import (
	"database/sql"
	"fmt"
	"go-crud-master/config"
	"go-crud-master/entities"
	"time"
)

type TaskModel struct {
	conn *sql.DB
}

func NewTaskModel() *TaskModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &TaskModel{
		conn: conn,
	}
}

func (p *TaskModel) FindAll() ([]entities.Task, error) {

	rows, err := p.conn.Query("select * from task")
	if err != nil {
		return []entities.Task{}, err
	}
	defer rows.Close()

	var dataTask []entities.Task
	for rows.Next() {
		var task entities.Task
		rows.Scan(&task.Id,
			&task.Task,
			&task.Assignee,
			&task.Deadline,
			&task.Action)

		if task.Action == "1" {
			task.Action = "Sudah Selesai"
		} else {
			task.Action = "WIP"
		}
		// 2006-01-02 => yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", task.Deadline)
		// 02-01-2006 => dd-mm-yyyy
		task.Deadline = tgl_lahir.Format("02-01-2006")

		dataTask = append(dataTask, task)
	}

	return dataTask, nil

}

func (p *TaskModel) Create(task entities.Task) bool {

	result, err := p.conn.Exec("insert into task (id, task, assignee, deadline, action) values(?,?,?,?,?)",
		task.Id, task.Task, task.Assignee, task.Deadline, task.Action)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *TaskModel) Find(id int64, task *entities.Task) error {

	return p.conn.QueryRow("select * from task where id = ?", id).Scan(
		&task.Id,
		&task.Task,
		&task.Assignee,
		&task.Deadline,
		&task.Action)
}

func (p *TaskModel) Update(task entities.Task) error {

	_, err := p.conn.Exec(
		"update task set task = ?, assignee = ?, deadline = ?, action = ? where id = ?",
		task.Task, task.Assignee, task.Deadline, task.Action, task.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *TaskModel) Delete(id int64) {
	p.conn.Exec("delete from task where id = ?", id)
}
