package main

import "fmt"

// Function returning an interface{}
func getData() interface{} {
	return 3.14
}

func senddata() interface{} {
	return "Amaresha"
}

func main() {
	data := getData()
	newdata := senddata()

	value, ok := data.(float64)
	if ok {
		fmt.Println("Extracted value:", value) // ✅ Output: Extracted value: 3.14
	} else {
		fmt.Println("Type assertion failed")
	}

	newvalue, ok := newdata.(string)
	if ok {
		fmt.Println("Extracted  value:", newvalue)
	} else {
		fmt.Println("type assettion failed")
	}
}
