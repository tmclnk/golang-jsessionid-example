package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

// Client is a stateful client
type Client struct {
	url    string
	client http.Client
}

// Creates a client with a stateful url and cookie jar which will ignore redirects
func NewClient(url string) *Client {
	// Let the cookie jar manage the cookies for us
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}
	return &Client{
		url: url,
		client: http.Client{
			Jar: jar,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				// redirects over to :80 might fail if we follow them,
				// so we are using this to disable that
				return http.ErrUseLastResponse
			},
		},
	}
}

func (c *Client) Login(username string, password string) (*http.Response, error) {
	url := fmt.Sprintf("%s/static/auth/j_spring_security_check", c.url)
	req, err := http.NewRequest("POST", url, nil)

	// this part is admittedly ludicrous; adding the query params to the Query object
	// isn't enough; you have to re-encode it yourself
	q := req.URL.Query()
	q.Add("j_username", username)
	q.Add("j_password", password)
	req.URL.RawQuery = q.Encode()
	res, err := c.client.Do(req)
	return res, err
}

func (c *Client) GetUserDefaults() (*http.Response, error) {
	serviceUrl := fmt.Sprintf("%s/rest/PODService/AgilityUserDfltsPod/GetUserDefaults", c.url)
	payload := strings.NewReader(`{
	"request": {
		"AppData": "(Agility Mobile POD) BH POD Login/GetUserDefaults - 2022-07-08T13:48:30.543991",
		"UserIdentification": "bonzo"
	}
}`)

	req, _ := http.NewRequest("PUT", serviceUrl, payload)
	return c.client.Do(req)
}
