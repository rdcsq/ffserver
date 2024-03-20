package api

import (
	"encoding/json"
	"ffserver/env"
	"ffserver/ffmpeg"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type GenerateThumbnailDto struct {
	Url    string `json:"url"`
	Format string `json:"format"`
}

func GenerateThumbnail(w http.ResponseWriter, r *http.Request) {
	var dto GenerateThumbnailDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		RespondWithError(w, 400, "bad_body")
		return
	}

	if dto.Url == "" {
		RespondWithError(w, 400, "no_url")
		return
	}

	if dto.Format != "webp" && dto.Format != "png" && dto.Format != "jpeg" && dto.Format != "jpg" {
		RespondWithError(w, 400, "unsupported_format")
		return
	}

	id, err := ffmpeg.GenerateThumbnail(dto.Url, dto.Format)
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			// TODO: add ffmpeg output
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error_code": "ffmpeg",
			})
			return
		}
		log.Println(err)
		RespondWithInternalServerError(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"id":          id,
		"downloadUrl": fmt.Sprintf("%v/thumbnail/%v.%v", env.Domain, id, dto.Format),
	})
}

func GetThumbnail(w http.ResponseWriter, r *http.Request) {
	// /thumbnail/
	file := filepath.Join(os.TempDir(), fmt.Sprintf("thumbnail_%v", r.URL.Path[11:]))

	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, file)
}
