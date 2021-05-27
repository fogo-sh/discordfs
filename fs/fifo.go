package fs

import (
	"log"
	"os"
	"syscall"
)

// WriteFifo creates a named pipe at the given path, outputting the given channel to said pipe.
func WriteFifo(path string, c chan []byte) {
	syscall.Mkfifo(path, 0700)

	go func() {
		var file, err = os.OpenFile(path, os.O_WRONLY, 0200)

		if err != nil {
			log.Fatal(err)
		}

		for {
			file.Write(<-c)
		}
	}()
}
