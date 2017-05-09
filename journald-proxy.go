package main

import (
	"bufio"
	"github.com/coreos/go-systemd/journal"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var line string
	var err error
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		journal.Send(line, journal.PriNotice, nil)
	}
}
