package controllers

import (
	"database/sql"
	"expenses/config"
	"expenses/models"
	"fmt"
	"html/template"
	"net/http"
)

func ActionHome(db *sql.DB, w http.ResponseWriter, r *http.Request){
	category := r.URL.Query().Get("category")
	dateFrom := r.URL.Query().Get("from")
	dateTo := r.URL.Query().Get("to")

	tmpl, err := template.ParseFiles("./web/layout.html", "./web/home.html")

	if err != nil{
		http.Error(w, "Ошибка загрузки шаблона!", http.StatusInternalServerError)
		return
	}

	expenses, err := config.GetExpenses(db, category, dateFrom, dateTo)
	if err != nil {
		fmt.Println("Ошибка загрузки расходов:", err)
		http.Error(w, "Не удалось загрузить расходы!", http.StatusInternalServerError)
		return
	}

	categories, err := config.GetCategory(db)
	if err != nil {
		http.Error(w, "Не удалось загрузить категории!", http.StatusInternalServerError)
		return
	}

	data := models.HomeData{
		Expenses:   expenses,
		Categories: categories,
		SelectedCategory: category,
		DateFrom: dateFrom,
		DateTo: dateTo,
		
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Ошибка отображения страницы!", http.StatusInternalServerError)
		return
	}
	
	// fmt.Printf("Категория: %s, дата от: %s, дата до: %s\n", category, dateFrom, dateTo)
}