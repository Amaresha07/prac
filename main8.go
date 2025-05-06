package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := user{name: "amaresha", age: 21}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData))
}
