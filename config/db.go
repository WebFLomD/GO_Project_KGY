
package config

import (
	"database/sql"
	"expenses/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/personal_expenses")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Получение расходов конкретного пользователя
func GetExpenses(db *sql.DB, userID int, category string, dateFrom string, dateTo string) ([]models.Expense, error) {
	var conditions []string
	var args []interface{}

	// Обязательно учитываем владельца расходов
	conditions = append(conditions, "expenses.user_id = ?")
	args = append(args, userID)

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

	query := `SELECT expenses.id, expenses.description, expenses.amount,
				expenses.date, expenses.category_id, categories.name
				FROM expenses
				JOIN categories ON expenses.category_id = categories.id
				WHERE ` + strings.Join(conditions, " AND ")

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var expense models.Expense

		if err := rows.Scan(
			&expense.ID,
			&expense.Description,
			&expense.Amount,
			&expense.Date,
			&expense.CategoryID,
			&expense.CategoryName,
		); err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

// Получение конкретного расхода пользователя
func GetExpenseByID(db *sql.DB, id int, userID int) (models.Expense, error) {
	var expense models.Expense

	query := `SELECT id, description, amount, date, category_id
				FROM expenses
				WHERE id = ? AND user_id = ?`

	err := db.QueryRow(query, id, userID).Scan(
		&expense.ID,
		&expense.Description,
		&expense.Amount,
		&expense.Date,
		&expense.CategoryID,
	)

	if err != nil {
		return models.Expense{}, err
	}

	return expense, nil
}

// Добавление расхода пользователя
func AddExpense(db *sql.DB, userID int, description string, amount int, date string, categoryID int) error {
	query := `INSERT INTO expenses
				(description, amount, date, category_id, user_id)
				VALUES (?, ?, ?, ?, ?)`

	_, err := db.Exec(query, description, amount, date, categoryID, userID)

	return err
}

// Редактирование расхода пользователя
func UpdateExpense(db *sql.DB, id int, userID int, description string, amount int, date string, categoryID int) error {
	query := `UPDATE expenses
				SET description = ?, amount = ?, date = ?, category_id = ?
				WHERE id = ? AND user_id = ?`

	_, err := db.Exec(
		query,
		description,
		amount,
		date,
		categoryID,
		id,
		userID,
	)

	return err
}

// Удаление расхода пользователя
func DeleteExpense(db *sql.DB, id int, userID int) error {
	query := `DELETE FROM expenses WHERE id = ? AND user_id = ?`

	_, err := db.Exec(query, id, userID)

	return err
}

// Получение категорий
func GetCategory(db *sql.DB) ([]models.Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var category models.Category

		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
