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
 Pp_Name  string                `json:pp_name`
 Data     string                `json:data`
}

type Post_Image struct {
 Image_Name  string                `json:post_image_name`
 Data             string                `json:data`
}

type MultiBlob struct {
 Pp_Names []string              `json:pp_names`
 Data     []string              `json:data`
}
