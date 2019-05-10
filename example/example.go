package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/b4b4r07/go-pipe"
)

func main() {
	var b bytes.Buffer
	if err := pipe.Command(&b,
		exec.Command("ls", "/Users/b4b4r07/Downloads"),
		exec.Command("grep", "Vim"),
	); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, &b); err != nil {
		log.Fatal(err)
	}
}
