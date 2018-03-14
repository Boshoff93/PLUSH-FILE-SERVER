package main

import (
  b64 "encoding/base64"
  "encoding/json"
  "fmt"
  "net/http"
  "strings"
  "os"
)

func addProfilePicture(w http.ResponseWriter, r *http.Request){

  var blob Blob
  if err := json.NewDecoder(r.Body).Decode(&blob); err != nil {
          http.Error(w, err.Error(), 400)
          return
  }
  // Spliiting up data:image/jpeg;base64,/9j/ffdgfd...
  imageParts := strings.Split(blob.Data, ",")
  string64 := imageParts[1]
  sDec, err := b64.StdEncoding.DecodeString(string64)
  finished := make(chan bool)
  if err != nil {
    fmt.Println("Could not decode profile picture, error: " + err.Error())
    json.NewEncoder(w).Encode(Error{Error: err.Error()})
    finished <- true
    return
  }

  go func() {
    path := "./images/profile_pictures/"
    os.MkdirAll(path, os.ModePerm)
    check(err)
    fileDec, err := os.Create(path + blob.Pp_Name)
    check(err)
    defer fileDec.Close()
    _, err = fileDec.Write(sDec)
    check(err)
    fileDec.Sync()
    json.NewEncoder(w).Encode(blob)
    finished <- true
  }()
  <- finished
}
