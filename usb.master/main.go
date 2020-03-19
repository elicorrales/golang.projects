package main

import (
        "io/ioutil"
        "os"
)

func main() {

    files, err := ioutil.ReadDir("/dev")
    if err != nil {
        println(err)
        os.Exit(1)
    }

    for _, file := range files {
        println(file)
    }
}
