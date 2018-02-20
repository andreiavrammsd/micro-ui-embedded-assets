package main

import (
	"html/template"
	"reflect"
	"strings"
	"time"
)

type Data struct {
	Items           []interface{}
	Endpoints       map[string]string
	CurrentEndpoint string
	Pagination      *Pagination
}

type Pagination struct {
	Offset string
	Limit  string
	Prev   string
	Next   string
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"datetime": datetime,
		"add":      add,
		"title":    title,
	}
}

func datetime(timestamp interface{}) string {
	if timestamp != nil && reflect.TypeOf(timestamp).String() == "float64" {
		ts := reflect.ValueOf(timestamp).Float()
		t := time.Unix(int64(ts), 0)

		return t.Format(timeFormat)
	}

	return "not set"
}

func add(args ...interface{}) float64 {
	var sum float64

	for _, arg := range args {
		sum += reflect.ValueOf(arg).Float()
	}

	return sum
}

func title(arg string) string {
	return strings.Replace(strings.Title(strings.Replace(arg, "-", " ", -1)), ":", " ", -1)
}
