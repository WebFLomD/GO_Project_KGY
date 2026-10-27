package config

import (
	"database/sql"
	"expenses/models"
	"strings"

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
func GetExpenses (db *sql.DB, category string, dateFrom string, dateTo string) ([]models.Expense, error){
	var conditions []string
	var args []interface{}

	// Фильтр по категории
    if category != "" {
        conditions = append(conditions, "expenses.category_id = ?")
        args = append(args, category)
    }

    // Фильтр по начальной дате
    if dateFrom != "" {
        conditions = append(conditions, "expenses.date >= ?")
        args = append(args, dateFrom)
    }

    // Фильтр по конечной дате
    if dateTo != "" {
        conditions = append(conditions, "expenses.date <= ?")
        args = append(args, dateTo)
    }

	// SQL запрос
	query := `SELECT expenses.id, expenses.description, expenses.amount,
				expenses.date, expenses.category_id, categories.name
        	FROM expenses
        	JOIN categories ON expenses.category_id = categories.id`

	if len(conditions) > 0 {
    	query += " WHERE " + strings.Join(conditions, " AND ")
    }

	rows, err := db.Query(query, args...)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense

	for rows.Next(){
		var expense models.Expense	
		
		if err := rows.Scan(&expense.ID, &expense.Description, &expense.Amount, &expense.Date, &expense.CategoryID, &expense.CategoryName); err != nil{
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

	row := db.QueryRow("SELECT id, description, amount, date, category_id FROM expenses WHERE id = ?", id)
	err := row.Scan(&expense.ID, &expense.Description, &expense.Amount, &expense.Date, &expense.CategoryID)

	if err != nil{
		return models.Expense{}, err
	}

	return expense, nil	
}

// Функция добавление в БД
func AddExpense(db *sql.DB, description string, amount int, date string, categoryID int) error{
	_, err := db.Exec("INSERT INTO expenses (description, amount, date, category_id) VALUES (?, ?, ?, ?)",
							description, amount, date, categoryID)

	if err != nil{
		return err
	}

	return nil
}

// Функция редактирование в БД
func UpdateExpense(db *sql.DB, id int, description string, amount int, date string, categoryID int) error {
    _, err := db.Exec(
        "UPDATE expenses SET description = ?, amount = ?, date = ?, category_id = ? WHERE id = ?",
        description, amount, date, categoryID, id,
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

// Функция получение данные с БД - Категория
func GetCategory (db *sql.DB) ([]models.Category, error){
	rows, err := db.Query("SELECT id, name from categories")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	
	for rows.Next(){
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil{
			return nil, err
		}

		categories = append(categories, category)
	}

	if  err = rows.Err(); err != nil{
		return nil, err
	}

	return categories, nil
}