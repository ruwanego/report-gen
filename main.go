package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// convertToDirectLink Converts the google drive view link of the audit.csv file
// to a direct link
func convertToDirectLink(auditURL string) string {
	// eg. https://drive.google.com/open?id=1q4ubKjRBCPS1eViYyiLcivp4cA7iG41d
	// 1q4ubKjRBCPS1eViYyiLcivp4cA7iG41d
	u, err := url.Parse(auditURL)
	if err != nil {
		log.Fatal(err)
	}

	id := u.Query()["id"]

	tmpl, err := template.New("url").Parse("https://drive.google.com/uc?export=download&id={{.Id}}")
	if err != nil {
		panic(err)
	}

	var tmplBytes bytes.Buffer

	err = tmpl.Execute(&tmplBytes, id)
	if err != nil {
		panic(err)
	}

	return tmplBytes.String()
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

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

func main() {
	resultSheet, err := excelize.OpenFile("sp-survey-202003-data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := resultSheet.GetRows("Sheet1")

	columnHeaders := rows[0]
	// fix column header names by removing "data-"
	for i := 0; i < len(columnHeaders); i++ {
		columnHeaders[i] = strings.Replace(columnHeaders[i], "data-", "", -1)
	}

	// compact mode choices
	modeChoices := make(map[int][]string)
	const START = 36
	const END = 2510
	for i := 1; i < len(rows); i++ {
		var s []string
		for j := START; j < END; j++ {
			if strings.TrimSpace(rows[i][j]) != "" {
				s = append(s, rows[i][j])
			}
		}
		modeChoices[i] = s
		s = nil
	}
	prettyPrint(modeChoices)

	whole := make([]map[string]string, len(rows)-1)

	for j := 1; j < len(rows); j++ {
		elementMap := make(map[string]string)
		for i := 0; i < len(rows[j]); i++ {
			if i >= 36 && i <= 2509 {
				continue
			}
			elementMap[columnHeaders[i]] = strings.TrimSpace(rows[j][i])
		}
		elementMap["mode_choice_1"] = modeChoices[j][0]
		elementMap["mode_choice_2"] = modeChoices[j][1]
		elementMap["mode_choice_3"] = modeChoices[j][2]
		elementMap["mode_choice_4"] = modeChoices[j][3]
		elementMap["mode_choice_5"] = modeChoices[j][4]
		whole[j-1] = elementMap
	}

	// if i >= 20 && i <= 2494 {
	// 	if strings.TrimSpace(rows[j][i]) != "" {
	// 		prettyPrint(rows[j][i])
	// 		elementMap[rows[0][i]] = rows[j][i]
	// 		continue
	// 	}
	// }
	jsonString, err := json.Marshal(whole)

	f2, err2 := os.Create("results.json")
	if err2 != nil {
		fmt.Println(err)
		return
	}
	_, err3 := f2.WriteString(string(jsonString))
	if err3 != nil {
		fmt.Println(err)
		f2.Close()
		return
	}
	// fmt.Println(l, "bytes written successfully")
	err4 := f2.Close()
	if err4 != nil {
		fmt.Println(err)
		return
	}
}
