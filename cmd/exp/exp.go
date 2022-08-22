package main

import (
	"database/sql"
	"fmt"

	"github.com/choongmanee/lenslocked/models"
	"github.com/google/uuid"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
}

func main() {
	cfg := models.DefaultPostgresConfig()

	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL,
			amount INT,
			description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	// id := uuid.Must(uuid.NewRandom())
	// name := "Jon Calhoun"
	// email := "jon@calhoun.io"

	// row := db.QueryRow(`
	// 	INSERT INTO users(id, name, email)
	// 	VALUES ($1, $2, $3) RETURNING id;`, id, name, email)

	// var returnedId uuid.UUID
	// err = row.Scan(&returnedId)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("User Created. userId: %s", returnedId)

	row := db.QueryRow(`
		SELECT id, email from users
		WHERE "name" = $1;`, "Chung Calhoun")

	var id, email string
	err = row.Scan(&id, &email)
	if err == sql.ErrNoRows {
		fmt.Println("Error, no rows!")
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("user information: id=%s email=%s", id, email)

	userID := id
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	id := uuid.Must(uuid.NewRandom())
	// 	_, err := db.Exec(`
	// 		INSERT INTO orders(id, user_id, amount, description)
	// 		VALUES($1, $2, $3, $4);`, id, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders")

	type Order struct {
		ID          uuid.UUID
		UserID      string
		Amount      int
		Description string
	}

	var orders []Order

	rows, err := db.Query(`
		SELECT id, amount, description
		FROM orders
		WHERE user_id=$1`, userID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserID = userID
		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}

	fmt.Println("Orders:", orders)

}
