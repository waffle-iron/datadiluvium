package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Page struct {
    Title string `json:"title"`
    Schema   string `json:"schema"`
}

func (p Page) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func main() {

    pages := getPages()
    for _, p := range pages {
        fmt.Println(p.toString())
        /* This should be parallelized for each of the schemas unless
        they're dependent and need to be executed in a specific order. */
    }

    fmt.Println(toJson(pages))
}

func getPages() []Page {
    raw, err := ioutil.ReadFile("./deluge.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []Page
    json.Unmarshal(raw, &c)
    return c
}
