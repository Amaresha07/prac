// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Welcome to the Home Page!")
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "This is the About Page.")
// }

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	http.HandleFunc("/about", aboutHandler)
// 		fmt.Println("Server running at http://localhost:8080/")
// 	http.ListenAndServe(":8080", nil)
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	var ptr = 10
// 	num := &ptr
// 	fmt.Println(&ptr)
// 	fmt.Println(&num)
// 	size := unsafe.Sizeof(num)
// 	fmt.Println(size)

// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// // Define a User struct
// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// // Create a handler function for the /user POST route
// func createUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Check if the method is POST
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Parse the incoming JSON body into the User struct
// 	var user User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	// Send a JSON response back with the user data
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message": "User created",
// 		"user":    user,
// 	})
// }

// func main() {
// 	// Handle the /user POST route using the createUserHandler
// 	http.HandleFunc("/user", createUserHandler)

// 	// Start the server on port 8080
// 	fmt.Println("Server is running on port 8080...")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	name string `json:"name"`
	age  int    `jsom:"age"`
}

func creathandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "invalid response", http.StatusMethodNotAllowed)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	w.Header().Set("constant type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
		"user":    user,
	})
}

func main() {
	http.HandleFunc("/user", creathandler)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
