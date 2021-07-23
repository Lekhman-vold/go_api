package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}



type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Go go go!!!"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("HomePage"))
}

func handleRequest() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/articles", allArticles)
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func main()  {
	const HOST = "127.0.0.1"
	const PORT = ":8080"
	log.Printf("Server started >>> %v%v", HOST, PORT)
	handleRequest()
}