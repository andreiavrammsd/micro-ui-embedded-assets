package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"path/filepath"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")

	endpoint, exists := endpoints[path]
	if exists == false {
		HomeHandler(w, r)
		return
	}

	u, err := url.Parse(endpoint)
	checkError(err)

	q := r.URL.Query()
	offset := r.URL.Query().Get("offset")
	limit := strconv.Itoa(getLimit(r.URL.Query().Get("limit")))
	q.Set("limit", limit)
	u.RawQuery = q.Encode()

	log.Println(u.String())

	resp, err := http.Get(u.String())
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	data := &Data{}
	err = json.Unmarshal(body, &data.Items)
	checkError(err)

	data.Endpoints = endpoints
	data.CurrentEndpoint = path

	data.Pagination = &Pagination{
		Offset: offset,
		Limit:  limit,
		Prev:   getPaginationUrl(offset, limit, "prev"),
		Next:   getPaginationUrl(offset, limit, "next"),
	}

	templateStringPosition := strings.LastIndex(path, ":") + 1
	templatePath := path[templateStringPosition:len(path)]
	assets := []string{
		fmt.Sprintf("template/%s.html", templatePath),
		"template/footer.html",
		"template/header.html",
		"template/nav.html",
		"template/pagination.html",
	}
	content, err := getAssetsContent(assets)
	checkError(err)

	t, err := template.New(path + ".html").Funcs(getFuncMap()).Parse(string(content))
	checkError(err)

	t.Execute(w, data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	assets := []string{
		"template/home.html",
		"template/nav.html",
	}
	content, err := getAssetsContent(assets)
	checkError(err)

	t, err := template.New("home.html").Funcs(getFuncMap()).Parse(string(content))
	checkError(err)

	data := &Data{
		Endpoints: endpoints,
	}
	t.Execute(w, data)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	ext := filepath.Ext(path)
	
	content, err := getAssetsContent([]string{path})
	checkError(err)
	
	contentTypes := map[string]string{
		"css": "text/css",
		"js": "text/javascript",
	}
	
	w.Header().Set("Content-Type", contentTypes[ext])
	w.Write(content)
}
