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
	Name string
	Img  string
	Glbl string
	Id   int
}

var templates = template.Must(template.ParseFiles("HTML/hpage.html"))
var templates2 = template.Must(template.ParseFiles("HTML/artist.html"))
var VarApi TemplateData

func artist(w http.ResponseWriter, r *http.Request) {

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
		VarApi = TemplateData{
			Name: ApiObject[i].Name,
			Img:  ApiObject[i].Image,
			Id:   ApiObject[i].ID,
		}
		templates.Execute(w, VarApi)
	}

}

func home(w http.ResponseWriter, r *http.Request) {

	Api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	ApiDataArtist, err := ioutil.ReadAll(Api.Body)
	if err != nil {
		log.Fatal(err)
	}
	var ApiObject API
	json.Unmarshal(ApiDataArtist, &ApiObject)

	templates2.Execute(w, err)
}

func main() {
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/artist", artist)

	log.Fatal(http.ListenAndServe(":55", nil))
}
