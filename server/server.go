package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	mw "server/middleware"
	helper "server/services/helpers"
	

	vision "cloud.google.com/go/vision/apiv1"
)

var PORT string

func main() {
	fmt.Println("Server is running...")

	http.HandleFunc("/", mw.CORS(handleLinkSearch))

	PORT = os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// Golang server
	fmt.Printf("API Server running at port " + PORT + "/\n")
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		fmt.Println("(server.go) Golang server has stopped due to:")
		log.Fatal(err)
	}
}

type recievedData struct {
	PhotoURL string
}

type returnData struct {
	Text               string
	FullMatchingLinks  []string
	ImageMatchingLinks []string
}

func handleLinkSearch(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var recievedData recievedData
		json.Unmarshal([]byte(reqBody), &recievedData)

		ctx := context.Background()
		client, err := vision.NewImageAnnotatorClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		defer client.Close()

		annotations := helper.DetectWeb(ctx, client, recievedData.PhotoURL)
		if annotations == nil {
			fmt.Println("Error")
		}

		fullMatchingLinks, imageMatchingLinks := helper.PrintResults(annotations)
		var returnData returnData
		if fullMatchingLinks == nil {
			if imageMatchingLinks == nil {
				returnData.Text = "No results found"
			} else {
				returnData.ImageMatchingLinks = imageMatchingLinks
				returnData.Text = "Partial results found"
			}
		} else {
			returnData.Text = "Results found!"
			returnData.FullMatchingLinks = fullMatchingLinks
		}

		returnRequestData, _ := json.Marshal(returnData)
		w.Write(returnRequestData)
	default:
		var returnData returnData
		returnData.Text = "Request error!"
		returnRequestData, _ := json.Marshal(returnData)
		w.WriteHeader(http.StatusNotFound)
		w.Write(returnRequestData)
	}
}