package main

import (
  b64 "encoding/base64"
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/gorilla/mux"
  "strings"
)

func getPostImages(w http.ResponseWriter, r *http.Request){
  params:= mux.Vars(r)
  var postImages Post_Images
  var posts_names_string = params["posts_with_images"]
  fmt.Println("Right here")
  fmt.Println(posts_names_string)
  if(posts_names_string != "null") {
    postImages.Post_Names = strings.Split(posts_names_string, ","); // for one name
    finished := make(chan bool)
    go func() {
      for _, element := range postImages.Post_Names {
        fmt.Print(element)
        if (element == "empty") {
          postImages.Data = append(postImages.Data, "empty")
        } else {
          fmt.Println(element)
          var string64 []byte
          path := "./images/post_pictures/"
          string64, _ = ioutil.ReadFile(path + element)
          encodedString := b64.StdEncoding.EncodeToString(string64)
          postImages.Data = append(postImages.Data, "data:image/jpeg;base64,"+ encodedString)
        }
      }
      json.NewEncoder(w).Encode(postImages)
      finished <- true
    }()
    <- finished
  } else {
    var empty []string
    json.NewEncoder(w).Encode(empty)
  }
}
