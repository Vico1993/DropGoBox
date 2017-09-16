package main

import (
    "flag"
)

const access_token = "QO8JziS-WMAAAAAAAAAFr8Lob1Xt0oAWqzrkT6kwjLdbAygnJGiYiuV0VINKmY4G"

func main( ) {
    source := flag.String("method", "list", "une list")
	flag.Parse()
    println(*source)
}
