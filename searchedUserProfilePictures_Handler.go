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

func searchedUserProfilePictures(w http.ResponseWriter, r *http.Request){
  fmt.Println("got here")
  params:= mux.Vars(r)
  var multiBlob MultiBlob
  var pp_names_string = params["pp_names"]
  fmt.Println(pp_names_string)
  if(pp_names_string != "null") {
    multiBlob.Pp_Names = strings.Split(pp_names_string, ",,"); // for one name
    if(len(multiBlob.Pp_Names) == 1) {
      multiBlob.Pp_Names[0] = pp_names_string[0:len(pp_names_string)-1]
    }
    finished := make(chan bool)
    go func() {
      for _, element := range multiBlob.Pp_Names {
        if (element == "not found") {
          multiBlob.Data = append(multiBlob.Data, "empty")
        } else {
          fmt.Println("this is the:"+ element)
          var string64 []byte
          path := "./images/profile_pictures/"
          string64, _ = ioutil.ReadFile(path + element)
          encodedString := b64.StdEncoding.EncodeToString(string64)
          multiBlob.Data = append(multiBlob.Data, "data:image/jpeg;base64,"+ encodedString)
        }
      }
      fmt.Println(multiBlob.Data)
      json.NewEncoder(w).Encode(multiBlob)
      finished <- true
    }()
    <- finished
  } else {
    var empty []string
    json.NewEncoder(w).Encode(empty)
  }
}
