package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jcav67/go-mysql/database"
	"github.com/jcav67/go-mysql/handlers"
	"github.com/jcav67/go-mysql/models"
)

func main() {

	db, err := database.Db_Connection()
	if err != nil {
		log.Fatal("ERROR IN DATABASE CONNECTION")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
		log.Fatal("ERROR IN DATABASE CONNECTION")
	}
	fmt.Println("Conexion a la base de datos exitosa")

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener un contacto por Id")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar un contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Println("Seleccione una opción:")

		// Leer la opcion seleccionada
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListarContacts(db)
		case 2:
			fmt.Println("Ingrese el Id del contacto a buscar: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactById(db, idContact)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
		case 4:
			updatedContact := inputContactDetails(option)
			handlers.UpadateContact(db, updatedContact)
		case 5:
			fmt.Println("Ingrese el Id del contacto a eliminar: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContactById(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Por favor, intente nuevamente")
		}
	}
}

// Funcion para ingresar los detalles del contacto desde la entrada
func inputContactDetails(option int) models.Contact {
	//  Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Println("Ingrese el Id del contacto a editar: ")
		var idContact int
		fmt.Scanln(&idContact)

		contact.Id = idContact
	}

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el email del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
