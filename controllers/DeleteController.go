
package controllers

import (
	"database/sql"
	"expenses/config"
	"expenses/middleware"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func DeleteExpense(db *sql.DB, store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	// Получаем ID авторизованного пользователя
	userID, ok := middleware.CurrentUserID(store, r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Некорректный ID расхода", http.StatusBadRequest)
		return
	}

	// Удаляем только расход текущего пользователя
	err = config.DeleteExpense(db, id, userID)
	if err != nil {
		http.Error(w, "Не удалось удалить расход", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
