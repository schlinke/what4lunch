package dbaccess

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	// Load pq to import the driver "postgres"
	_ "github.com/lib/pq"
)

// Menu is a structure containing the name of the restaurant, the url where the menu
// can be found and a searchstring to find to current menu
type Menu struct {
	Name, URL, Searchstring string
}

// GetMenu returns all Menus from the db
func GetMenu() []Menu {
	db := connect()

	return getMenuList(db, "SELECT name,url,searchstring FROM menu_url")
}

// GetLunch returns all Lunchmenus from the db
func GetLunch() []Menu {
	db := connect()

	return getMenuList(db, "SELECT name,url,searchstring FROM lunch_url")
}

func getMenuList(db *sql.DB, query string) []Menu {
	var menus []Menu

	rows, err := db.Query(query)
	if err != nil {

		log.Fatal(err)

	}

	if err != nil {

		log.Fatal(err)

	}
	defer rows.Close()

	for rows.Next() {
		var name, url, searchstring string
		err = rows.Scan(&name, &url, &searchstring)
		if err != nil {
			panic(err)
		}
		menus = append(menus, Menu{name, url, searchstring})
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
