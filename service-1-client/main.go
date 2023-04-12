package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting client...requesting from http://service-2-server.default.svc.cluster.local:4317")

	for i := 0; i < 1; {
		resp, err := http.Get("http://service-2-server.default.svc.cluster.local:4317")
		if err != nil {
			log.Println(err)
			log.Println("Sleep container indefinitely.")
			select {}
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		strBody := string(body)
		log.Println(strBody)
		time.Sleep(5 * time.Second)
	}
}
