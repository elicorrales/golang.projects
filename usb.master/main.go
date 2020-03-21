package main

import (
        "io/ioutil"
        "os"
)

func main() {

    files, err := ioutil.ReadDir("/dev")
    //files, err := ioutil.ReadDir("c:/Users/devchu")
    if err != nil {
        println(err)
        os.Exit(1)
    }

    for _, file := range files {
        println(file.Name())
    }
}
