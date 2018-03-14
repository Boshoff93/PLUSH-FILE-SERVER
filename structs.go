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
 Pp_Name  string              `json:pp_name`
 Data     string              `json:data`
}
