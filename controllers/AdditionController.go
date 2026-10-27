package controllers

import (
	"database/sql"
	"expenses/config"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func ActionAddition (db *sql.DB, w http.ResponseWriter, r *http.Request){
	// GET
	if r.Method == http.MethodGet{
		tmpl, err := template.ParseFiles("./web/layout.html", "./web/add.html")
		if err != nil{
			fmt.Fprintln(w, "Ошибка загрузки шаблона!")
			return
		}

		categories, err := config.GetCategory(db)
		if err != nil {
			http.Error(w, "Не удалось загрузить данные!", http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, "layout", categories)
		return
	}

	// POST
	if r.Method != http.MethodPost{
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	description := r.FormValue("description")
	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil{
		fmt.Fprintln(w, "Ошибка загрузки") 
		return
	}
	date := r.FormValue("date")
	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil{
		fmt.Fprintln(w, "Ошибка загрузки") 
		return
	}

	err = config.AddExpense(db, description, amount, date, categoryID)
	if err != nil{
		http.Error(w, "Не удалось добавить расход", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}