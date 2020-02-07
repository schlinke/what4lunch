package urlparser

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	dc "github.com/schlinke/what4lunch/datecalculator"
	"github.com/schlinke/what4lunch/dbaccess"
)

func findMenuPdf(url string, searchstring string) string {
	fileurl := ""

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, element *goquery.Selection) {
		href, _ := element.Attr("href")
		if matched, _ := regexp.MatchString(searchstring, href); matched {
			fileurl = href
			return
		}
	})

	return fileurl
}

// ParseURL replaces the placesholder in the url with its correct values
// i.e. {year} to 2020
func ParseURL(url string, t time.Time) string {

	newurl := strings.Replace(url, "{year}", strconv.Itoa(dc.GetYear(t)), -1)
	newurl = strings.Replace(newurl, "{intmonth}", strconv.Itoa(dc.GetIntMonth(t)), -1)
	if strings.Contains(newurl, "{mondaydate}") {
		a, b := dc.GetDayAndMonth(dc.GetDayOfCW(t, 1))
		newurl = strings.Replace(newurl, "{mondaydate}", fmt.Sprintf("%d.%d.", a, b), -1)
	}
	if strings.Contains(newurl, "{fridaydate}") {
		a, b := dc.GetDayAndMonth(dc.GetDayOfCW(t, 5))
		newurl = strings.Replace(newurl, "{fridaydate}", fmt.Sprintf("%d.%d.", a, b), -1)
	}

	return newurl
}

// GetMenusFromWww downloads all current menus with are saved in the db
func GetMenusFromWww() {
	lunch := dbaccess.GetLunch()

	for _, element := range lunch {
		path := "lunch/" + strconv.Itoa(dc.GetCW(time.Now()))
		file := element.Name + ".pdf"
		downloadMenu(path, file, findMenuPdf(element.URL, element.Searchstring))
	}

	// for k, v := range lunch {
	// 	path := "lunch/" + strconv.Itoa(dc.GetCW(time.Now()))
	// 	file := k + ".pdf"
	// 	downloadMenu(path, file, ParseURL(v, time.Now()))

	// }

	// menu := dbaccess.GetMenu()
	// for k, v := range menu {
	// 	path := "menu/" + strconv.Itoa(dc.GetCW(time.Now()))
	// 	file := k + ".pdf"
	// 	downloadMenu(path, file, ParseURL(v, time.Now()))

	// }
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
