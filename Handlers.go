package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Command - Spam: /spam/{emoji}/{num of emoji}")
	fmt.Fprintln(w, "Command - Insert claps emojis between each word: /insertclaps/{sentence}")
	fmt.Fprintln(w, "Skrt Skrt.")
}

func InsertClaps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {	
	input := ps.ByName("input")
	
	output := CreateClapsOutput(input)

	WriteJson(w, output)
}



func InsertClapsPostRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	output := CreateClapsOutput(string(r.Form.Get("text")))
	WriteJson(w, output)
}

func CreateClapsOutput(input string) string {
	inputArr := strings.Split(input, " ")

	for i := 0; i < len(inputArr); i++ {
		word := inputArr[i]
		inputArr[i] = strings.ToUpper(word) + " :clap:"
	}

	return strings.Join(inputArr, " ")
}



func Spam(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	output := ""
	emoji := ps.ByName("emoji")
	num, err := strconv.Atoi(ps.ByName("num"))
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		for i := 0; i < num; i++ {
			output += emoji
		}

		WriteJson(w, output)
	}
}

func WriteJson(w http.ResponseWriter, output string) {
	outputJson := OutputJson{"in_channel", output}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outputJson)
}
