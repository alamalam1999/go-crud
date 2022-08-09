package main

import (
	"go-crud-master/controllers/taskcontroller"
	"net/http"
)

func main() {

	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/task", taskcontroller.Index)
	http.HandleFunc("/task/index", taskcontroller.Index)
	http.HandleFunc("/task/add", taskcontroller.Add)
	http.HandleFunc("/task/edit", taskcontroller.Edit)
	http.HandleFunc("/task/delete", taskcontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
