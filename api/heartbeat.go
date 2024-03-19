package api

import (
	"encoding/json"
	"ffserver/ffmpeg"
	"net/http"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"serverVersion": "dev",
		"ffmpegVersion": ffmpeg.GetVersion(),
	})
}
