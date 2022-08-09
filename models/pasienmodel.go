package models

import (
	"database/sql"
	"fmt"
	"go-crud-master/config"
	"go-crud-master/entities"
	"time"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel) FindAll() ([]entities.Pasien, error) {

	rows, err := p.conn.Query("select * from task")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPasien []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(&pasien.Id,
			&pasien.Task,
			&pasien.Assignee,
			&pasien.Deadline,
			&pasien.Action)

		if pasien.Action == "1" {
			pasien.Action = "Sudah Selesai"
		} else {
			pasien.Action = "WIP"
		}
		// 2006-01-02 => yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", pasien.Deadline)
		// 02-01-2006 => dd-mm-yyyy
		pasien.Deadline = tgl_lahir.Format("02-01-2006")

		dataPasien = append(dataPasien, pasien)
	}

	return dataPasien, nil

}

func (p *PasienModel) Create(pasien entities.Pasien) bool {

	result, err := p.conn.Exec("insert into task (id, task, assignee, deadline, action) values(?,?,?,?,?)",
		pasien.Id, pasien.Task, pasien.Assignee, pasien.Deadline, pasien.Action)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *PasienModel) Find(id int64, pasien *entities.Pasien) error {

	return p.conn.QueryRow("select * from task where id = ?", id).Scan(
		&pasien.Id,
		&pasien.Task,
		&pasien.Assignee,
		&pasien.Deadline,
		&pasien.Action)
}

func (p *PasienModel) Update(pasien entities.Pasien) error {

	_, err := p.conn.Exec(
		"update task set task = ?, assignee = ?, deadline = ?, action = ? where id = ?",
		pasien.Task, pasien.Assignee, pasien.Deadline, pasien.Action, pasien.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *PasienModel) Delete(id int64) {
	p.conn.Exec("delete from task where id = ?", id)
}
