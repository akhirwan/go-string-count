package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// input
	input := "aaaaaaasssssiiiAA"

	// optional if don't use space
	input = strings.ReplaceAll(input, " ", "")
	// optional if lower case OR strings.ToUpper(input) if upper case
	// input = strings.ToLower(input)

	// make unique array input
	keys := make(map[string]bool)
	unique := []string{}

	for _, val := range input {
		if _, v := keys[string(val)]; !v {
			keys[string(val)] = true
			unique = append(unique, string(val))
		}
	}

	// count input array value
	mapResult := make(map[string]int)

	for i := 0; i < len(unique); i++ {
		mapResult[string(unique[i])] = strings.Count(input, string(unique[i]))
	}

	// marshal json result
	result, err := json.Marshal(mapResult)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(result))
	}
}
