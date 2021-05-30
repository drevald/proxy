package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"bufio"
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
//	os.Setenv("HTTP_PROXY", "http://localhost:8001")
	proxyUrl, _ := url.Parse("http://localhost:8001")
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	_,err := myClient.Get("http://ya.ru:80")
	if err != nil {
		fmt.Println("Fail to access remote address via proxy", err.Error())
	}
}


func TestClient(t *testing.T) {
	fmt.Println("Test Client")
	go main()

	myClient := http.Client{}
	resp, err := myClient.Get("http://ya.ru:80")
	if err != nil {
		fmt.Println("Fail to access remote address", err.Error())
	}
	fmt.Println(resp)
}

func TestTcpClient(t *testing.T) {
	servAddr := "ya.ru:80"
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    co,err := net.DialTCP("tcp", nil, tcpAddr)
    fmt.Println("TCP Address is", tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    } 
	
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(co, text+"\n")

		message, _ := bufio.NewReader(co).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
				fmt.Println("TCP client exiting...")
				return
		}
	}


    //buf := make([]byte, 100)
    // // for {
    //     _, err = co.Read(buf)
    //     if err != nil {
    //         fmt.Println("Fail to read from server " + servAddr, err.Error())
    //     }
    //     fmt.Println(buf)
    // // }	
}

func TestTcpSocket(t *testing.T) {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println("Error", err.Error());
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
