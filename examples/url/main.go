package main

import (
	"fmt"
	"net/url"
)

func main() {
	// q := url.Values{"path": []string{"a"}}
	// u := url.URL{
	// 	Scheme:   "https",
	// 	Host:     "example.com",
	// 	RawQuery: q.Encode(),
	// }
	u := url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "v1/endpoint",
	}

	fmt.Println(u.String())
}
