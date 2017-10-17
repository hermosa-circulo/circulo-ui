package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/hermosa-circulo/circulo-tools/api"
)

func main() {
    func1 := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello")
    }
    router := api.NewRouter()
    router.HandleFunc("/api/", func1)
    //fmt.Println(router.routes)
    srv := &http.Server{
            Handler:      router,
            Addr:         "127.0.0.1:8000",
            WriteTimeout: 15 * time.Second,
            ReadTimeout:  15 * time.Second,
    }
    srv.ListenAndServe()
    fmt.Println("ok")
}
