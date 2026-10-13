package controllers

import (
	"database/sql"
	"expenses/config"
	"html/template"
	"net/http"
	"strconv"
)

func EditExpense(db *sql.DB, w http.ResponseWriter, r *http.Request){
	
	switch r.Method{
		case http.MethodGet:
			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil{
				http.Error(w, "Некорректный ID", http.StatusBadRequest)
				return
			}

			expense, err := config.GetExpenseByID(db, id)
			if err != nil{
				http.Error(w, "Не удалось найти расход", http.StatusNotFound)
				return
			}

			tmpl, err := template.ParseFiles("./web/layout.html", "./web/edit.html")
			if err != nil {
				http.Error(w, "Ошибка загрузки шаблона!", http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "layout", expense)
			if err != nil{
				http.Error(w, "Ошибка отображения страницы", http.StatusInternalServerError)
				return
			}
		case http.MethodPost:
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil{
				http.Error(w, "Ошибка загрузки", http.StatusInternalServerError) 
				return
			}
			description := r.FormValue("description")
			amount, err := strconv.Atoi(r.FormValue("amount"))
			if err != nil{
				http.Error(w, "Ошибка загрузки", http.StatusInternalServerError) 
				return
			}
			date := r.FormValue("date")
			
			err = config.UpdateExpense(db, id, description, amount, date)
			if err != nil{
				http.Error(w, "Не удалось обновить данные", http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}