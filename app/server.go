package app

import (
	"log"
	"net/http"
	"os"
)

const indexPage = "public/index.html"
var members []Member
var errorText string


func StartHttpServer() {
	members = []Member{
		{Name: "Bob", Email: "bob@mail.com", RegistrationDate: "17.01.2019"},
		{Name: "Jane", Email: "jane@mail.com", RegistrationDate: "22.04.2020"},
		{Name: "Sally", Email: "sally@mail.com", RegistrationDate: "07.03.2021"},
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	f, _ := os.Create("/var/log/golang/golang-server.log")
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/", handleIndex)

	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}
