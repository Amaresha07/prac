package main

import "fmt"

func checktype(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	case float64:
		fmt.Println("float", v)
	case bool:
		fmt.Println("boolean", v)
	default:
		fmt.Println("unknow type")
	}
}

func main() {
	checktype(10)
	checktype("Amaresha")
	checktype(8.55)
	checktype(true)

}
