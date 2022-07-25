package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filePath := os.Args[1]

	fmt.Println("file path: " + filePath)

	body, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal("File not found")
	}

	stringBody := string(body)

	index := strings.Index(stringBody, "#:")

	if index < 0 {
		log.Fatal("Cannot extract translation data from file")
	}

	stringBody = stringBody[index:]
	stringBody = strings.ReplaceAll(stringBody, "\"", "")
	stringBody = strings.ReplaceAll(stringBody, "\n", "")

	splitString := strings.Split(stringBody, "#:")

	var csvData = [][]string{{"string", "translation"}}

	for _, value := range splitString {
		msgIdPos := strings.Index(value, "msgid")
		msgStrPos := strings.Index(value, "msgstr")

		if msgIdPos < 0 || msgStrPos < 0 {
			continue
		}

		msgId := value[msgIdPos+6 : msgStrPos]
		msgStr := value[msgStrPos+7:]

		row := []string{msgId, msgStr}

		csvData = append(csvData, row)
	}

	currentTime := time.Now()
	currentTimeString := currentTime.Format("20060102150405")

	filename := fmt.Sprintf("PoExport%s.csv", currentTimeString)

	csvFile, err := os.Create(filename)

	if err != nil {
		log.Fatal("Cannot create export csv file")
	}

	w := csv.NewWriter(csvFile)

	err = w.WriteAll(csvData)

	if err != nil {
		log.Fatal("Cannot write content into CSV file")
	}
}
