package api

import (
	"encoding/json"
	"ffserver/ffprobe"
	"log"
	"net/http"
	"os/exec"
)

type GetStreamsDto struct {
	Url string `json:"url"`
}

func GetStreams(w http.ResponseWriter, r *http.Request) {
	var dto GetStreamsDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		RespondWithError(w, http.StatusBadRequest, "bad_body")
		return
	}

	if dto.Url == "" {
		RespondWithError(w, http.StatusBadRequest, "no_url")
		return
	}

	streams, err := ffprobe.GetStreams(dto.Url)
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			// TODO: add ffprobe output
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error_code": "ffprobe",
			})
			return
		}
		log.Println(err)
		RespondWithInternalServerError(w)
		return
	}

	w.Write([]byte(streams))
}
