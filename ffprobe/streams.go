package ffprobe

import (
	"log"
	"os/exec"
)

func GetStreams(source string) (string, error) {
	out, err := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_streams", source).Output()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(out), nil
}
