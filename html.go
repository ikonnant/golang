package main

import (
  "fmt"
  "net/http"
  "html/template"
)

const templateURL = "templates/"

type ViewData struct {
  Title string
  Users []User
}

type User struct {
  Name string
  Age int
}

func main() {
  data := ViewData {
    Title : "Users List",
    Users : []User {
      User {Name: "Tom", Age: 21},
      User {Name: "Kate", Age: 23},
      User {Name: "Alice", Age: 25},
    },
  }

  tmplParse("/", "main", data)

  fmt.Println("Server is listening...")
  http.ListenAndServe(":8080", nil)
}

func templateFn(page string) string {
  return templateURL + page + "/index.html"
}

func tmplParse(url string, page string, data ViewData) {
  http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles(templateFn(page))
    tmpl.Execute(w, data)
  })
}