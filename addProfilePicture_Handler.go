package main

import (
  b64 "encoding/base64"
  "encoding/json"
  "fmt"
  "net/http"
  "strings"
  "os"
)

type Error struct {
  Error    string               `json: "error"`
}

type Blob struct {
 User_Id  string              `json:user_id`
 Path     string              `json:"path"`
 Data     string              `json:"data"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func addProfilePicture(w http.ResponseWriter, r *http.Request){

  var blob Blob
  if err := json.NewDecoder(r.Body).Decode(&blob); err != nil {
          http.Error(w, err.Error(), 400)
          return
  }
  // Spliiting up data:image/jpeg;base64,/9j/ffdgfd...
  imageParts := strings.Split(blob.Data, ",")
  htmlEmbed := imageParts[0]
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
    os.MkdirAll(blob.Path, os.ModePerm)
    check(err)
    file, err := os.Create(blob.Path + blob.User_Id + "_pp_image")
    check(err)
    defer file.Close()
    _, err = file.WriteString(htmlEmbed+"\n")
    check(err)
    _, err = file.Write(sDec)
    check(err)
    file.Sync()

    finished <- true
    json.NewEncoder(w).Encode(blob)
  }()
  <- finished
}
