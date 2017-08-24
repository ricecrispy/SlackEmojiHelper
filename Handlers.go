package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//Index - The landing page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Skrt Skrt.")
}

//Logic to add the claps emoji to the user input
func createClapsOutput(input string) string {
	inputArr := strings.Split(input, " ")
	var word string

	for i := range inputArr {
		word = inputArr[i]
		inputArr[i] = strings.ToUpper(word) + " :clap:"
	}

	return strings.Join(inputArr, " ")
}

//InsertClapsGetRequest - GET request for InsertClaps
func InsertClapsGetRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	input := ps.ByName("input")
	writeJSON(w, createClapsOutput(input))
}

//InsertClapsPostRequest - POST request for InsertClaps
func InsertClapsPostRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	userName := r.Form.Get("user_name")
	userID := r.Form.Get("user_id")
	modifiedText := createClapsOutput(string(r.Form.Get("text")))
	finalText := fmt.Sprintf("%s (%s) said: %s", userName, userID, modifiedText)
	url := r.Form.Get("response_url")
	writeJSONToResponseURL(finalText, url)
}

//Logic for writing the output in Json format for Slack
func writeJSON(w http.ResponseWriter, output string) {
	outputJSON := OutputJSON{"in_channel", output}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outputJSON)
}

//Logic for writing the output in Json format for Slack via the request's response_url
func writeJSONToResponseURL(output string, url string) {
	outputJSON := OutputJSON{"in_channel", output}
	jBuffer := new(bytes.Buffer)
	json.NewEncoder(jBuffer).Encode(outputJSON)
	_, err := http.Post(url, "application/json", jBuffer)
	if err != nil {
		panic(err)
	}
}
