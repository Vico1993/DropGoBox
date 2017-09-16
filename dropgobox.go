package main

import (
	"flag"
	"fmt"
	"strings"
)

// Access Token Dropbox
const accessToken = "QO8JziS-WMAAAAAAAAAFr8Lob1Xt0oAWqzrkT6kwjLdbAygnJGiYiuV0VINKmY4G"

func getList() {
	fmt.Println("---- Récupération  ----\n")

	var url = "https://api.dropboxapi.com/2/files/list_folder"
	var result = post(url, accessToken)

	println(strings.NewReader(result))
}

func main() {
	source := flag.String("method", "list", "une list")
	flag.Parse()
	fmt.Println(*source)

	if *source == "list" {
		getList()
	}

	// list
	// upload
	// download
	// delete
	// move
	// copy
	// mkdir
	// monitor
	// share
	// saveurl
	// search
	// info
	// space
	// unlink
}
