package shared

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(filePath string) <-chan string {
	lines := make(chan string)

	go func() {
		defer close(lines)

		f, err := os.Open(filePath)
		if err != nil {
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			v := strings.TrimSpace(scanner.Text())
			if v != "" {
				lines <- v
			}
		}
	}()

	return lines
}
