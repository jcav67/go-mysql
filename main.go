package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jcav67/go-mysql/database"
	"github.com/jcav67/go-mysql/handlers"
)

func main() {

	db, err := database.Db_Connection()
	if err != nil {
		log.Fatal("ERROR ACA")
		log.Fatal("ERROR IN DATABASE CONNECTION")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
		log.Fatal("ERROR IN DATABASE CONNECTION")
	}

	handlers.ListarContacts(db)

	fmt.Println("Conexion a la base de datos exitosa")

}
