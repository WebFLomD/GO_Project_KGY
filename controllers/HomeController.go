
package controllers

import (
	"database/sql"
	"expenses/config"
	"expenses/middleware"
	"expenses/models"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

func ActionHome(db *sql.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	// Получаем ID авторизованного пользователя
	userID, ok := middleware.CurrentUserID(store, r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Получаем параметры фильтрации
	category := r.URL.Query().Get("category")
	dateFrom := r.URL.Query().Get("from")
	dateTo := r.URL.Query().Get("to")

	// Загружаем шаблоны
	tmpl, err := template.ParseFiles("./web/layout.html", "./web/home.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона!", http.StatusInternalServerError)
		return
	}

	// Получаем расходы только текущего пользователя
	expenses, err := config.GetExpenses(db, userID, category, dateFrom, dateTo)
	if err != nil {
		fmt.Println("Ошибка загрузки расходов:", err)
		http.Error(w, "Не удалось загрузить расходы!", http.StatusInternalServerError)
		return
	}

	// Получаем категории
	categories, err := config.GetCategory(db)
	if err != nil {
		http.Error(w, "Не удалось загрузить категории!", http.StatusInternalServerError)
		return
	}

	// Подготавливаем данные для шаблона
	data := models.HomeData{
		Expenses:         expenses,
		Categories:       categories,
		SelectedCategory: category,
		DateFrom:         dateFrom,
		DateTo:           dateTo,
	}

	// Отображаем страницу
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Ошибка отображения страницы!", http.StatusInternalServerError)
		return
	}
}
