package main

import (
    "flag"
    "fmt"
    "net/http"
    "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world!")
}

func main() {
    port:= flag.Int("port", 8080, "port number")
    flag.Parse()

    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + strconv.Itoa(*port), nil)
}
