package main

import (
  b64 "encoding/base64"
  "encoding/json"
  "fmt"
  "net/http"
  "strings"
  "os"
)

func addPostPicture(w http.ResponseWriter, r *http.Request){

  var post_image Post_Image
  if err := json.NewDecoder(r.Body).Decode(&post_image); err != nil {
          http.Error(w, err.Error(), 400)
          return
  }
  // Spliiting up data:image/jpeg;base64,/9j/ffdgfd...
  imageParts := strings.Split(post_image.Data, ",")
  string64 := imageParts[1]
  sDec, err := b64.StdEncoding.DecodeString(string64)
  finished := make(chan bool)
  if err != nil {
    fmt.Println("Could not decode post picture, error: " + err.Error())
    json.NewEncoder(w).Encode(Error{Error: err.Error()})
    finished <- true
    return
  }

  go func() {
    path := "./images/post_pictures/"
    os.MkdirAll(path, os.ModePerm)
    check(err)
    fileDec, err := os.Create(path + post_image.Image_Name)
    check(err)
    defer fileDec.Close()
    _, err = fileDec.Write(sDec)
    check(err)
    fileDec.Sync()
    json.NewEncoder(w).Encode(post_image)
    finished <- true
  }()
  <- finished
}
