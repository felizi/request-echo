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
	fmt.Println("===================================================================")

	fmt.Println("=============================== HOST ==============================")
	fmt.Println(r.Host)

	fmt.Println("=============================== URL ===============================")
	fmt.Printf("%v %v %v\n", r.Method, r.URL, r.Proto)

	fmt.Println("=============================== HEADERS ===========================")
	for k, v := range r.Header {
		fmt.Printf("%s:%s\n", k, v)
	}

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("=============================== BODY ==============================")
	fmt.Printf("%s\n", body)

	fmt.Println("===================================================================")

	w.WriteHeader(200)
}
