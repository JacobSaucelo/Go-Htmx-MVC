package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	models "webGoCrud/models"
	"webGoCrud/utils"
)

type Users models.UsersType

var tmpl *template.Template

func DisplayUsers(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("views/index.html"))
	page := utils.GetUsers()
	fmt.Println("page: ", page)
	tmpl.Execute(w, page)

}

// func (u *Users) AddUser(w http.ResponseWriter, r http.Request) {
// 	name := r.PostFormValue("user-input-name")
// 	age := r.PostFormValue("user-input-age")

// 	tmpl = template.Must(template.ParseFiles("views/index.html"))
// }
