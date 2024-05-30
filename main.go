package main

import (
    "fmt"
    "html/template"
    "net/http"
)

var templates = template.Must(template.ParseFiles("templates/base.html"))

func index() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        b := struct {
            Title        template.HTML
            BusinessName string
            Slogan       string
        }{
            Title:        template.HTML("Ex3_Week6, AI & GPT"),
            BusinessName: "Business,",
            Slogan:       "we get things done!",
        }
        err := templates.ExecuteTemplate(w, "base", &b)
        if err != nil {
            http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    })
}

func main() {
    http.Handle("/", index())
    fmt.Println("Starting server at :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Printf("main: couldn't start server: %v\n", err)
    }
}
