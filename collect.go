package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func CollectData() []Data {

	exelName := os.Getenv("EXEL_NAME")
	if exelName == "" {
		log.Fatalln("oyy.... file exel e apa namanya ?")
	}

	sheetName := os.Getenv("SHEET_NAME")
	if exelName == "" {
		log.Println("nama sheet nya oy jangan lupa")
	}

	var start int
	var err error
	startIndex := os.Getenv("START_INDEX")
	if startIndex == "" {
		start = 2
	} else {
		start, err = strconv.Atoi(startIndex)
		if err != nil {
			log.Fatal("start_index kudu nomer oy, dan bilangan bulat")
		}
	}

	var end int
	endIndex := os.Getenv("END_INDEX")
	if endIndex == "" {
		log.Fatal("end_index wajib diisi")
	} else {
		end, err = strconv.Atoi(endIndex)
		if err != nil {
			log.Fatal("end_index kudu nomer oy, dan bilangan bulat")
		}
	}

	xlsx, err := excelize.OpenFile(exelName)
	if err != nil {
		log.Fatal(err)
	}

	var rows []Data
	for i := start; i <= end; i++ {
		row := Data{
			Email: xlsx.GetCellValue(sheetName, fmt.Sprintf("B%d", i)),
			Name:  xlsx.GetCellValue(sheetName, fmt.Sprintf("C%d", i)),
		}

		rows = append(rows, row)
	}

	return rows
}
