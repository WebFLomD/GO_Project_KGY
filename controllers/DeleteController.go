package controllers

import (
	"database/sql"
	"expenses/config"
	"net/http"
	"strconv"
)

func DeleteExpense(db *sql.DB, w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil{
		http.Error(w, "Некорректный ID расхода", http.StatusBadRequest) 
		return
	}

	err = config.DeleteExpense(db, id)
	if err != nil{
		http.Error(w, "Не удалось удалить расход", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}