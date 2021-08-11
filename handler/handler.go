package handler

import (
	"html/template"
	"net/http"
)

type PageContent struct {
	Title string
	About string
}

func render_templete(res http.ResponseWriter, html_page string, page *PageContent) {
	/*
		Takes a response writer(Response) as the first argument,html page and content to be embeded in html page.
	*/
	html_page = "./templates/" + html_page
	t, err := template.ParseFiles(html_page)

	/*
		Check if theres an error retrieving html page
	*/
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	/*
		Check if there is an error rendering html page
	*/
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
		Title: "Go Web App",
		About: "One of Golang’s biggest advantages is that it offers the clarity and ease-of-use that other languages lack. Golang’s advantages make it easy for new programmers to quickly understand the language and for seasoned veterans to easily read each other’s code.",
	}

	render_templete(res, "index.html", Page)
}

func MainHandler() {
	http.HandleFunc("/", index)
}
