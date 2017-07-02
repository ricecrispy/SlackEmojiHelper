package main

import (
	"fmt"
	"log"
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
	inputArr := strings.Split(input, " ")

	fmt.Fprintln(w, len(inputArr))

	for i := 0; i < len(inputArr); i++ {
		word := inputArr[i]
		inputArr[i] = word + " :clap:"
	}

	output := strings.Join(inputArr, " ")
	fmt.Fprintln(w, output)
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
		fmt.Fprintln(w, output)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/insertclaps/:input", InsertClaps)
	router.GET("/spam/:emoji/:num", Spam)

	log.Fatal(http.ListenAndServe("https://slackemojihelper.herokuapp.com", router))
}
