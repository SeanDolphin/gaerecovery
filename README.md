#GAERecovery 

[![GoDoc](https://godoc.org/gopkg.in/SeanDolphin/gaerecovery.v1?status.png)](http://godoc.org/gopkg.in/SeanDolphin/gaerecovery.v1)  
[![Build Status](https://travis-ci.org/SeanDolphin/gaerecovery.svg?branch=master)](https://travis-ci.org/SeanDolphin/gaerecovery)  
[![Coverage Status](https://coveralls.io/repos/SeanDolphin/gaerecovery/badge.svg?branch=master)](https://coveralls.io/r/SeanDolphin/gaerecovery?branch=master)  
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

Google Appengine Recovery middleware for [Negroni](https://github.com/codegangsta/negroni).

Mostly a copy of the Negroni Recovery middleware with small changes to make it function
log into GAE correctly. 

## Installation

The import path for the package is *gopkg.in/SeanDolphin/gaerecovery.v1*.

To install it, run:

    go get gopkg.in/SeanDolphin/gaerecovery.v1
    
## Usage

~~~ go
package main

import (
    "fmt"
    "net/http"

    "github.com/codegangsta/negroni"
    "gopkg.in/SeanDolphin/gaerecovery.v1"
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
