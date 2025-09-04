package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting Get response", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n", res)
	//fmt.Println("reponse", res)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error ")
		return
	}

	fmt.Println("response", string(data))
}
/*
Web service
Type of response: *http.Response
response {
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
*/
