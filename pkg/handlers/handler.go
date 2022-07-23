package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/mpvl/unique"
)

//this is the struct which we use to take input from bosy in json
type Rules struct {
	Rules []struct {
		ID       string   `json:"id"`
		Head     string   `json:"head"`
		Body     string   `json:"body,omitempty"`
		Requires []string `json:"requires,omitempty"`
	} `json:"rules"`
}

//this is the struct use to output in json
type newStruct struct {
	Rules []struct {
		ID       string   `json:"id"`
		Head     string   `json:"head"`
		Body     string   `json:"body,omitempty"`
		Requires []string `json:"requires,omitempty"`
	} `json:"rules"`
}
type result struct {
	Rules []struct {
		ID   string `json:"id"`
		Head string `json:"head"`
		Body string `json:"body,omitempty"`
	} `json:"rules"`
}

//this is our important function, here we start the process
//r --> request, w --> response
func HandleHttpPolicy(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("filepath")
	fmt.Println("filePath =>", filePath)
	var rule Rules
	response := newStruct{}

	//read r.Body --> JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &rule) //Unmarshall to convert to struct
	if err != nil {
		panic(err)
	}

	//list of rules which appended
	append_rule := []string{}

	var a []int
	//loop through the rules from our input json
	for i := range rule.Rules {
		bytesP, _ := json.Marshal(rule.Rules[i]) //for checking in debugging
		bytesS := string(bytesP)
		fmt.Println(bytesS)

		// if len(rule.Rules[i].Requires) == 0 || strings.Compare(append_rule[0], rule.Rules[i].Requires[0]) == 0 {
		if len(rule.Rules[i].Requires) == 0 { //if no requires, we add to the first
			response.Rules = append(response.Rules, rule.Rules[i])
			append_rule = append(append_rule, rule.Rules[i].ID)
		}

		for x := range rule.Rules[i].Requires {
			r := rule.Rules[i].Requires[x]
			fmt.Println("r: ", r)
			if strings.Compare(append_rule[0], r) == 0 {
				append_rule = append(append_rule, rule.Rules[i].ID)
				response.Rules = append(response.Rules, rule.Rules[i])
			} else if len(append_rule) > 1 {
				if strings.Compare(append_rule[1], r) == 0 {
					append_rule = append(append_rule, rule.Rules[i].ID)
					response.Rules = append(response.Rules, rule.Rules[i])
				}
			} else {
				a = append(a, i)
			}
		}
	}
	//appending the rules missed by first loop
	unique.Ints(&a)
	fmt.Println(a)
	for i := range rule.Rules {
		for x := range rule.Rules[i].Requires {
			r := rule.Rules[i].Requires[x]
			if strings.Compare(append_rule[i], r) == 0 {
				append_rule = append(append_rule, rule.Rules[i].ID)
				response.Rules = append(response.Rules, rule.Rules[i])
			}
		}
	}
	//create a new struct
	resultStruct := result{}
	//creating a map, because we want to delete requires in the output
	//in go, we cant delete keys in struct, we can though delete in map
	maap := structs.Map(response)
	//deleting the requires key with the values
	delete(maap, "requires")
	//after deleting the key/value again decoding the map to struct
	mapstructure.Decode(maap, &resultStruct)
	//saving the response json with the responseFile.txt name
	file, _ := json.MarshalIndent(resultStruct, "", " ")
	path := filePath
	fileName := fmt.Sprintf("%s/responseFile.txt", path)
	// fileName := "cmd/web/responseFile.txt"
	_ = ioutil.WriteFile(fileName, file, 0644) //644?
	//json response send as a output
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultStruct)
}
