package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/katsuyuki-nakamura/unmo/unmo"
)

func prompt(unmo *unmo.Unmo) string {
	return unmo.Name() + "," + unmo.ResponderName() + "> "
}

func main() {
	fmt.Println("Unmo System prototype : proto")
	proto := unmo.NewUnmo("proto")

	for {
		fmt.Print("> ")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		input := stdin.Text()
		input = strings.TrimRight(input, "\n")
		if input == "" {
			break
		}
		response := proto.Dialogue(input)
		fmt.Println(prompt(proto) + response)
	}
}
