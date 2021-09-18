package main

import (
	"fmt"
	"log"
	"os"

	"github.com/perryd01/unaswrappergo"
)

var apiKey string

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	apiKey = os.Getenv(apiKey)
}

func main() {
	uo, err := unaswrappergo.AuthwithAPIKey(apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(uo.Login.Token)
}
