package main

import (
	"ffserver/api"
	"ffserver/env"
	"ffserver/ffmpeg"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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

type JwtPayload struct {
	Source string
	jwt.RegisteredClaims
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtString := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
		if jwtString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(jwtString, &JwtPayload{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error decoding JWT")
			}
			return env.AuthSecret, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*JwtPayload)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Printf("auth,%v\n", claims.Source)

		w.Header().Set("Content-Type", "application/json")

		handler.ServeHTTP(w, r)
	})
}
