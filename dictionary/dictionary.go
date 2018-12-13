package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dictionary struct {
	random  []string
	pattern map[string]string
}

var instance = newDictionary()

func GetInstance() *Dictionary {
	return instance
}

func newDictionary() *Dictionary {
	d := new(Dictionary)

	{
		file, err := os.Open("dics/random.txt")
		if err != nil {
			fmt.Println("failed to open dics/random.txt")
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("failed to scan dics/random.txt")
				os.Exit(1)
			}
			line := scanner.Text()
			line = strings.TrimRight(line, "\n")
			d.random = append(d.random, line)
		}
	}

	{
		file, err := os.Open("dics/pattern.txt")
		if err != nil {
			fmt.Println("failed to open dics/pattern.txt")
			os.Exit(1)
		}
		defer file.Close()

		d.pattern = map[string]string{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("failed to scan dics/pattern.txt")
				os.Exit(1)
			}
			line := scanner.Text()
			line = strings.TrimRight(line, "\n")
			if line == "" {
				break
			}
			result := strings.SplitN(line, "\t", 2)
			if len(result) != 2 {
				fmt.Println("failed to split")
				os.Exit(1)
			}
			pattern := result[0]
			phrase := result[1]
			d.pattern[pattern] = phrase
		}
	}
	return d
}

func (d Dictionary) Random() []string {
	return d.random
}

func (d Dictionary) Pattern() map[string]string {
	return d.pattern
}
