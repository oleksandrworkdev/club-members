package app

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func renderIndexPage(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles(indexPage))
	data := IndexPageData{
		Members: getMembers(),
		ErrorText: errorText,
	}
	tmpl.Execute(w, data)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Println("Error", err)
			renderIndexPage(w)
		}
		fName := r.PostForm["name"][0]
		fEmail := r.PostForm["email"][0]

		emailExist := doesEmailExist(fEmail)
		if !emailExist {
			registrationDate := time.Now().Format("02.01.2006")
			members = append(members, Member{Name: fName, Email: fEmail, RegistrationDate: registrationDate})
			errorText = ""
		} else {
			errorText = "Such email already exist"
		}

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		log.Printf("Serving %s to %s...\n", indexPage, r.RemoteAddr)
		renderIndexPage(w)
	}
}