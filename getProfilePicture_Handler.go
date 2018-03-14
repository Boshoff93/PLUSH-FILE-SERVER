package main

import (
  b64 "encoding/base64"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "github.com/gorilla/mux"
)

func getProfilePicture(w http.ResponseWriter, r *http.Request){

  params:= mux.Vars(r)
  var blob Blob
  blob.Pp_Name = params["pp_name"]

  finished := make(chan bool)
  go func() {
    var string64 []byte
    path := "./images/profile_pictures/"
    string64,_ = ioutil.ReadFile(path + blob.Pp_Name)
    encodedString := b64.StdEncoding.EncodeToString(string64)
    var blob Blob
    blob.Data = encodedString
    json.NewEncoder(w).Encode(blob)
    finished <- true
  }()
  <- finished
}
