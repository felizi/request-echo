package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	port := flag.String("p", "8888", "port of server")
	sleep := flag.Int("s", 0, "response sleep in milliseconds")
	flag.Parse()

	fmt.Printf("Request echo on port %s\n", *port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, *sleep)
	})

	fmt.Println(http.ListenAndServe(":"+*port, nil))

}

func handler(w http.ResponseWriter, r *http.Request, sleep int) {
	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Printf("Host: %v\n", r.Host)
	fmt.Printf("URL: %v %v %v\n", r.Method, r.URL, r.Proto)
	fmt.Println("Headers:")
	for k, v := range r.Header {
		fmt.Printf("%s:%s\n", k, v)
	}
	fmt.Printf("%s:%s\n", "Host", r.Host)

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Body: %s\n", body)

	fmt.Println("Cookies:")
	for _, cookie := range r.Cookies() {
		fmt.Printf("Cookie: %s\n", cookie.Name)
		fmt.Printf("Value: %s\n", cookie.Value)
	}

	time.Sleep(time.Millisecond * time.Duration(sleep))

	w.WriteHeader(200)
}
