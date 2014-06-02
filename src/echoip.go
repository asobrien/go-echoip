// echoip.go
//
// A small program to echo the IP of a requesting client.

package main

import (
    "net/http"
    "fmt"
    "strings"
    "flag"
)

var (
    httpAddr = flag.String("port", "9999", "HTTP server port to run echoip")
)

func echoIp(w http.ResponseWriter, r *http.Request) {
    ipRequest := r.RemoteAddr
    lastIndexSep := strings.LastIndex(r.RemoteAddr,":")
    ip := ipRequest[:lastIndexSep]
    fmt.Fprintf(w, "%s", ip)
}

func main() {
    flag.Parse()
    fullHttpAddr := ":" + *httpAddr
    http.HandleFunc("/", echoIp)
    http.ListenAndServe(fullHttpAddr, nil)
}