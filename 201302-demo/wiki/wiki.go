// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "errors"
    "regexp"
    "io/ioutil"
    "net/http"
    "html/template"
)

const lenPath = len("/view/")

type Page struct {
    Title string
    Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func getTitle(w http.ResponseWriter, r *http.Request) (title string, err error) {
    title = r.URL.Path[lenPath:]
    if !titleValidator.MatchString(title) {
        http.NotFound(w, r)
        err = errors.New("Invalid Page Title")
    }
    return
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    //t, err := template.ParseFiles(tmpl + ".html")
    //if err != nil {
    //    http.Error(w, err.Error(), http.StatusInternalServerError)
    //    return
    //}
    //err = t.Execute(w, p)
    //if err != nil {
    //    http.Error(w, err.Error(), http.StatusInternalServerError)
    //}
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    saveerr := p.save()
    if saveerr != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":55556", nil)
}
