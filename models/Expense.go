package models

type Expense struct {
	ID            int
	Description   string
	Amount        int
	Date          string
	CategoryID    int
	CategoryName  string
}

type EditExpenseData struct{
	Expense     Expense
	Categories  []Category
}