package api

import (
	"encoding/json"
	"ffserver/ffprobe"
	"net/http"
)

type GetStreamsDto struct {
	Url string `json:"url"`
}

func GetStreams(w http.ResponseWriter, r *http.Request) {
	var dto GetStreamsDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		RespondWithError(w, 400, "Bad Request")
		return
	}

	if dto.Url == "" {
		RespondWithError(w, 400, "Empty URL")
		return
	}

	streams, err := ffprobe.GetStreams(dto.Url)
	if err != nil {
		// TODO: properly handle errors
		RespondWithInternalServerError(w)
		return
	}

	w.Write([]byte(streams))
}
