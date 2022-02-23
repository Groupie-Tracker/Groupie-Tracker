package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type API struct {
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
	Id   int
}

type Artists1 struct {
	Artists []API
}

var templates = template.Must(template.ParseFiles("HTML/hpage.html"))
var templates2 = template.Must(template.ParseFiles("HTML/artist.html"))
var templates3 = template.Must(template.ParseFiles("HTML/truc.html"))
var VarApi TemplateData
var VarArtists Artists1
var ApiObject []API
var data string
var Id_data string
var test string

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

	json.Unmarshal(ApiData, &ApiObject)

	VarArtists = Artists1{
		Artists: ApiObject,
	}
	templates.Execute(w, VarArtists)
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

	test := r.FormValue("Oui")
	fmt.Println("r.formvalue :", test)

	// // test = test.Base(r.URL.test)
	// test := r.URL.Path
	// fmt.Println("le test   ", path.Base(r.URL.Path), "test ", test)

	// if test != "favicon.ico" {
	// 	Id_data = test
	// }
	// fmt.Println("iddata avant", Id_data)
	templates2.Execute(w, err)
	// fmt.Println(Id_data)
}

func details(w http.ResponseWriter, r *http.Request) {

	temporaire, _ := strconv.Atoi(Id_data)
	fmt.Println("azeaze ", temporaire)
	VarApi = TemplateData{
		Name: ApiObject[temporaire].Name,
	}

	templates3.Execute(w, VarApi)
}

func main() {
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/artist", artist)
	http.HandleFunc("/artist/", details)

	log.Fatal(http.ListenAndServe(":55", nil))
}
