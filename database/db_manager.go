package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Db_Connection() (*sql.DB, error) {
	// cargar variables de entorno desde .env
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	// Abrir una conexion a la base de datos
	db, err := sql.Open("mysql", dns)

	if err != nil {
		return nil, err
	}

	return db, nil
}
