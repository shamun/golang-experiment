package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "path/filepath"
)

const http_root = "./public"
const server_port = ":55555"

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/html/", htmlHandler)

    http.ListenAndServe(server_port, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "rootHandler: %s\n", r.URL.Path)
    fmt.Fprintf(w, "URL: %s\n", r.URL)
    fmt.Fprintf(w, "Method: %s\n", r.Method)
    fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
    fmt.Fprintf(w, "Proto: %s\n", r.Proto)
    fmt.Fprintf(w, "HOST: %s\n", r.Host)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "viewHandler: %s", r.URL.Path)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "htmlHandler: %s\n", r.URL.Path)

    filename := http_root + r.URL.Path
    fileext := filepath.Ext(filename)

    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Fprintf(w, " 404 Not Found!\n")
        w.WriteHeader(http.StatusNotFound)
        return
    }

    var contype string
    switch fileext {
        case ".html", ".htm":
            contype = "text/html"
        case ".css":
            contype = "text/css"
        case ".js":
            contype  = "application/javascript"
        case ".png":
            contype = "image/png"
        case ".jpg", ".jpeg":
            contype = "image/jpeg"
        default:
            contype = "text/plain"
    }
    fmt.Fprintf(w, "ext %s, ct = %s\n", fileext, contype)

    w.Header().Set("Content-Type", contype)
    fmt.Fprintf(w, "%s", content)
}
