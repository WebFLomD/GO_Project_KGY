package main

import (
	"fmt"
	"net/http"
)

// Главная страница
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "Welcome to Expenses API!")
}

// Страница About
func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "Expenses API")
	fmt.Fprintln(w, "Simple HTTP server written in Go.")
	fmt.Fprintln(w, "This project will be used to learn how to build a REST API.")
}

// Ping
func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "pong")
}

func main() {
	// Создаём встроенный маршрутизатор Go
	mux := http.NewServeMux()

	// Регистрируем маршруты
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /about", about)
	mux.HandleFunc("GET /ping", ping)

	// Сообщение в терминале
	fmt.Println("Server started on http://localhost:8080")

	// Запускаем HTTP-сервер
	err := http.ListenAndServe(":8080", mux)

	// Если сервер завершился с ошибкой
	if err != nil {
		fmt.Println("Server error:", err)
	}
}