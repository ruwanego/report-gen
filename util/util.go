package util

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// ConvertToDirectLink Converts the google drive view link of the audit.csv file
// to a direct link
func ConvertToDirectLink(auditURL string) string {
	// eg. we have https://drive.google.com/open?id=1q4ubKjRBCPS1eViYyiLcivp4cA7iG41d
	// we want to convert it to : https://drive.google.com/uc?export=download&id=1q4ubKjRBCPS1eViYyiLcivp4cA7iG41d
	// SOMETHING
	u, err := url.Parse(auditURL)
	if err != nil {
		log.Fatal(err)
	}
	// SOMETHING
	id := u.Query()["id"][0]
	log.Print(id)
	s := "https://drive.google.com/uc?export=download" + "&" + "id=" + id
	return s
}
