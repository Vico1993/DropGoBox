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

func getList(path string) string {
	var url = "https://api.dropboxapi.com/2/files/list_folder"
	var data = `{ "path": "` + path + `", "recursive": false, "include_media_info": false, "include_deleted": false, "include_has_explicit_shared_members": false, "include_mounted_folders": true }`
	var result = post(url, accessToken, data, "application/json", false)

	return result
}

func getUpload(filePath string, dropboxFilePath string) string {
	var url = "https://content.dropboxapi.com/2/files/upload"
	var data = `{"path": "` + dropboxFilePath + `","mode": "overwrite","autorename": true,"mute": false}`
	var result = postFile(url, accessToken, data, filePath)

	return result
}

func getDownload(dropboxFilePath string) string {
	var url = "https://content.dropboxapi.com/2/files/download"
	var data = `{ "path": "` + dropboxFilePath + `" }`
	var result = post(url, accessToken, data, "text/plain", true)

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
		dropdata = getList("")
	case "upload":
		dropdata = getUpload("./README.md", "/Homework/test.txt")
	case "download":
		dropdata = getDownload("/convention victor.pdf")
		// dropdata = getDownload("/Homework/data.txt")
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
