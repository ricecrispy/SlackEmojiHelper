package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//Index - The landing page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Command - Spam: /spam/{emoji}/{num of emoji}")
	fmt.Fprintln(w, "Command - Insert claps emojis between each word: /insertclaps/{sentence}")
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
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	output := createClapsOutput(string(r.Form.Get("text")))
	writeJSON(w, output)
}

//SpamGetRequest - GET reqeust for Spam
func SpamGetRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	output := ""
	emoji := ps.ByName("emoji")
	num, err := strconv.Atoi(ps.ByName("num"))
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		for i := 0; i < num; i++ {
			output += emoji
		}

		writeJSON(w, output)
	}
}

//SpamPostRequest - POST request for Spam
func SpamPostRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	input := string(r.Form.Get("text"))
	writeJSON(w, input)
}

//Logic for writing the output in Json format for Slack
func writeJSON(w http.ResponseWriter, output string) {
	outputJSON := OutputJSON{"in_channel", output}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outputJSON)
}
