package controller

import (
	"myapp/demo/mvc_golang/my-app/model"
	"myapp/demo/mvc_golang/my-app/view"
	"net/http"
	"strconv"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := model.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	view.ViewUser(w, user)
}
