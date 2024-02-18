package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	addition       int = 1
	subtraction    int = 1
	multiplication int = 1
	division       int = 1
)

func calculate(expression string) (float64, error) {
	tokens := tokenize(expression)
	postfix, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}
	result, err := evaluatePostfix(postfix)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func tokenize(expression string) []string {
	// Удаляем все пробелы из выражения
	expression = strings.ReplaceAll(expression, " ", "")

	var tokens []string
	var buffer strings.Builder

	for _, char := range expression {
		if isOperator(char) {
			// Если встретили оператор, добавляем предыдущее число и оператор в список токенов,
			// а затем сбрасываем буфер
			if buffer.Len() > 0 {
				tokens = append(tokens, buffer.String())
				buffer.Reset()
			}
			tokens = append(tokens, string(char))
		} else if char == '(' || char == ')' {
			// Если встретили скобку, добавляем предыдущее число (если есть) и скобку в список токенов,
			// а затем сбрасываем буфер
			if buffer.Len() > 0 {
				tokens = append(tokens, buffer.String())
				buffer.Reset()
			}
			tokens = append(tokens, string(char))
		} else {
			buffer.WriteRune(char)
		}
	}

	// Добавляем последнее число в список токенов, если оно есть
	if buffer.Len() > 0 {
		tokens = append(tokens, buffer.String())
	}

	return tokens
}

func infixToPostfix(tokens []string) ([]string, error) {
	var postfix []string
	var stack []string

	for _, token := range tokens {
		switch token {
		case "+", "-":
			// Приоритет операций "+" и "-" равен 1,
			// поэтому выталкиваем из стека все операции с большим или равным приоритетом
			for len(stack) > 0 && (stack[len(stack)-1] == "+" || stack[len(stack)-1] == "-" || stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/") {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "*", "/":
			// Приоритет операций "*" и "/" равен 2,
			// поэтому выталкиваем из стека все операции с большим или равным приоритетом
			for len(stack) > 0 && (stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/") {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "(":
			stack = append(stack, token)
		case ")":
			// Выталкиваем все операции из стека в постфиксную форму до тех пор, пока не встретим открывающую скобку "("
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// Удаляем открывающую скобку "(" из стека
			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return nil, fmt.Errorf("Не найдена открывающая скобка")
			}
		default:
			postfix = append(postfix, token)
		}
	}

	// Выталкиваем все оставшиеся операции из стека в постфиксную форму
	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, fmt.Errorf("Не найдена закрывающая скобка")
		}
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix, nil
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func evaluatePostfix(tokens []string) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		switch token {
		case "+":
			if len(stack) < 2 {
				return 0, fmt.Errorf("Недостаточно операндов для операции сложения")
			}
			for i := 0; i < addition; i++ {
				time.Sleep(time.Second)
			}
			result := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "-":
			if len(stack) < 2 {
				return 0, fmt.Errorf("Недостаточно операндов для операции вычитания")
			}
			for i := 0; i < subtraction; i++ {
				time.Sleep(time.Second)
			}
			result := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "*":
			if len(stack) < 2 {
				return 0, fmt.Errorf("Недостаточно операндов для операции умножения")
			}
			for i := 0; i < multiplication; i++ {
				time.Sleep(time.Second)
			}
			result := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "/":
			if len(stack) < 2 {
				return 0, fmt.Errorf("Недостаточно операндов для операции деления")
			}
			if stack[len(stack)-1] == 0 {
				return 0, fmt.Errorf("Деление на ноль")
			}
			for i := 0; i < division; i++ {
				time.Sleep(time.Second)
			}
			result := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		default:
			number, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("Некорректный токен: %s", token)
			}
			stack = append(stack, number)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("Некорректное выражение")
	}

	return stack[0], nil
}

func IDConstructor(exp string) string {
	exp = strings.ReplaceAll(exp, " ", "")
	exp = strings.ReplaceAll(exp, "(", "")
	exp = strings.ReplaceAll(exp, ")", "")
	exp = strings.ReplaceAll(exp, "+", "0")
	exp = strings.ReplaceAll(exp, "-", "1")
	exp = strings.ReplaceAll(exp, "*", "2")
	exp = strings.ReplaceAll(exp, "/", "3")
	return exp
}

func CreateDatabase() {
	// Открываем соединение с базой данных
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создаем таблицу
	createTable := `
        CREATE TABLE IF NOT EXISTS answers (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            res REAL
        );
    `

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("База данных успешно создана!")
}

func InsertData(id string, res float64) {
	// Открываем соединение с базой данных
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Выполняем SQL-запрос для добавления записи
	insertQuery := `INSERT INTO answers (id, res) VALUES (?, ?)`
	_, err = db.Exec(insertQuery, id, res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Элемент успешно добавлен в базу данных!")
}

func FindByID(id string) (float64, bool) {
	// Открываем соединение с базой данных
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Выполняем SQL-запрос для поиска записи по id
	query := `SELECT id, res FROM answers WHERE id = ?`
	row := db.QueryRow(query, id)

	// Инициализируем переменные для хранения значений из базы данных
	var userID int
	var res float64

	// Сканируем результаты запроса в переменные
	err = row.Scan(&userID, &res)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Запись не найдена")
			return 0, false
		}
		log.Fatal(err)
	}

	return res, true
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/calculator.html")
	tmpl.Execute(w, "")

	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Ожидается выражение", http.StatusMethodNotAllowed)
		return
	}

	// Получаем тело запроса
	exp := r.FormValue("expression")
	id := IDConstructor(exp)
	_, f := FindByID(id)
	if f {
		w.WriteHeader(http.StatusOK)
		ans, _ := FindByID(id)
		ansstr := strconv.FormatFloat(ans, 'g', -1, 64)
		w.Write([]byte(ansstr))
		return
	} else {
		ans, err := calculate(exp)
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

func delaysHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/delays.html")
	tmpl.Execute(w, "")

	addition, err := strconv.Atoi(r.FormValue("addition"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err "))
	}
	subtraction, err = strconv.Atoi(r.FormValue("subtraction"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err "))
	}
	multiplication, err = strconv.Atoi(r.FormValue("multiplication"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err "))
	}
	division, err = strconv.Atoi(r.FormValue("division"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err "))
	}
	fmt.Println(addition, subtraction, multiplication, division)
	w.Write([]byte(string(addition)))
	w.Write([]byte(string(subtraction)))
	w.Write([]byte(string(multiplication)))
	w.Write([]byte(string(division)))
}

func main() {
	CreateDatabase()
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/static/calculator", calculatorHandler)
	http.HandleFunc("/static/delays", delaysHandler)
	http.ListenAndServe(":8080", nil)
}
