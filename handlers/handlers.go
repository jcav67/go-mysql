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

// Obtener un contacto por Id
func GetContactById(db *sql.DB, contactID int) {
	//Query para buscar  un contacto
	query := "SELECT * FROM contact where id = ?"

	// Ejecutar la consulta
	rowsResult := db.QueryRow(query, contactID)

	//instanciacion del modelo
	contact := models.Contact{}
	var valueEmail sql.NullString

	err := rowsResult.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontró registro con el ID %d", contactID)
		} else {
			log.Fatal("Error en la base de datos")
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "No Email"
	}
	fmt.Println("\nInformación del contacto:")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("ID %d, Nombre: %s, email: %s, telefono: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("---------------------------------------------------------------")
}

// Crear un nuevo contacto
func CreateContact(db *sql.DB, newContact models.Contact) {
	//Query para insertar nuevo contacto
	query := "INSERT INTO contact(name, email, phone) VALUES(?,?,?)"

	//Ejectuar la sentencia SQL
	res, err := db.Exec(query, newContact.Name, newContact.Email, newContact.Phone)

	if err != nil {
		log.Fatal("ERROR ACA")
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("Nuevo contacto registrado con éxito")
	fmt.Println("---------------------------------------------------------------")
	newContactId, _ := res.LastInsertId()
	GetContactById(db, int(newContactId))
}

// Update un nuevo contacto
func UpadateContact(db *sql.DB, newContact models.Contact) {
	//Query para insertar nuevo contacto
	query := "UPDATE contact SET name = ?, email = ?, phone=? WHERE id = ?"

	//Ejectuar la sentencia SQL
	_, err := db.Exec(query, newContact.Name, newContact.Email, newContact.Phone, newContact.Id)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("Registro actualizado correctamente")
	fmt.Println("---------------------------------------------------------------")
	GetContactById(db, newContact.Id)
}

// Borrar un registro de la bd
func DeleteContactById(db *sql.DB, contactID int) {
	// Query pára elimianr un registro
	query := "DELETE FROM contact WHERE id = ?"

	//Ejecutar la sentencia SQL
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("El registro con id %v fue eliminado con exito\n", contactID)
	fmt.Println("---------------------------------------------------------------")
}
