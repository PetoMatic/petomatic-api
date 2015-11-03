package main

import (
    "log"
    "net/http"
    "os"
)

func checkError(desc string, err error) {
    if err != nil {
        log.Fatal("%s %s", desc, err)
        os.Exit(1)
    }
}


func main() {
    checkError("Error connecting to shopsdb:", InitDBConn("127.0.0.1", "5432", "petomaticdb", "petomatic"))
    rrr := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", rrr))
}
