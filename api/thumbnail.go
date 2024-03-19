package api

import (
	"encoding/json"
	"ffserver/env"
	"ffserver/ffmpeg"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type GenerateThumbnailDto struct {
	Url    string `json:"url"`
	Format string `json:"format"`
}

func GenerateThumbnail(w http.ResponseWriter, r *http.Request) {
	var dto GenerateThumbnailDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		RespondWithError(w, 400, "Bad Request")
		return
	}

	if dto.Url == "" {
		RespondWithError(w, 400, "Empty URL")
		return
	}

	if dto.Format != "webp" && dto.Format != "png" && dto.Format != "jpeg" && dto.Format != "jpg" {
		RespondWithError(w, 400, "Bad format")
		return
	}

	id, err := ffmpeg.GenerateThumbnail(dto.Url, dto.Format)
	if err != nil {
		// TODO: properly handle errors
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
