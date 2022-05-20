package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GeneratePreview(videoName string) (string, error) {
	path := "public/"
	suffix := ".jpg"
	input := path + videoName
	pictureName := strings.Split(videoName, ".")[0] + suffix
	output := path + pictureName
	cmdArguments := []string{"-i", input, "-t", "1", "-r", "1", "-q:v", "2", "-f", "image2", output}

	cmd := exec.Command("ffmpeg", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		//log.Fatal(err)
		return "", err
	}
	fmt.Printf("command output: %q", out.String())
	return pictureName, nil
}
