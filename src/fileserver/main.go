package main

import (
    "net/http"
    "fmt"
    "github.com/russross/blackfriday"
)

func main() {
    http.HandleFunc("/markdown", GenerateMarkdown)
    // http.Handle("/", http.FileServer(http.Dir("public")))
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

//needs some work
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
    rw.Write(markdown)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
}