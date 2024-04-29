package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jcav67/go-mysql/models"
)

// Listar todos los contactos desde la base de datos
func ListarContacts(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"

	// Ejecutar la consulta
	rowsResult, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rowsResult.Close()

	// Iterar sobre los resultados y mostrarlos
	fmt.Println("\nLista de contactos:")
	fmt.Println("---------------------------------------------------------------")
	for rowsResult.Next() {
		// Instancia del modelo contact
		contact := models.Contact{}
		var valueEmail sql.NullString
		err := rowsResult.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "No Email"
		}

		fmt.Printf("ID %d, Nombre: %s, email: %s, telefono: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("---------------------------------------------------------------")
	}
}
