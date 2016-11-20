package main

import (
    "os"
    "io"
    "log"
    "net/http"
    "time"
)

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        logfile, err := os.OpenFile("./log/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
        if err != nil {
            panic("cannot open test.log:" + err.Error())
        }
        defer logfile.Close()

        log.SetOutput(io.MultiWriter(logfile, os.Stdout))
        log.SetFlags(log.Ldate | log.Ltime)

        q := r.URL.Query()

        log.Printf(
            "%s\t%s\t%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            q.Get("stamp"),
            q.Get("ref"),
            time.Since(start),
        )
    })
}
