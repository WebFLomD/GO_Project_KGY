package controllers

import (
	"fmt"
	"net/http"
)

func ActionAbout (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "О нас")
}