package ffmpeg

import (
	"ffserver/utils"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func GenerateThumbnail(source string, format string) (string, error) {
	id := utils.RandomString(32)

	output := filepath.Join(os.TempDir(), "thumbnail_"+id+"."+format)
	if err := exec.Command("ffmpeg", "-loglevel", "error", "-i", source, "-vframes", "1", output).Run(); err != nil {
		return "", err
	}

	// delete the image after 60 seconds
	go func() {
		time.Sleep(60 * time.Second)
		os.Remove(output)
	}()

	return id, nil
}
