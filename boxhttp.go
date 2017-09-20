package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func get(url string, accessToken string) string {
	fmt.Println("---- Récupération  ----")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer : "+accessToken)

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("erreur ReadAll: ", err)
	}
	bodyString := string(body)

	return bodyString
}

func post(url string, accessToken string, dataString string, contentType string, header bool) string {
	fmt.Println("---- POST REQUEST  ----")

	var tmp = dataString
	if header {
		tmp = ""
	} else {
		tmp = dataString
	}
	data := bytes.NewBuffer([]byte(tmp))

	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", contentType)
	if header {
		req.Header.Add("Dropbox-API-Arg", dataString)
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// RECOIT UN FICHIER
	var contentTypeResp = resp.Header.Get("Content-Type")
	if contentTypeResp == "application/octet-stream" {

		var newFile *os.File

		if newFile, err = os.Create("./download/test.txt"); err != nil {
			log.Fatal("error Create : ", err)
		}
		defer newFile.Close()

		if _, err = io.Copy(newFile, resp.Body); err != nil {
			// os.Remove(dst)
		}

		return "DEBUG"
	}

	// ELSE CAS CLASSIC
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("erreur ReadAll: ", err)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
	}

	return string(prettyJSON.Bytes())
}

func postFile(url string, accessToken string, dataString string, path string) string {
	fmt.Println("---- POST FILE REQUEST  ----")

	// Prepare files
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(" --> Open: ", err)
	}
	defer file.Close()

	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		log.Fatal(" --> NewRequest: ", err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Dropbox-API-Arg", dataString)

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(" --> Do: ", err)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(" --> erreur ReadAll: ", err)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Println(" --> JSON parse error: ", error)
	}

	return string(prettyJSON.Bytes())
}
