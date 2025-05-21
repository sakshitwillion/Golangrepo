package main

import (
	"myapp/demo/mvc_golang/my-app/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/user", controller.HandleUser)
	http.ListenAndServe(":8080", nil)
}
