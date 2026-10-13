package config

import (
	"database/sql"
	"expenses/models"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error){
	// Открываем подключение к БД
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/personal_expenses")
	if err != nil {
		return nil, err
	}

	// Проверяем, доступна ли БД
	err = db.Ping()
	if err != nil{
		db.Close()
		return nil, err
	}

	// Возвращаем подключение
	return db, nil	
}

// Функция получение данные с БД
// Пример кода: https://go.dev/doc/database/querying
func GetExpenses (db *sql.DB) ([]models.Expense, error){
	rows, err := db.Query("SELECT id, description, amount, date FROM expenses")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense

	for rows.Next(){
		var expense models.Expense	
		
		if err := rows.Scan(&expense.ID, &expense.Description, &expense.Amount, &expense.Date); err != nil{
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return expenses, nil
}

// Функция получение ID
func GetExpenseByID(db *sql.DB, id int)(models.Expense, error){
	var expense models.Expense

	row := db.QueryRow("SELECT id, description, amount, date FROM expenses WHERE id = ?", id)
	err := row.Scan(&expense.ID, &expense.Description, &expense.Amount, &expense.Date)

	if err != nil{
		return models.Expense{}, err
	}

	return expense, nil	
}

// Функция добавление в БД
func AddExpense(db *sql.DB, description string, amount int, date string) error{
	_, err := db.Exec("INSERT INTO expenses (description, amount, date) VALUES (?, ?, ?)",
							description, amount, date)

	if err != nil{
		return err
	}

	return nil
}

// Функция редактирование в БД
// func UpdateExpense(db *sql.DB, id int, description string, amount int, date string) error{
// 	_, err := db.Exec("UPDATE expenses SET description = ?, amount = ?, date = ? WHERE id = ?", description, amount, date, id)
	
// 	if err != nil{
// 		return err
// 	}
// 	return nil
// }

func UpdateExpense(db *sql.DB, id int, description string, amount int, date string) error {
    _, err := db.Exec(
        "UPDATE expenses SET description = ?, amount = ?, date = ? WHERE id = ?",
        description, amount, date, id,
    )
    return err
}
// Функция удаление данные с БД
func DeleteExpense (db *sql.DB, id int) error{
	_, err := db.Exec("DELETE FROM expenses WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}