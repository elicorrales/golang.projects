package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "time"
    "sync"
)

const port = "8090"

var mutex = &sync.Mutex{}

var value  = "0"

func post_value(writer http.ResponseWriter , request *http.Request ) {
    vars := mux.Vars(request)
    valueIn := vars["value"]
    //mutex.Lock()
    value = valueIn
    time.Sleep(1 * time.Millisecond)
    valueNow := value
    //mutex.Unlock()
    if valueNow == valueIn {
        fmt.Fprintf(writer, ".")
    } else {
        fmt.Print("E")
        fmt.Fprintf(writer, "E")
    }
}


func main() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/value/{value:[0-9]+}",post_value)
    http.Handle("/",rtr)
    fmt.Println("Server up and listening at " + port)
    http.ListenAndServe(":" + port,nil)
}
