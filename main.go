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

    router.HandleFunc("/plush-file-server/profilePicture", ValidateMiddleware(addProfilePicture)).Methods("POST")
    router.HandleFunc("/plush-file-server/postImage", ValidateMiddleware(addPostPicture)).Methods("POST")
    //POST method is used, as a body is required to send over the path of the profile picture
    router.HandleFunc("/plush-file-server/profilePicture/{pp_name}", ValidateMiddleware(getProfilePicture)).Methods("GET")

    router.HandleFunc("/plush-file-server/searchedUserProfilePictures/{pp_names}", ValidateMiddleware(searchedUserProfilePictures)).Methods("GET")
    router.HandleFunc("/plush-file-server/getPostImages/{posts_with_images}", ValidateMiddleware(getPostImages)).Methods("GET")
    router.HandleFunc("/plush-file-server/removePostPicture", ValidateMiddleware(removePostPicture)).Methods("DELETE")
    http.ListenAndServe(":8001", handlers.CORS(headersOk, methodsOk, originsOk)(loggedRouter))
}
