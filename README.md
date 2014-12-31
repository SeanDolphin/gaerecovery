#GAERecovery 
[![GoDoc](https://godoc.org/github.com/SeanDolphin/gaerecovery?status.png)](http://godoc.org/github.com/SeanDolphin/gaerecovery)

Google Appengine Recovery middleware for [Negroni](https://github.com/codegangsta/negroni).

Mostly a copy of the Negroni Recovery middleware with small changes to make it function
log into GAE correctly. 

## Usage

~~~ go
package main

import (
    "fmt"
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/SeanDolphin/gaerecovery"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    	  fmt.Fprintf(w, "Welcome to the home page!")
    })

    n := negroni.New()
    n.Use(gaerecovery.Recovery())
    n.UseHandler(mux)
    n.Run(":3000")
}
~~~
