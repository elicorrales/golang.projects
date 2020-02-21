package main


import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)


var value  = "0"

func post_value(writer http.ResponseWriter , request *http.Request ) {
    vars := mux.Vars(request)
    fmt.Println("method:" + request.Method)
    fmt.Println(request.URL)
    fmt.Println(vars)
    fmt.Println(vars["value"])
    value = vars["value"]
    fmt.Fprintf(writer, "Hello")
}


func main() {

    rtr := mux.NewRouter()

    rtr.HandleFunc("/value/{value:[0-9]+}",post_value)

    http.Handle("/",rtr)

    http.ListenAndServe(":8090",nil)
}
