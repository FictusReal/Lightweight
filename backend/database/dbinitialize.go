package database

import (
	"context"
	"log"
	"os"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v5"
)

// func DBInit() {
// 	dsn := "postgresql://azaan:G5y2e6a8ojMkDnNoPA_tLA@sea-jackal-4744.6wr.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
// 	ctx := context.Background()
// 	conn, err := pgx.Connect(ctx, dsn)
// 	defer conn.Close(context.Background())
// 	if err != nil {
// 		log.Fatal("failed to connect database", err)
// 	}
// 	//fmt.Println("WE Are NOW Initializing Databse!!!")

// 	var now time.Time
// 	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
// 	if err != nil {
// 		log.Fatal("failed to execute query", err)
// 	}

// 	fmt.Println(now)

// }

func DBInit() {

	// Read in connection string
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Insert initial rows
	var accounts User
	// for i := 0; i < len(accounts); i++ {
	// 	account
	// }
	accounts.ID = "1"
	accounts.Name = "Zoeb"

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return insertRows(context.Background(), tx, accounts)
	})

	if err == nil {
		log.Println("New rows created.")
	} else {
		log.Fatal("error: ", err)
	}

	// Print out the balances
	log.Println("Initial balances:")
	printBalances(conn)

}
