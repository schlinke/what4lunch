package urlparser

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	dc "github.com/schlinke/what4lunch/datecalculator"
	"github.com/schlinke/what4lunch/dbaccess"
)

func findMenuPdf(url string, searchstring string) string {
	fileurl := ""
	r, err := regexp.Compile(searchstring)
	if err != nil {
		log.Fatal(err)
	}

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	fileurl = r.FindString(bodyString)

	return fileurl
}

// GetAll downloads all menus save in the DB as Pdf-files
func GetAll() {
	GetLunchFromWww()
	GetMenusFromWww()
}

// GetMenusFromWww downloads all current menus with are saved in the db
func GetMenusFromWww() {
	menu := dbaccess.GetMenu()
	path := "menu"

	downloadPdf(menu, path)
}

// GetLunchFromWww downloads all current lunch menus with are saved in the db
func GetLunchFromWww() {
	lunch := dbaccess.GetLunch()
	path := "lunch"

	downloadPdf(lunch, path)
}

// TODO find a better name for this function
// this function iterate over the list of urls and set the correct path
// after this the function starts the download
func downloadPdf(menu []dbaccess.Menu, folder string) {
	for _, element := range menu {
		path := folder + "/" + strconv.Itoa(dc.GetCW(time.Now()))
		file := element.Name + ".pdf"
		error := downloadMenu(path, file, findMenuPdf(element.URL, element.Searchstring))

		if error != nil {
			fmt.Println(error)
		}
	}
}

func checkDir(path string) (err error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0770)
	}

	if err != nil {
		return err
	}

	return nil
}

func downloadMenu(path string, file string, url string) (err error) {
	err = checkDir(path)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(path + "/" + file)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// GetLinksLunch does something
func GetLinksLunch() map[string]string {
	_, cw := time.Now().ISOWeek()
	links := make(map[string]string)

	lunchlist := dbaccess.GetLunchNames()
	for _, element := range lunchlist {
		filepath := "/lunch/" + strconv.Itoa(cw) + "/" + element + ".pdf"
		filepath = strings.ReplaceAll(filepath, " ", "_")
		links[element] = filepath
	}

	return links
}

// GetLinksMenu does something
func GetLinksMenu() map[string]string {
	_, cw := time.Now().ISOWeek()
	links := make(map[string]string)

	lunchlist := dbaccess.GetMenuNames()
	for _, element := range lunchlist {
		filepath := "/menu/" + strconv.Itoa(cw) + "/" + element + ".pdf"
		filepath = strings.ReplaceAll(filepath, " ", "_")
		links[element] = filepath
	}

	return links
}
