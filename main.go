package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dns := "jmilo67:Camilo_67@tcp(localhost:3306)/db_contacts"

	// Abrir una conexion a la base de datos
	db, err := sql.Open("mysql", dns)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexion a la base de datos exitosa")
}
