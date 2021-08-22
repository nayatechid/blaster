package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"time"
)

func BlastToEmail(subject string, receivers [][]string, t *template.Template, delay time.Duration) {
	var data = receivers
	var dataError []Data
	var success, sleep int
	var dataE int

	mappedData := make([]Data, 0)
	for _, d := range data {
		mappedData = append(mappedData, Data{
			Name:  d[0],
			Email: d[1],
		})
	}

	maxPerBatch := 5
	fmt.Println("starting")
	for _, val := range mappedData {
		if sleep == maxPerBatch {
			time.Sleep(delay)
			sleep = 0
		}
		sleep++
		fmt.Printf("sending reminder to --> %s \n", val.Email)
		if err := SendMail(subject, t, val); err != nil {
			fmt.Printf("error sending interview to --> %s, with error %v\n", val.Email, err.Error())
			dataError = append(dataError, val)
			dataE++
			continue
		}
		success++
		fmt.Println("email sended")
	}
	fmt.Println("finished")

	if len(dataError) == 0 {
		fmt.Printf("data success --> %d\n", success)
		fmt.Println("no data error")
	} else {
		jsonB, err := json.Marshal(dataError)
		if err != nil {
			fmt.Println(dataError)
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("data success --> %d\n", success)

		fmt.Printf("data error --> %d\n", dataE)
		fmt.Println(string(jsonB))
	}

}
