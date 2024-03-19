package main

import (
	"ffserver/api"
	"ffserver/env"
	"ffserver/ffmpeg"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting")
	env.LoadEnv()
	log.Printf("ffmpeg version: %v\n", ffmpeg.GetVersion())
	log.Printf("Starting server in %v\n", env.ListeningAddress)

	heartbeat := http.HandlerFunc(api.Heartbeat)
	getStreams := http.HandlerFunc(api.GetStreams)

	http.Handle("GET /", middleware(heartbeat))
	http.Handle("POST /streams", middleware(getStreams))
	http.ListenAndServe(env.ListeningAddress, nil)
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		handler.ServeHTTP(w, r)
	})
}
