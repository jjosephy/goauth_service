package main

import (
    "fmt"
    //"github.com/dgrijalva/jwt-go"
    "github.com/jjosephy/goauth_service/handler"
    //"io/ioutil"
    //"log"
    //"time"
    "net/http"
)

const (
    PORT       = ":8443"
    PRIV_KEY   = "./private_key"
    PUBLIC_KEY = "./cert.pem"
)

// Main entry point used to set up routes //
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/token", handler.TokenHandler())
    err := http.ListenAndServeTLS(PORT, PUBLIC_KEY, PRIV_KEY, mux)
    if err != nil {
       fmt.Printf("main(): %s\n", err)
    }
}
