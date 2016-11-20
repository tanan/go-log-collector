package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

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
}
