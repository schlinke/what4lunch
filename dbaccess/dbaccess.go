package dbaccess

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	// Load pq to import the driver "postgres"
	_ "github.com/lib/pq"
)

func getMenus(db *sql.DB) map[string]string {
	rows, err := db.Query("SELECT name,url FROM menu_url")
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

// ConnectAndGet return the Titles and URLs of all menus save in the DB
// For this the function establishes a connect with the cockroachdb
func ConnectAndGet() map[string]string {
	db, err := sql.Open("postgres", getDbConnectionString())
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return getMenus(db)

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
