package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	port := flag.String("p", "8888", "port of server")
	flag.Parse()

	fmt.Printf("Request echo on port %s\n", *port)

	http.HandleFunc("/", handler)

	fmt.Println(http.ListenAndServe(":"+*port, nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Printf("Host: %v\n", r.Host)
	fmt.Printf("URL: %v %v %v\n", r.Method, r.URL, r.Proto)
	fmt.Println("Headers:")
	for k, v := range r.Header {
		fmt.Printf("%s:%s\n", k, v)
	}

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Body: %s\n", body)

	fmt.Println("Cookies:")
	for _, cookie := range r.Cookies() {
		fmt.Printf("Cookie: %s\n", cookie.Name)
		fmt.Printf("Value: %s\n", cookie.Value)
	}

	w.WriteHeader(200)
}
