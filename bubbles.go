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
	// special case for 0. For some reason, it is at the very end
	// of the circled latin small letter unicode block.
	// We can't subtract from 'a' because 0 comes before 'a', it would produce a negative number
	// We can't Abs value it, the code will be 25 codes off,
	// subtracting '0' from 'J' (capital J) works to produce it's code (0x24EA/9450)
	// but it will be more efficient to just special case this.
	if r == '0' {
		return rune(0x24EA)
	} else if unicode.IsLetter(r) {
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
