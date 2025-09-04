package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL")
	myURL := "https://jsonplaceholder.typicode.com/todos/resources?key1=value1"
	fmt.Printf("Type of URL: %T\n", myURL)
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println(" Can't parse URL", err)
	}
	fmt.Println("Type of URL| %T\n", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Scheme:", parsedURL.Host)
	fmt.Println("Scheme:", parsedURL.Path)
	fmt.Println("Scheme:", parsedURL.RawQuery)
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=iamprince"
	newUrl := parsedURL.String()
	fmt.Println("new URL", newUrl)
}
/*
URL
Type of URL: string
Type of URL| %T
 https://jsonplaceholder.typicode.com/todos/resources?key1=value1
Scheme: https
Scheme: jsonplaceholder.typicode.com
Scheme: /todos/resources
Scheme: key1=value1
new URL https://jsonplaceholder.typicode.com/newPath?username=iamprince
*/
