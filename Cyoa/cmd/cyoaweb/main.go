package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/benjamin-daniel/gophercises/Cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the cyoa web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := Cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	h := Cyoa.NewHandler(story, nil)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
