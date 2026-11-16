
package main

import (
	"expenses/config"
	"expenses/controllers"
	"expenses/middleware"
	"log"
	"net/http"
	// "os"

	"github.com/gorilla/sessions"
)

func main() {
	// Подключение к базе данных
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	// Получение секретного ключа
	// secret := os.Getenv("SESSION_KEY")

	// if len(secret) < 32 {
	// 	log.Fatal("Задайте SESSION_KEY длиной не менее 32 символов")
	// }

	secret := "my-super-secret-key-for-local-development-123456"

	// Создание хранилища сессий
	store := sessions.NewCookieStore([]byte(secret))

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	// Регистрация и авторизация
	http.HandleFunc("/registration", controllers.RegistrationHandler(db, store))
	http.HandleFunc("/login", controllers.LoginHandler(db, store))
	http.HandleFunc("/logout", controllers.LogoutHandler(store))

	// Открытая страница
	http.HandleFunc("/about", controllers.ActionAbout)

	// Главная страница
	http.Handle("/", middleware.AuthRequired(store, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		controllers.ActionHome(db, store, w, r)
	})))

	// Добавление расхода
	http.Handle("/add", middleware.AuthRequired(store, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.ActionAddition(db, store, w, r)
	})))

	// Редактирование расхода
	http.Handle("/edit", middleware.AuthRequired(store, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.EditExpense(db, store, w, r)
	})))

	// Удаление расхода
	http.Handle("/delete", middleware.AuthRequired(store, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteExpense(db, store, w, r)
	})))

	// Запуск сервера
	log.Println("Сервер запущен: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
