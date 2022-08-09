package taskcontroller

import (
	"go-crud-master/entities"
	"go-crud-master/libraries"
	"go-crud-master/models"
	"html/template"
	"net/http"
	"strconv"
)

var validation = libraries.NewValidation()
var taskModel = models.NewTaskModel()

func Index(response http.ResponseWriter, request *http.Request) {

	task, _ := taskModel.FindAll()

	data := map[string]interface{}{
		"task": task,
	}

	temp, err := template.ParseFiles("views/task/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/task/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var task entities.Task
		task.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		task.Task = request.Form.Get("task")
		task.Assignee = request.Form.Get("assignee")
		task.Deadline = request.Form.Get("deadline")
		task.Action = request.Form.Get("action")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(task)

		if vErrors != nil {
			data["task"] = task
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data task berhasil disimpan"
			taskModel.Create(task)
		}

		temp, _ := template.ParseFiles("views/task/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var task entities.Task
		taskModel.Find(id, &task)

		data := map[string]interface{}{
			"task": task,
		}

		temp, err := template.ParseFiles("views/task/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var task entities.Task
		task.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		task.Task = request.Form.Get("task")
		task.Assignee = request.Form.Get("assignee")
		task.Deadline = request.Form.Get("deadline")
		task.Action = request.Form.Get("action")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(task)

		if vErrors != nil {
			data["task"] = task
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data  berhasil diperbarui"
			taskModel.Update(task)
		}

		temp, _ := template.ParseFiles("views/task/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	taskModel.Delete(id)

	http.Redirect(response, request, "/task", http.StatusSeeOther)
}
