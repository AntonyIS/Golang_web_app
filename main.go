package main

import (
	"fmt"
	"net/http"

	"github.com/AntonyIS/Golang_web_app/handler"
)

func main() {
	handler.MainHandler()
	fmt.Println("Server running on port 5000")
	http.ListenAndServe(":5000", nil)
}
