package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type API []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type TemplateData struct {
	Test string
	Img  string
}

var templates = template.Must(template.ParseFiles("HTML/hpage.html"))

func home(w http.ResponseWriter, r *http.Request) {

	Api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	ApiData, err := ioutil.ReadAll(Api.Body)
	if err != nil {
		log.Fatal(err)
	}
	var ApiObject API
	json.Unmarshal(ApiData, &ApiObject)
	for i := 0; i < len(ApiObject); i++ {
		fmt.Println(ApiObject[i].Name)
		fmt.Println(ApiObject[i].Image)
	}
	for i := 0; i < len(ApiObject); i++ {
		VarApi := Temp{
			Test: ApiObject[i].Name,
			Img:  ApiObject[i].Image,
		}
		templates.Execute(w, VarApi)
	}
}

func main() {
	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":55", nil))
}
