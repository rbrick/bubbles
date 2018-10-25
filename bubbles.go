package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

//mapperFunc translates a alphanumeric character into an enclosed "bubble" character and outputs it into stdout or a file.
var mapperFunc = func(r rune) rune {
	if unicode.IsLetter(r) {
		if unicode.IsLower(r) {
			return (r - 'a') + 0x24D0
		} else {
			return (r - 'A') + 0x24B6
		}

	} else if unicode.IsDigit(r) {
		return (r - '1') + 0x2460
	}
	return r

}

var fileFlag = flag.String("f", "", `create a new file with the output "bubble-fied" `)

func main() {

	flag.Parse()

	output := os.Stdout

	if *fileFlag != "" {
		f, err := os.OpenFile(*fileFlag, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err == nil {
			output = f
		}

	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(output, strings.Map(mapperFunc, scanner.Text()))
	}
}
