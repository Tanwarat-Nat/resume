package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	out, err := os.Create("index.html")
	check(err)
	defer out.Close()

	data := map[string]interface{}{}

	file, err := ioutil.ReadFile("data.json")
	check(err)

	err = json.Unmarshal(file, &data)
	check(err)

	t, err := template.ParseGlob("templete/*")
	check(err)

	err = t.Execute(out, data)
	check(err)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
