package ffmpeg

import (
	"log"
	"os/exec"
	"strings"
)

func GetVersion() string {
	out, err := exec.Command("ffmpeg", "-version").Output()
	if err != nil {
		log.Fatalln(err)
	}
	output := string(out)
	end_index := strings.Index(output, "-")
	start_index := strings.LastIndex(output[:end_index], " ")
	return output[start_index+1 : end_index]
}
