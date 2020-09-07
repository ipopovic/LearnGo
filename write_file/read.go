package main

import (
    "fmt"
    "io"
    "os"
)

var path = "test.txt"

func main() {

	readFile()
}

func readFile() {
    // Open file for reading.
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    // Read file, line by line
    var text = make([]byte, 1024)
    for {
        _, err = file.Read(text)

        // Break if finally arrived at end of file
        if err == io.EOF {
            break
        }

        // Break if error occured
        if err != nil && err != io.EOF {
            isError(err)
            break
        }
    }

    fmt.Println("Reading from file.......")
    fmt.Println(string(text))
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}
