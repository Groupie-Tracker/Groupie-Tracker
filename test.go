package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
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

type Artists1 struct {
	Artists []API
	PathID  string
}

type DescritpionPage struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

var templates = template.Must(template.ParseFiles("HTML/hpage.html"))
var templates2 = template.Must(template.ParseFiles("HTML/artist.html"))
var templates3 = template.Must(template.ParseFiles("HTML/truc.html"))
var ApiObject []API
var data string
var Id_data string
var test DescritpionPage

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

	VarArtists := Artists1{
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

	templates2.Execute(w, err)
}

func details(w http.ResponseWriter, r *http.Request) {
	pathID := r.URL.Path
	pathID = path.Base(pathID)
	pathIDint, _ := strconv.Atoi(pathID)
	VarArtists := DescritpionPage{
		ID:           ApiObject[pathIDint-1].ID,
		Image:        ApiObject[pathIDint-1].Image,
		Members:      ApiObject[pathIDint-1].Members,
		CreationDate: ApiObject[pathIDint-1].CreationDate,
		FirstAlbum:   ApiObject[pathIDint-1].FirstAlbum,
		Locations:    ApiObject[pathIDint-1].Locations,
		ConcertDates: ApiObject[pathIDint-1].ConcertDates,
		Relations:    ApiObject[pathIDint-1].Relations,
	}

	templates3.Execute(w, VarArtists)
}

func main() {
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", home)
	http.HandleFunc("/artist", artist)
	http.HandleFunc("/artist/", details)

	log.Fatal(http.ListenAndServe(":55", nil))
}
