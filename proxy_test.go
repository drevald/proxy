package main

import (
	"fmt"
	"testing"
	"math"
	"net/http/httptest"
	"io"
	"strings"
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

