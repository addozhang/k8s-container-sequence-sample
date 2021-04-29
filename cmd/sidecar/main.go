package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const StatusKey = "STATUS"

func main() {
	os.Setenv(StatusKey, "STARTING")
	http.HandleFunc("/ready", checkReady)
	go func() {
		time.Sleep(time.Second * 10)
		os.Setenv(StatusKey, "STARTED")
		log.Println("status: " + os.Getenv(StatusKey))
	}()
	log.Println("status: " + os.Getenv(StatusKey))
	http.ListenAndServe(":8080", nil)
}

func checkReady(w http.ResponseWriter, req *http.Request) {
	if status := os.Getenv(StatusKey); status == "STARTED" {
			w.WriteHeader(200)
			fmt.Fprintf(w, "success")
			return
	}
	w.WriteHeader(503)
	fmt.Fprintf(w, "fail")
}

