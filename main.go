package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	port := flag.String("p", "8888", "port of server")
	sleep := flag.Int("s", 0, "response sleep in milliseconds")
	failureRate := flag.Int("f", 0, "failure rate between 0 and 100")
	failureStatusCode := flag.Int("fsc", 400, "status code to response when failure rate reached- default 400")
	statusCode := flag.Int("sc", 200, "status code to response - default 200")
	flag.Parse()

	fmt.Printf("Request echo on port %s\n", *port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, *failureRate, *failureStatusCode, *sleep, *statusCode)
	})

	fmt.Println(http.ListenAndServe(":"+*port, nil))

}

func handler(w http.ResponseWriter, r *http.Request, failureRate, failureStatusCode, sleep, statusCode int) {
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

	cors(w, r)

	if mustFail(failureRate) {
		w.WriteHeader(failureStatusCode)
	} else {
		time.Sleep(time.Millisecond * time.Duration(sleep))
		w.WriteHeader(statusCode)
	}
}

func mustFail(rate int) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	return randInt(0, 100) < rate
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func cors(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("access-control-allow-credentials", "true")
		w.Header().Set("access-control-allow-headers", r.Header.Get("access-control-request-headers"))
		w.Header().Set("access-control-request-method", r.Header.Get("access-control-request-method"))
		w.Header().Set("access-control-allow-origin", r.Header.Get("origin"))
	}
}
