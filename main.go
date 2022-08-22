package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

var reverseUrl string
var port string
var printRequest bool

func printUsage() {
	fmt.Println("usage: reverse-proxy <url> <port> <print request>")
	fmt.Println("       required <url>, string")
	fmt.Println("       optional <port>, integer, default=1338")
	fmt.Println("       optional <print request>. bool [true, false]. requires port")
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host
	proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s proxy_url: %s%s\n", req.Method, reverseUrl, req.URL)
	if printRequest {
		fmt.Println(req)
	}
	serveReverseProxy(reverseUrl, res, req)
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("no url to mirror")
		printUsage()
		os.Exit(1)
	}
	reverseUrl = args[1]
	if len(args) > 2 {
		port = args[2]
		if len(args) > 3 {
			var err error
			if printRequest, err = strconv.ParseBool(args[3]); err != nil {
				fmt.Println("third param must be true or false")
				printUsage()
				os.Exit(1)
			}
		} else {
			printRequest = false
		}
	} else {
		port = "1338"
		fmt.Printf("no port provided. using fallback port :%s\n", port)
	}
	fmt.Printf("domain %s now serving on localhost:%s\n", reverseUrl, port)
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}
