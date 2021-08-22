package main

import (
	"embed"
	"html/template"
	"os"
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
	command, err := ParseArguments(os.Args[1:])
	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	t, err := template.ParseFS(Template, command["template"])
	if err != nil {
		panic(err)
	}

	receivers, err := CollectReceiversFromCSV(command["receivers"], 1, 0)
	if err != nil {
		panic(err)
	}

	BlastToEmail(command["subject"], receivers, t, time.Minute)
}
