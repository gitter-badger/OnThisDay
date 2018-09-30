package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(time.Now())
	WebClient()
}

func WebClient() {
	port := ":8080"
	mux := http.DefaultServeMux
	fmt.Printf("Starting server on %v", port)
	mux.HandleFunc("/init", InitHandler)
	mux.HandleFunc("/oauth/redirect", RedirectHandler)
	// mux.HandleFunc("/auth", WorldHelloHandler)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Server is down!")
		fmt.Printf("%+v", err)
	}
}

func InitHandler(repwri http.ResponseWriter, req *http.Request) {
	http.ServeFile(repwri, req, "src/web_root/index.html")
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	uri := "https://slack.com/api/oauth.access?code=" + r.URL.Query().Get("code") + "&client_id=" + os.Getenv("CLIENT_ID") + "&client_secret=" + os.Getenv("CLIENT_SECRET") + "&redirect_uri=" + os.Getenv("REDIRECT_URI")
	resp, Geterr := http.Get(uri)
	if Geterr != nil {
		log.Fatal("Failed to authenticate")
	}
	var JSONResp interface{}
	jsonDecoder := json.NewDecoder(resp.Body)
	JSONerr := jsonDecoder.Decode(&JSONResp)
	if JSONerr != nil {
		log.Fatal("Failed to decode the body")
	}
	fmt.Printf("%+v", JSONResp)
}
