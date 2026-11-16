package config

import (
    "database/sql"
    "expenses/models"
)

func CreateUser(db *sql.DB, username, email, passwordHash string) error {
    query := `
        INSERT INTO users (username, email, password_hash)
        VALUES (?, ?, ?)
    `

    _, err := db.Exec(query, username, email, passwordHash)
    return err
}

func GetUserByUsername(db *sql.DB, username string) (models.User, error) {
    var user models.User

    query := `
        SELECT id, username, email, password_hash
        FROM users
        WHERE username = ?
    `

    err := db.QueryRow(query, username).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.PasswordHash,
    )

    return user, err
}
