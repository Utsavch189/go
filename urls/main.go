package main

import (
	"fmt"
	"net/url"
)

func main() {
	url1 := "http://localhost:5000/api?id=1&limit=100"

	// Get parts of url1
	resullt, _ := url.Parse(url1)
	fmt.Println(resullt.Scheme)
	fmt.Println(resullt.Host)
	fmt.Println(resullt.Hostname())
	fmt.Println(resullt.Port())
	fmt.Println(resullt.Path)
	fmt.Println(resullt.RawQuery)
	fmt.Println(resullt.Query())

	/*
		OUTPUT:
			http
			localhost:5000
			localhost
			5000
			/api
			id=1&limit=100
			map[id:[1] limit:[100]]
	*/

	// Construct a url from parts
	myUrl := &url.URL{
		Scheme:   "http",
		Host:     "localhost:6000",
		Path:     "/api/v1",
		RawQuery: "id=2&limit=478",
	}
	fmt.Println(myUrl) // http://localhost:6000/api/v1?id=2&limit=478
}
