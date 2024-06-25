package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// https://duncanleung.com/go-json-encoding-and-decoding/

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func main() {
	resp, err := http.Get("https://ws.audioscrobbler.com/2.0/?method=user.getinfo&user=rj&api_key=" + getEnvVar("LASTFMAPIKEY") + "&format=json")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(resp)
	}
}
