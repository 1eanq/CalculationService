package CalculationService

import (
	calculator "CalculationService/internal/calculator"
	"CalculationService/internal/database"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/calculator.html")
	tmpl.Execute(w, "")

	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Ожидается выражение", http.StatusMethodNotAllowed)
		return
	}

	// Получаем тело запроса
	exp := r.FormValue("expression")
	id := database.IDConstructor(exp)
	_, f := database.FindByID(id)

	//Проверка, был ли запрос с таким же ID
	if f {
		w.WriteHeader(http.StatusOK)
		ans, _ := database.FindByID(id)
		ansstr := strconv.FormatFloat(ans, 'g', -1, 64)
		w.Write([]byte(ansstr))
		return
	} else {
		ans, err := calculator.Calculate(exp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		ansstr := strconv.FormatFloat(ans, 'g', -1, 64)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ansstr))
	}
	fmt.Println("Запрос обработан")
}

func DelaysHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/delays.html")
	tmpl.Execute(w, "")

	err := fmt.Errorf("err")
	calculator.Addition, err = strconv.Atoi(r.FormValue("addition"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err "))
	}
	calculator.Subtraction, err = strconv.Atoi(r.FormValue("subtraction"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err "))
	}
	calculator.Multiplication, err = strconv.Atoi(r.FormValue("multiplication"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err "))
	}
	calculator.Division, err = strconv.Atoi(r.FormValue("division"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err "))
	}
}
