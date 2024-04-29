package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jcav67/go-mysql/database"
	"github.com/jcav67/go-mysql/handlers"
	"github.com/jcav67/go-mysql/models"
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
	newContact := models.Contact{Name: "usuarioGo2", Email: "go2@go.com", Phone: "123465"}

	handlers.ListarContacts(db)
	handlers.GetContactById(db, 3)
	handlers.CreateContact(db, newContact)

	fmt.Println("Conexion a la base de datos exitosa")

}
