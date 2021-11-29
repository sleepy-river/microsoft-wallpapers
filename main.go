package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	INPUT_FILE            = "wallpapers.txt"
	OUTPUT_FILE_EXTENSION = ".png"
)

func main() {
	data, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		lineWords := strings.Split(line, ". ")
		if len(lineWords) > 1 {
			url := strings.TrimSpace(strings.Split(line, ". ")[1])
			err = DownloadFile(url)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Print("\033[u\033[K") // restore the cursor position and clear the line
				fmt.Printf("Downloading: %s", progressBar(i, len(lines)))
			}
		}
	}
}

func DownloadFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	urlParts := strings.Split(url, "/")
	subdirectory := urlParts[2]
	if _, err := os.Stat(subdirectory); os.IsNotExist(err) {
		err := os.Mkdir(subdirectory, 0700)
		if err != nil {
			return err
		}
	}
	filePath := subdirectory + "/" + urlParts[len(urlParts)-1] + OUTPUT_FILE_EXTENSION

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func progressBar(currentValue, maxValue int) string {
	percentage := float64(currentValue) / float64(maxValue)
	maxBarPieces := 20
	currentBarPieces := int(percentage * float64(maxBarPieces))

	var progress string
	for i := 0; i < currentBarPieces; i++ {
		progress += "="
	}
	for i := 0; i < maxBarPieces-currentBarPieces; i++ {
		progress += "-"
	}
	progress += fmt.Sprintf("> (%.4g%%)", percentage*100.0)

	return progress
}
