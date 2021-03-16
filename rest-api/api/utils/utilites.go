package utils

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

func BodyParser(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	return body
}

func ToJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	checkErr(err)
}

func ToXML(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json; charset=UTF8")
	w.WriteHeader(statusCode)
	err := xml.NewEncoder(w).Encode(data)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
