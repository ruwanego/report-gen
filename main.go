package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ruwanego/report-gen/util"
)

func main() {
	const FileName = "sp-survey-202003-data.xlsx"
	const SheetName = "Sheet1"

	resultSheet, err := excelize.OpenFile(FileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := resultSheet.GetRows(SheetName)

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

	fmt.Print(modeChoices)

	whole := make([]map[string]string, len(rows)-1)

	for j := 1; j < len(rows); j++ {
		elementMap := make(map[string]string)
		for i := 0; i < len(rows[j]); i++ {
			if i >= START && i <= END {
				continue
			}
			elementMap[columnHeaders[i]] = strings.TrimSpace(rows[j][i])
		}
		elementMap["mode_choice_1"] = modeChoices[j][0]
		elementMap["mode_choice_2"] = modeChoices[j][1]
		elementMap["mode_choice_3"] = modeChoices[j][2]
		elementMap["mode_choice_4"] = modeChoices[j][3]
		elementMap["mode_choice_5"] = modeChoices[j][4]
		elementMap["meta-audit"] = util.ConvertToDirectLink(elementMap["meta-audit"])
		whole[j-1] = elementMap
	}

	var results []util.Result
	var jsonString, _ = json.Marshal(whole)
	err = json.Unmarshal(jsonString, &results)
	if err != nil {
		panic(err)
	}

	jsonString, _ = json.Marshal(results)

	fixedJSONString := string(bytes.Replace([]byte(jsonString), []byte("\\u0026"), []byte("&"), -1))

	f2, err2 := os.Create("results.json")
	if err2 != nil {
		fmt.Println(err)
		return
	}

	_, err3 := f2.WriteString(fixedJSONString)
	if err3 != nil {
		fmt.Println(err)
		f2.Close()
		return
	}

	err4 := f2.Close()
	if err4 != nil {
		fmt.Println(err)
		return
	}
}
