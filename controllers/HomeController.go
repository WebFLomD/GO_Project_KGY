package controllers

import (
	"database/sql"
	"expenses/config"
	"fmt"
	"html/template"
	"net/http"
)

func ActionHome(db *sql.DB, w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("./web/layout.html", "./web/home.html")

	if err != nil{
		http.Error(w, "Ошибка загрузки шаблона!", http.StatusInternalServerError)
		return
	}

	expenses, err := config.GetExpenses(db)
	if err != nil{
		http.Error(w, "Не удалось загрузить расходы!", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", expenses)
	fmt.Printf("%+v\n", expenses)
}