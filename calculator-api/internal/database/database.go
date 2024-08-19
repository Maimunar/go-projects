package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Calculation struct {
	Id        string
	Input     string
	Output    int
	CreatedAt string
}

type SQLiteRepository struct {
	db *sql.DB
	// Migrate*()
	// AddCalculation()
	// GetCalculations()
}

const dbName = "./internal/database/calculator.db"

func OpenDB() SQLiteRepository {
	fmt.Println("Opening database...")
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal("Error opening database", err)
	}
	return SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Close() {
	fmt.Println("Closing database...")
	err := r.db.Close()
	if err != nil {
		log.Fatal("Error closing database", err)
	}
}

func (r *SQLiteRepository) Migrate() {
	query := `
		CREATE TABLE IF NOT EXISTS calculations(
			id TEXT PRIMARY KEY, 
			input TEXT NOT NULL, 
			output INTEGER NOT NULL,
			createdAt TEXT
		);
    `

	_, err := r.db.Exec(query)
	if err != nil {
		log.Fatal("Error migrating database", err)
	}
}

func (r *SQLiteRepository) AddCalculation(input string, output int) error {
	stmt, err := r.db.Prepare("INSERT INTO calculations (id, input, output, createdAt) VALUES (?, ?, ?, datetime('now')) ")
	if err != nil {
		return err
	}
	id := uuid.New().String()

	_, err = stmt.Exec(id, input, output)
	defer stmt.Close()

	return err
}

func (r *SQLiteRepository) GetCalculations() ([]Calculation, error) {
	rows, err := r.db.Query("SELECT id, input, output, createdAt FROM calculations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	calculations := []Calculation{}
	for rows.Next() {
		var c Calculation

		err := rows.Scan(&c.Id, &c.Input, &c.Output, &c.CreatedAt)
		if err != nil {
			return nil, err
		}

		calculations = append(calculations, c)
	}
	return calculations, nil
}
