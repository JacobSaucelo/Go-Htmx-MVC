package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	models "webGoCrud/models"
	"webGoCrud/utils"
)

type Users models.UsersType

// var tmpl *template.Template
var tmpl = template.Must(template.ParseFiles("views/index.html"))

func DisplayUsers(w http.ResponseWriter, r *http.Request) {
	page := utils.GetUsers()

	fmt.Println("page: ", page)
	tmpl.Execute(w, page)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	userData := utils.GetUsers()
	name := r.PostFormValue("user-name")
	age := r.PostFormValue("user-age")

	response := userData.CreateUser(name, age)

	// tmpl = template.Must(template.ParseFiles("views/index.html"))
	tmpl.ExecuteTemplate(w, "users-list-element", models.UserType{
		Id:   response.Id,
		Name: response.Name,
		Age:  response.Age,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userData := utils.GetUsers()
	userId := r.FormValue("id")
	userData.DeleteUser(userId)

	// fmt.Println("userId: ", response)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// func (u *Users) AddUser(w http.ResponseWriter, r http.Request) {
// 	name := r.PostFormValue("user-input-name")
// 	age := r.PostFormValue("user-input-age")

// 	tmpl = template.Must(template.ParseFiles("views/index.html"))
// }
