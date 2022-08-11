package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

func MountOpenApi(router *mux.Router, baseUrl string) {
	fatmanName := os.Getenv("FATMAN_NAME")
	fatmanVersion := os.Getenv("FATMAN_VERSION")

	router.HandleFunc(baseUrl+"/static/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		tmplt := template.New("openapi.json")
		tmplt, _ = tmplt.ParseFiles("./swaggerui/openapi.json")

		context := map[string]string{
			"fatmanName":    fatmanName,
			"fatmanVersion": fatmanVersion,
		}
		tmplt.Execute(w, context)
	})
	router.HandleFunc(baseUrl+"/static/{filename}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		filename := vars["filename"]
		http.ServeFile(w, r, "./swaggerui/"+filename)
	})

	router.HandleFunc(baseUrl+"", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.URL.Path+"/", http.StatusTemporaryRedirect)
	})
	router.HandleFunc(baseUrl+"/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swaggerui/index.html")
	})
}
