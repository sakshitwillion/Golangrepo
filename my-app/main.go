package main

import (
	"net/http"

	"github.com/sakshitwillion/Golangrepo/controller"
)

func main() {
	http.HandleFunc("/user", controller.HandleUser)
	http.ListenAndServe(":8080", nil)
}
