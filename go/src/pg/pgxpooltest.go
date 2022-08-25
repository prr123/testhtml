package main

import (
	"context"
	"fmt"
	"os"
//	"github.com/jackc/pgx/v4/pgxpool"
	pgx "github.com/jackc/pgx/v5"
)

func main() {
	var dbname string

	numArgs := len(os.Args)

	dbg := false
	switch numArgs {
		case 1:
			fmt.Println("error insufficient cmdline args!")
			fmt.Printf("  usage is: pgxpool dbname [dbg]!")
			os.Exit(-1)

		case 2:
			dbname = os.Args[1]
			if dbname == "dbg" {
				fmt.Println("error 'dbg' is not a valid dbname!")
				os.Exit(-1)
			}

		case 3:
			dbname = os.Args[1]
			if dbname == "dbg" {
				fmt.Println("error 'dbg' is not a valid dbname!")
				os.Exit(-1)
			}
			dbgresp := os.Args[2]
			if dbgresp == "dbg" {
				dbg = true
			} else {
				fmt.Printf("error '%s' is not a valid option!\n", dbgresp)
				os.Exit(-1)
			}


		default:
			fmt.Println("error too many cmdline args!")
			fmt.Printf("  usage is: pgxpool dbname [dbg]!")
			os.Exit(-1)

	}

	fmt.Printf("pgxpool dbg: %t\n",dbg)

	ctx := context.Background()

	connString := "postgresql://azuldbusr:azul@localhost:5433/azultestdb"

	if dbg {
		fmt.Printf("connString: %s\n", connString)
		config, err := pgx.ParseConfig(connString)
//		fmt.Printf("config: %v \nerr: %v\n", config, err)
//		fmt.Printf("%v\n",*config)
		if err != nil {
			fmt.Printf("error parsing connection string: %v\n", err)
			os.Exit(-1)
		}
		fmt.Printf("*** parsed connection string successfully ***\nConfiguration Parameters:\n")
		fmt.Printf("Host: %s\n",(*config).Host)
		fmt.Printf("Port: %d\n",(*config).Port)
		fmt.Printf("DB:   %s\n",(*config).Database)
		fmt.Printf("User: %s\n",(*config).User)
		fmt.Printf("Pwd:  %s\n",(*config).Password)
//		fmt.Printf("TLS:  %v\n",(*config).TLSConfig)
		fmt.Printf("Runtime Parameters: %d\n", len((*config).RuntimeParams))
		for k,v := range (*config).RuntimeParams {
			fmt.Printf("key: %s value: %s\n", k, v)
		}

	}
	dbpool, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close(ctx)

	fmt.Printf("\n*** connected to db ***\n")

	var greeting string
	queryStr := "select 'Hello, world!'"
	fmt.Printf("\nDb query: %s\n", queryStr)

	err = dbpool.QueryRow(context.Background(), queryStr).Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Query response: %s\n", greeting)

	fmt.Printf("\n*** success ***\n")
}
