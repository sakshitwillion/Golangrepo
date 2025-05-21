package view

import (
	"html/template"
	"myapp/demo/mvc_golang/my-app/model"
	"net/http"
)

func ViewUser(w http.ResponseWriter, user model.User) {
	tmpl, err := template.New("user").Parse(`
        <h1>User Details</h1>
        <p>ID: {{.ID}}</p>
        <p>Name: {{.Name}}</p>
    `)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
