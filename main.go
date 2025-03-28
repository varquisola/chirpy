package main

import (
    "log"
    "net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

    handler := http.NewServeMux()

    server := http.Server{
        Handler: handler,
        Addr:    ":8080",
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("Error with the server")
        return
    }

}
