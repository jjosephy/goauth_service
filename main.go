package main

import (
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "io/ioutil"
    "log"
    "time"
)

var (
    jwtTestDefaultKey []byte
    defaultKeyFunc    jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return jwtTestDefaultKey, nil }
)

func main() {
    fmt.Println("here")
    var e error


    if jwtTestDefaultKey, e = ioutil.ReadFile("cert.pem"); e != nil {
        panic(e)
    }

    // Create the token
	tokens := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	tokens.Claims["foo"] = "bar"
	tokens.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, ex := tokens.SignedString(jwtTestDefaultKey)

    if ex != nil {
        log.Printf("Error %v", ex)
        return
    }

    log.Println(tokenString)
    log.Println("")

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return defaultKeyFunc(tokens)
    })

    fmt.Println()
    if err == nil && token.Valid {
        fmt.Println(token.Claims["foo"])
    } else {
        fmt.Println("err")
    }
}
