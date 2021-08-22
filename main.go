package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"time"

	"github.com/joho/godotenv"
)

//go:embed template.html
var Template embed.FS

type Data struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	t, err := template.ParseFS(Template, "template.html")
	if err != nil {
		panic(err)
	}

	var data = CollectData()
	var dataError []Data
	var success, sleep int
	var dataE int

	fmt.Println("starting")
	for _, val := range data {
		if sleep == 5 {
			time.Sleep(2 * time.Minute)
			sleep = 0
		}
		sleep++
		fmt.Printf("sending reminder to --> %s ", val.Email)
		if err := SendReminder(t, val); err != nil {
			fmt.Printf("error sending interview to --> %s, with error %v", val.Email, err.Error())
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
