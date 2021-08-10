package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageContent struct {
	Title   string
	Content string
}

func render_templete(res http.ResponseWriter, html_page string, page *PageContent) {
	/*
		Takes a response writer(Response) as the first argument,html page and content to be embeded in html page.
	*/
	html_page = "./templates/" + html_page
	t, err := template.ParseFiles(html_page)

	// Check if theres an error retrieving html page
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// Check if there is an error rendering html page
	err = t.Execute(res, page)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	/*
		Takes http response writer and http request.
		Returns html page with content
	*/
	Page := &PageContent{
		Title:   "Go Web App",
		Content: "Welcome to my first Go web app",
	}
	fmt.Printf("HELLO %v", Page.Title)
	render_templete(res, "index.html", Page)
}

func MainHandler() {
	http.HandleFunc("/", index)
}
