package client

import (
	"fmt"
	"log"
	"net/http"
)

func GetWorker() {
	resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		log.Fatalf("Error in http.Get: %s", err)
	}
	fmt.Println(resp)
}
