package main

import (

    "net/http"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "os"
)

func main() {
    router := mux.NewRouter()
    loggedRouter := handlers.LoggingHandler(os.Stdout, router)
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

    router.HandleFunc("/plush-file-server/addProfilePicture", addProfilePicture).Methods("POST")
    //POST method is used, as a body is required to send over the path of the profile picture
    router.HandleFunc("/plush-file-server/getProfilePicture", getProfilePicture).Methods("POST")
    http.ListenAndServe(":8001", handlers.CORS(headersOk, methodsOk, originsOk)(loggedRouter))
}
