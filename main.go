package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// I'm just hardcoding the url from the service info record...
const url = "https://api-8101-1.dmsi.dev:443/develtestdbPODService"

// TODO add username and password here
const username = ""
const password = ""

func main() {
	// since we're using cookies, I created a class that encapsulates
	// the url and the state of the cookies; once logged in, the client
	// carries the JSESSIONID cookie
	client := NewClient(url)
	login, err := client.Login(username, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer login.Body.Close()
	fmt.Println(login.StatusCode) // 302

	// client has logged-in state
	defaults, err := client.GetUserDefaults()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer login.Body.Close()
	fmt.Println(defaults.StatusCode) // 200

	fmt.Println(client.client.Jar) // dump the cookie jar to demonstrate what's happening here

	b, err := ioutil.ReadAll(defaults.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(string(b)) // hopefully some json
}
