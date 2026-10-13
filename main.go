package main

import (
	"expenses/config"
	"expenses/controllers"
	"log"
	"net/http"
)

func main(){
	// Подключение к БД
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()
	log.Println("Подключение к MySQL успешно!")
	
	// Страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.ActionHome(db, w, r) // Файл controllers/HomeControlle.go
	}) 
	http.HandleFunc("/about", controllers.ActionAbout) // Файл controllers/AboutControlle.go
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		controllers.ActionAddition(db, w, r) // Файл controllers/AdditionControlle.go
	})
	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		controllers.EditExpense(db, w, r) // Файл controllers/EditControlle.go
	}) 
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteExpense(db, w, r) // Файл controllers/DeleteControlle.go
	}) 
	
	// localhost:8080
	http.ListenAndServe(":8080", nil)
}