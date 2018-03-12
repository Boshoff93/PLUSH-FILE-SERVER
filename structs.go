package main

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Error struct {
  Error    string               `json: "error"`
}

type Blob struct {
 User_Id  string              `json:user_id`
 Path     string              `json:"path"`
 Data     string              `json:"data"`
}
