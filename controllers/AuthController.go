package controllers

import (
    "database/sql"
    "expenses/config"
    "expenses/middleware"
    "net/http"

    "github.com/gorilla/sessions"
    "golang.org/x/crypto/bcrypt"
)

func RegistrationHandler(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            http.ServeFile(w, r, "./web/registration.html")
            return
        }

        if r.Method != http.MethodPost {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
            return
        }

        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password")

        if username == "" || email == "" || password == "" {
            http.Error(w, "Заполните все поля", http.StatusBadRequest)
            return
        }

        if len(password) < 8 {
            http.Error(w, "Пароль должен содержать минимум 8 символов", http.StatusBadRequest)
            return
        }

        passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            http.Error(w, "Ошибка обработки пароля", http.StatusInternalServerError)
            return
        }

        err = config.CreateUser(db, username, email, string(passwordHash))
        if err != nil {
            http.Error(w, "Не удалось зарегистрировать пользователя. Возможно, логин или почта уже заняты.", http.StatusBadRequest)
            return
        }

        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
}

func LoginHandler(db *sql.DB, store *sessions.CookieStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            http.ServeFile(w, r, "./web/login.html")
            return
        }

        if r.Method != http.MethodPost {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
            return
        }

        username := r.FormValue("username")
        password := r.FormValue("password")

        user, err := config.GetUserByUsername(db, username)
        if err != nil {
            if err == sql.ErrNoRows {
                http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
                return
            }

            http.Error(w, "Ошибка базы данных", http.StatusInternalServerError)
            return
        }

        err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
        if err != nil {
            http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
            return
        }

        session, err := store.Get(r, middleware.SessionName)
        if err != nil {
            http.Error(w, "Ошибка создания сессии", http.StatusInternalServerError)
            return
        }

        session.Values[middleware.UserIDKey] = user.ID

        err = session.Save(r, w)
        if err != nil {
            http.Error(w, "Не удалось сохранить сессию", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func LogoutHandler(store *sessions.CookieStore) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        session, err := store.Get(r, middleware.SessionName)
        if err == nil {
            session.Options.MaxAge = -1
            _ = session.Save(r, w)
        }

        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
}
