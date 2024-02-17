package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

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
