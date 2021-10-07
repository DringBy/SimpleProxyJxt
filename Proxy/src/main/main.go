package main

import (
	"fmt"
	"net/http"
	"net/url"
)

var (
	//这里是否可以变成直接读取conf
	ProxyAddr = "http://127.0.0.1"
	port2 = "7000"
	port1 = "8000"
)

func SimpleReverseProxy(w http.ResponseWriter, r *http.Request) {
	ProxyURL, err := url.Parse(ProxyAddr+":"+port2)
	if err != nil{
		fmt.Println("URL err=", err)
		return
	}
	fmt.Println("Proxy URL: ",ProxyURL)
	// fmt.Println(ProxyURL.Scheme)
	// fmt.Println(ProxyURL.Host)
	r.URL.Scheme = ProxyURL.Scheme
	r.URL.Host = ProxyURL.Host
	// fmt.Println(r.URL.Scheme)
	// fmt.Println(r.URL.Host)
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		fmt.Println("Transport err=", err)
		return
	}
	fmt.Println("response=", resp.Status)
	defer resp.Body.Close()
}

func main(){
	http.HandleFunc("/api/v1/healthy", SimpleReverseProxy)
	fmt.Println("Start Serve on Port: ", port1)
	err := http.ListenAndServe("0.0.0.0:"+port1, nil)
	if err != nil {
		fmt.Println("Listen err=", err)
	}
	
}