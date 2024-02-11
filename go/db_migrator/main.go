package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	migrator "codeforge/db_migrator/migrator"

	"codeforge/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error in loading config file in migration : %v ", err)
	}

	file, err := os.OpenFile("schema.sql", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	config.Writer = file
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", config.PgUser, config.PgPassword, config.PgHost, config.PgPort, config.PgDBName, config.SslMode)
	log.Println(psqlInfo)
	conn, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = conn.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("Connected")
	}
	defer conn.Close()
	migrations(conn)

	dumpErr := migrator.Dump()
	if dumpErr != nil {
		log.Fatal(dumpErr)
	}
}

func migrations(conn *sqlx.DB) {
	migration := flag.String("migration", "", "up - For doing up migration, down - For doing down migration")
	count := flag.Int("count", 0, "Number of migrations to run")
	flag.Parse()
	if *count <= 0 {
		*count++
	}
	migrator := migrator.ProvideMigrator(conn.DB, *migration, *count)
	_, err := migrator.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}
}
