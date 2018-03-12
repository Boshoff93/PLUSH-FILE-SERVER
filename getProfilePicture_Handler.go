package main

import (
  b64 "encoding/base64"
  "encoding/json"
  "net/http"
  "bufio"
  "os"
  "fmt"
)

func getProfilePicture(w http.ResponseWriter, r *http.Request){

  var blob Blob
  if err := json.NewDecoder(r.Body).Decode(&blob); err != nil {
          http.Error(w, err.Error(), 400)
          return
  }

  finished := make(chan bool)
  go func() {
    var htmlEmbed string
    var string64 []byte

    file, err := os.Open(blob.Path)
    check(err)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      lines = append(lines, scanner.Text())
    }
    htmlEmbed = lines[0]
    string64 = []byte(lines[1])

    encodedString := b64.StdEncoding.EncodeToString(string64)
    //Constructing html base64 embeded image
    var base64EmbededImage = htmlEmbed + "," + encodedString
    var blob Blob
    blob.Data = base64EmbededImage
    json.NewEncoder(w).Encode(blob)
    finished <- true
  }()
  <- finished
}
