package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		total = 0
		done  bool
	)

	r := bufio.NewReader(os.Stdin)
	for !done {
		l, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			if err == io.EOF {
				done = true
			}
		}
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		n, err2 := strconv.Atoi(l)
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "%q was not an integer, skipping\n", l)
			continue
		}
		total += n
	}
	fmt.Println(total)
}
