package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
)

type RequestObject struct {
	Method string
	URL    *url.URL
	Header http.Header
	Data   map[string]interface{}
}

func main() {
	portFlag := *flag.String("port", "8077", "Port to listen on")
	listenAddress := *flag.String("address", "[::1]", "Address to listen on")

	http.HandleFunc("/api/v1/cache", func(w http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer request.Body.Close()
		var requestData map[string]interface{}
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			log.Fatal(err)
		}
		requestObject := RequestObject{
			Method: request.Method,
			URL:    request.URL,
			Header: request.Header,
			Data:   requestData,
		}
		log.Println(requestObject)
	})
	log.Fatal(http.ListenAndServe(listenAddress+":"+portFlag, nil))
}
