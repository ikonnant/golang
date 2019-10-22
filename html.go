package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const TemplatesFolder = "templates/"

type ViewData struct {
	Title string
	Users []User
}

type User struct {
	Name string
	Age  int
}

func main() {
	printPage("/test/", "main", ViewData{
		Title: "Users List",
		Users: []User{
			User{Name: "Tom", Age: 21},
			User{Name: "Kate", Age: 23},
			User{Name: "Alice", Age: 25},
		},
	})

	printPage("/test2/", "main", ViewData{
		Title: "Users List 2",
		Users: []User{
			User{Name: "Tom", Age: 21},
			User{Name: "Kate", Age: 23},
			User{Name: "Alice", Age: 25},
		},
	})

	fmt.Println("Server is listening...")
	_ = http.ListenAndServe(":8080", nil)
}

func pageTemplate(page string) string {
	return TemplatesFolder + page + "/index.html"
}

func printPage(url string, page string, data ViewData) {
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles(pageTemplate(page))
		_ = tmpl.Execute(w, data)
	})
}
