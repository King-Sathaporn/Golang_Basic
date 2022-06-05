package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Request to the server and get the response.
func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //ioutil.ReadAll is a function that reads all the data from the response body.
	fmt.Println(string(body))              // Print html code
}
