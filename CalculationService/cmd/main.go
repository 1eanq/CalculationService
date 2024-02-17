package main

import (
	"CalculationService/internal/database"
	handler "CalculationService/internal/handler"
	"net/http"
)

func main() {
	database.CreateDatabase()
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/static/calculator", handler.CalculatorHandler)
	http.HandleFunc("/static/delays", handler.DelaysHandler)
	http.ListenAndServe(":8080", nil)
}
