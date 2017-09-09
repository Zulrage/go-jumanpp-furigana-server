package main

import (
    // Standard library packages
    "net/http"
    "fmt"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "github.com/rs/cors"
    "gopkg.in/mgo.v2"
    "app/controllers"
)

func main() {
    // Instantiate a new router
    r := httprouter.New()

    // Get a UserController instance
    bc := controllers.NewFuriganaController(getSession())

    // Get a user resource
    r.GET("/furigana/:id", bc.GetFurigana)

    r.GET("/furigana", bc.GetAllFurigana)

    r.POST("/furigana", bc.CreateFurigana)

    r.DELETE("/furigana/:id", bc.RemoveFurigana)

    fmt.Println("Server up")
    // Fire up the server
    handler := cors.Default().Handler(r)
    fmt.Println(http.ListenAndServe("0.0.0.0:3003", handler))
}

func getSession() *mgo.Session {
    // Connect to our local mongo
    dbPlace := "mongodb://db:27017"
    s, err := mgo.Dial(dbPlace)

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    fmt.Println("Mongo session up at " + dbPlace)
    return s
}
