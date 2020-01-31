package dbaccess

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	// Load pq to import the driver "postgres"
	_ "github.com/lib/pq"
)

// GetMenu return all names and urls of all menus
func GetMenu() map[string]string {
	db := connect()

	return getMenuList(db, "SELECT name,url FROM menu_url")
}

// GetLunch return all names and urls of all menus
func GetLunch() map[string]string {
	db := connect()

	return getMenuList(db, "SELECT name,url FROM lunch_url")
}

func getMenuList(db *sql.DB, query string) map[string]string {
	rows, err := db.Query(query)
	if err != nil {

		log.Fatal(err)

	}

	if err != nil {

		log.Fatal(err)

	}
	defer rows.Close()

	menus := make(map[string]string)

	for rows.Next() {
		var name, url string
		err = rows.Scan(&name, &url)
		if err != nil {
			panic(err)
		}
		menus[name] = url
	}

	return menus

}

// Initialize the connect to a database which is descibed in dbinfo.txt
func connect() *sql.DB {
	db, err := sql.Open("postgres", getDbConnectionString())
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db

}

func getDbConnectionString() string {
	file, err := os.Open("dbinfo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	return string(b)
}
