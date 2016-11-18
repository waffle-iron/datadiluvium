package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Welcome to the early morning Amtrak Cascades, %s!", r.URL.Path[1:])
}

type Message struct {
   Text string
}

func about (w http.ResponseWriter, r *http.Request) {

   m := Message{"Welcome to the Twosi API, build v0.0.001."}
   b, err := json.Marshal(m)

   if err != nil {
       panic(err)
   }

    w.Write(b)
}

func main() {
   http.HandleFunc("/", handler)
   http.HandleFunc("/about/", about)
   log.Fatal(http.ListenAndServe(":8080", nil))
}
