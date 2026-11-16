
package controllers

import (
	"database/sql"
	"expenses/config"
	"expenses/middleware"
	"expenses/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func EditExpense(db *sql.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	// Получаем ID авторизованного пользователя
	userID, ok := middleware.CurrentUserID(store, r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Некорректный ID", http.StatusBadRequest)
			return
		}

		// Получаем расход только текущего пользователя
		expense, err := config.GetExpenseByID(db, id, userID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Расход не найден", http.StatusNotFound)
				return
			}

			http.Error(w, "Не удалось найти расход", http.StatusInternalServerError)
			return
		}

		categories, err := config.GetCategory(db)
		if err != nil {
			http.Error(w, "Не удалось загрузить категории", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("./web/layout.html", "./web/edit.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки шаблона!", http.StatusInternalServerError)
			return
		}

		data := models.EditExpenseData{
			Expense:    expense,
			Categories: categories,
		}

		err = tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			http.Error(w, "Ошибка отображения страницы", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Некорректный ID", http.StatusBadRequest)
			return
		}

		description := r.FormValue("description")

		amount, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			http.Error(w, "Некорректная сумма", http.StatusBadRequest)
			return
		}

		date := r.FormValue("date")

		categoryID, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, "Некорректная категория", http.StatusBadRequest)
			return
		}

		// Обновляем только расход текущего пользователя
		err = config.UpdateExpense(db, id, userID, description, amount, date, categoryID)
		if err != nil {
			http.Error(w, "Не удалось обновить данные", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
