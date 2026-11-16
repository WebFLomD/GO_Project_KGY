
package controllers

import (
	"database/sql"
	"expenses/config"
	"expenses/middleware"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func ActionAddition(db *sql.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	// Получаем ID авторизованного пользователя
	userID, ok := middleware.CurrentUserID(store, r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// GET
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/layout.html", "./web/add.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки шаблона!", http.StatusInternalServerError)
			return
		}

		categories, err := config.GetCategory(db)
		if err != nil {
			http.Error(w, "Не удалось загрузить данные!", http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "layout", categories)
		if err != nil {
			http.Error(w, "Ошибка отображения страницы!", http.StatusInternalServerError)
			return
		}

		return
	}

	// POST
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	description := r.FormValue("description")

	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil {
		fmt.Fprintln(w, "Ошибка загрузки")
		return
	}

	date := r.FormValue("date")

	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		fmt.Fprintln(w, "Ошибка загрузки")
		return
	}

	// Добавляем расход с ID владельца
	err = config.AddExpense(db, userID, description, amount, date, categoryID)
	if err != nil {
		http.Error(w, "Не удалось добавить расход", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
