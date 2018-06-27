package main

import (
  "encoding/json"
  "net/http"
  "os"
  "fmt"
)

func removePostPicture(w http.ResponseWriter, r *http.Request){

  var post_image Post_Image
  if err := json.NewDecoder(r.Body).Decode(&post_image); err != nil {
          http.Error(w, err.Error(), 400)
          return
  }

  fmt.Println(post_image)
  finished := make(chan bool)

  go func() {
    path := "./images/post_pictures/"
    os.Remove(path + post_image.Image_Name)

    json.NewEncoder(w).Encode(post_image)
    finished <- true
  }()
  <- finished
}
