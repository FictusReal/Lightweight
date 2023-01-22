package database

import (
	"context"
	"log"

	//"os"

	//"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"

	"github.com/jackc/pgx/v4"
)

func insertRows(ctx context.Context, tx pgx.Tx, user User) error {
	// Insert four rows into the "accounts" table.
	log.Println("Creating new rows...")
	if _, err := tx.Exec(ctx,
		"INSERT INTO users (id, display_name) VALUES ($1, $2)", user.ID, user.Name); err != nil {
		return err
	}
	return nil
}

func printData(conn *pgx.Conn) error {
	rows, err := conn.Query(context.Background(), "SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s: %s\n", id, name)
	}
	return nil
}

func deleteAccountRows(ctx context.Context, tx pgx.Tx, one string, two string) error {
	// Delete two rows into the "accounts" table.
	log.Printf("Deleting rows with IDs %s and %s...", one, two)
	if _, err := tx.Exec(ctx,
		"DELETE FROM accounts WHERE id IN ($1, $2)", one, two); err != nil {
		return err
	}
	return nil
}
