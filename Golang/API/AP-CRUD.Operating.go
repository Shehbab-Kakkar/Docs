package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error Getting:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response", res.Status)
		return
	}
	// 	data, err := io.ReadAll(res.Body)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, "Error reading:", err)
	// 		return
	// 	}
	// 	fmt.Println("Data:", string(data))
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding :", err)
		return
	}
	fmt.Println("Todo: ", todo)

	fmt.Println("Title response", todo.Title)

	fmt.Println("Completed response", todo.Completed)

}

func performPostRequest() {
	// 1. Create a Todo struct instance with some data
	todo := Todo{
		UserID:    23,             // Assign UserID field
		Title:     "Prince Kumar", // Assign Title field
		Completed: true,           // Assign Completed field
		// Id is omitted, so it will default to 0
	}

	// 2. Convert the Todo struct to JSON format for transmission
	jsonData, err := json.Marshal(todo) // Perform marshaling (Go struct -> JSON []byte)
	if err != nil {
		fmt.Println("Error marshalling :", err)
		return // Exit if marshaling fails
	}

	// 3. Convert the JSON bytes to a string (not strictly necessary, but done here)
	jsonString := string(jsonData)

	// 4. Create an io.Reader from the JSON string (required by http.Post)
	jsonReader := strings.NewReader(jsonString)

	// 5. Define the URL to which the POST request will be sent
	myURL := "https://jsonplaceholder.typicode.com/todos"

	// 6. Send HTTP POST request with JSON body
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return // Exit if the POST fails
	}
	defer res.Body.Close() // Ensure response body is closed after function returns

	// 7. Read the response body from the server
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading:", err)
		return // Exit if reading the response fails
	}

	// 8. Print the response data and HTTP status
	fmt.Println("Data:", string(data))          // Print the response JSON (usually echoes what you sent, plus assigned Id)
	fmt.Println("Response Status:", res.Status) // Print HTTP status (e.g., "201 Created")
}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"
	//create DELETE Request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating Delete Request", err)
		return
	}
	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Sending request", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response status of Deletation: ", res.Status)
}
func main() {
	fmt.Println("Learning CRUD")
	performGetRequest()
	performPostRequest()
	performDeleteRequest()
}
