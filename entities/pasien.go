package entities

type Pasien struct {
	Id       int64
	Task     string `validate:"required" label:"Task"`
	Assignee string `validate:"required" label:"Assignee"`
	Deadline string `valicate:"required" label:"Deadline"`
	Action   string `validate:"required" label:"Action"`
}
