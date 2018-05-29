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
	total := 0

	r := bufio.NewReader(os.Stdin)
	for {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		n, err2 := strconv.Atoi(l)
		if err2 != nil {
			fmt.Fprintf(os.Stdout, "%q was not an integer, skipping\n", l)
			continue
		}
		total += n
		if err == io.EOF {
			break
		}
	}
	fmt.Println(total)
}
