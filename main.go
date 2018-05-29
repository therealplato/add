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
	_ = total

	r := bufio.NewReader(os.Stdin)
	for {
		l, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			fmt.Fprintf(os.Stdout, "%q was not an integer, skipping\n", l)
			continue
		}
		total += n
	}
	fmt.Println(total)
}
