
package middleware

import (
    "net/http"

    "github.com/gorilla/sessions"
)

const SessionName = "expenses_session"

const UserIDKey = "user_id"

// Проверяем, авторизован ли пользователь
func AuthRequired(store *sessions.CookieStore, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := store.Get(r, SessionName)

        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        userID, ok := session.Values[UserIDKey].(int)

        if !ok || userID == 0 {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// Получаем ID авторизованного пользователя
func CurrentUserID(store *sessions.CookieStore, r *http.Request) (int, bool) {
    session, err := store.Get(r, SessionName)

    if err != nil {
        return 0, false
    }

    userID, ok := session.Values[UserIDKey].(int)

    if !ok || userID == 0 {
        return 0, false
    }

    return userID, true
}
