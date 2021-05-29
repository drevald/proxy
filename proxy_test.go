package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"os"
)

func TestAbs(t *testing.T) {
    got := math.Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %f; want 1", got)
    }
}

func TestHttpRequest(t *testing.T) {
	go main()	
	fmt.Println("Starting test")
	var reader io.Reader = strings.NewReader("some io.Reader stream to be read\n")
	req := httptest.NewRequest("GET", "localhost:8001", reader)
	fmt.Println("response", req.Body)
}

func TestServer(t *testing.T) {

	fmt.Println("TestServer")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

func TestProxiedClient(t *testing.T) {
	fmt.Println("TestProxiedClient")
	go main()
	os.Setenv("HTTP_PROXY", "http://localhost:8001")
	proxyUrl, _ := url.Parse("http://localhost:8001")
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	_,_ = myClient.Get("localhost:8001")
}