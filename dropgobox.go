package main

import (
	"flag"
	"fmt"
)

// Access Token Dropbox
const accessToken = "QO8JziS-WMAAAAAAAAAFr8Lob1Xt0oAWqzrkT6kwjLdbAygnJGiYiuV0VINKmY4G"

// DropElements represents return elements from Dropbox
type DropElements struct {
	Tag         string `json:".tag"`
	Name        string `json:"name"`
	PathLower   string `json:"path_lower"`
	PathDisplay string `json:"path_display"`
	ID          string `json:"id"`
}

// DropList represents return from Dropbox when we use /list.
type DropList struct {
	Elements []DropElements `json:"entries"`
	Cursor   string         `json:"cursor"`
	HasMore  bool           `json:"has_more"`
}

func getList() string {
	var url = "https://api.dropboxapi.com/2/files/list_folder"
	var result = post(url, accessToken, "")

	return result
}

func getUpload() string {
	var url = "https://api.dropboxapi.com/2/files/list_folder"
	var result = post(url, accessToken, "")

	return result
}

func defaultCommand() string {
	var data string
	data = "\n DropGoBox"
	return data
}

func main() {
	// Fonction utilisé lors de l'appelle via CMD
	source := flag.String("method", "default", "")
	flag.Parse()
	fmt.Println(*source)

	var dropdata string
	switch *source {
	case "list":
		dropdata = getList()
	case "upload":
		dropdata = getUpload()
	case "download":
		dropdata = getList()
	default:
		dropdata = defaultCommand()
	}

	println(dropdata)

	// ----> list <----
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
