package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Name struct {
	Name    string
	Surname string
}

type Value struct {
	Type  string
	Value Joke
}

type Joke struct {
	Joke string
}

func main() {
	http.HandleFunc("/", responseFunc)
	http.ListenAndServe(":5000", nil)
}

func responseFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	resp, err := http.Get("http://uinames.com/api/")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	nameData, _ := ioutil.ReadAll(resp.Body)
	var name Name
	json.Unmarshal([]byte(nameData), &name)

	resp, err = http.Get("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	jokeData, _ := ioutil.ReadAll(resp.Body)
	var value Value
	json.Unmarshal([]byte(jokeData), &value)

	jokey := string(value.Value.Joke)
	r := strings.NewReplacer("John", name.Name, "Doe", name.Surname)
	result := r.Replace(jokey)
	w.Write([]byte(result))
}
