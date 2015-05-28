# go-pipe

Unix-like pipelines for Go

## Description

How to pipe several commands? ...such as `ls ~/Download | grep Vim`. Use this package `go-pipe`!!

## Requirements

- Go

## Usage

1. Run `go get github.com/b4b4r07/go-pipe`

2. Put something like this in your `~/.bashrc` or `~/.zshrc`:

	```go
import "github.com/b4b4r07/go-pipe"
```

### example

```go
package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"

	pipe "github.com/b4b4r07/go-pipe"
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
```

## Installation

	$ go get github.com/b4b4r07/go-pipe

## License

[MIT](https://raw.githubusercontent.com/b4b4r07/dotfiles/master/doc/LICENSE-MIT.txt)

### thanks

[How to pipe several commands? - Stack Overflow](http://stackoverflow.com/questions/10781516/how-to-pipe-several-commands)

## Author

[BABAROT](http://tellme.tokyo) a.k.a. b4b4r07
