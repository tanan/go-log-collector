package main

import (
    "net/http"
    "time"

)

func Collect(w http.ResponseWriter, r *http.Request) {
    stamp := r.URL.Query().Get("stamp")
    if stamp == "" {
        stamp := generateUID()
        expiration := time.Now().Add(365 * 24 * time.Hour)
        cookie := http.Cookie{
            Name: "stamp",
            Value: stamp,
            Expires: expiration,
        }
        http.SetCookie(w, &cookie)
    }

    vis := r.URL.Query().Get("vis")
    if vis == "" {
        vis := generateVID()
        expiration := time.Now().Add(30 * time.Minute)
        cookie := http.Cookie{
            Name: "vis",
            Value: vis,
            Expires: expiration,
        }
        http.SetCookie(w, &cookie)
    }
}
