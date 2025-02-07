package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tag      []string `json:"Tag,omitempty"` /* omitempty means if Tag is nil then
	   this field will be removed from JSON
	*/
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React.js", 299, "Online.com", "Password@345", []string{"metatag", "core", "limited"}},
		{"React redux", 399, "Online.com", "Password@345", []string{"metatag", "core", "limited"}},
		{"React next", 499, "Online.com", "Password@345", []string{"metatag", "core", "limited"}},
		{"React mern", 0, "Online.com", "Password@345", nil},
	}
	finalJson, _ := json.MarshalIndent(courses, "", "\t")
	fmt.Printf(string(finalJson))
}
