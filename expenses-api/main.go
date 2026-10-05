package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Expense struct{
	Name string
	Price int
	Category string
}

var expenses []Expense

func addHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		tmpl, err := template.ParseFiles("templates/layout.html", "templates/add.html")

		if err != nil{
			fmt.Fprintln(w, "Ошибка загрузки шаблона") 
			return
		}

		tmpl.ExecuteTemplate(w, "layout", expenses)

		return
	}

	if r.Method != "POST"{
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	category := r.FormValue("category")

	if err != nil{
		fmt.Fprintln(w, "Ошибка загрузки") 
		return
	}

	expense := Expense{
		Name:     name,
		Price:    price,
		Category: category,
	}
	
	expenses = append(expenses, expense)

	http.Redirect(w, r, "/", 303)
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/home.html")

	if err != nil{
		fmt.Fprintln(w, "Ошибка загрузки шаблона") 
		return
	}

	tmpl.ExecuteTemplate(w, "layout", expenses)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "О нас")
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/add", addHandler)
	
	http.ListenAndServe(":8080", nil)
}