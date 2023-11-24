package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main(){
    m := http.NewServeMux() //spin up a new server instance

    const addr = ":8080" //should go in an .env file

    m.HandleFunc("/", handlePage) // handles the handler, which in this case is rendered http

    // standard lib. lets you have more control over the server's behaviors, so for standard lib
    // option this is probably best to link it up with env variables, or other things needed
    // ** Research all the options. Is it smart to have to option for all? Or to just nix it as
    // much as I can to keep it minimal?
    s := &http.Server{
        Addr: addr,
        Handler: m, 
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    fmt.Println("hi, the server hast started") // test string that prints in console to let ya know
    err := s.ListenAndServe() // initiate server, this way keeps it going untill it hits an error
    log.Fatal(err) // error logging, it's important
}

// handler function, this is what you render via HandleFunc up in main, so this can be extracted out
// into it's own file that is passed over here. This can keep it cleaner and separate in case you want
// to add more functions depending on what you do with the scaffold.
func handlePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "I am skynet")
    w.WriteHeader(200)
}
