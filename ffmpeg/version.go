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
	start_index := strings.Index(output, "version") + 8
	end_index := strings.Index(output[start_index:], " ")
	return output[start_index : start_index+end_index]
}
