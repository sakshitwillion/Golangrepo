package controller

import (
	"net/http"
	"strconv"

	"github.com/sakshitwillion/Golangrepo/model"
	"github.com/sakshitwillion/Golangrepo/view"
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
