package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var begins = []string{
	"AVERAGE SPEED ",
	"NOW HEADING ",
	"EN ROUTE ",
	"MY ROUTE ",
	"TRAVEL SPEED ",
	"HEADING ",
	"ETA IN "}

const (
	letters = 'Z' - 'A' + 1
	all     = letters + '9' - '0' + 1
)

func RuneNormalize(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r = r - 'A'
	} else {
		r = r - '0' + letters
	}
	return r
}

func RuneDecode(char rune, dial int) rune {
	char = unicode.ToUpper(char)
	if RuneCheckCodable(char) {
		char = RuneNormalize(char)
		char = char + rune(dial)
		char = char % all
		if char < 0 {
			char += all
		}
		if char < letters {
			char += 'A'
		} else {
			char += '0' - letters
		}
	}
	return char
}

func RuneCheckCodable(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func RuneDistance(r1, r2 rune) (distance int, SafeIgnore bool) {
	r1 = unicode.ToUpper(r1)
	r2 = unicode.ToUpper(r2)
	if RuneCheckCodable(r1) {
		if RuneCheckCodable(r2) {
			r1 = RuneNormalize(r1)
			r2 = RuneNormalize(r2)
			d := r2 - r1
			d = d % all
			if d < 0 {
				d += all
			}
			return int(d), false
		} else {
			return -1, false
		}
	} else {
		if RuneCheckCodable(r2) {
			return -1, false
		} else {
			return 0, true
		}
	}
}

func DecodeLine(line string, dials []int) string {
	numDials := len(dials)
	var ret []rune
	for i, r := range line {
		dial := i % numDials
		d := RuneDecode(r, dials[dial])
		ret = append(ret, d)
	}
	return string(ret)
}

func TryMatchingTests(line string, tests []string, numDials int) (dials []int, success bool) {
	for _, test := range tests {
		dials, success = TryMatching(line, test, numDials)
		if success {
			return
		}
	}
	return nil, false
}

func TryMatching(line, sample string, numDials int) (dials []int, success bool) {
	dials = make([]int, numDials)
	for i := range dials {
		dials[i] = -1
	}
	sampleRunes := []rune(sample)
	for i, r := range line {
		if i >= len(sampleRunes) {
			return dials, true
		}
		s := sampleRunes[i]
		dist, ignore := RuneDistance(r, s)
		if dials[i%numDials] < 0 {
			//first run
			if !ignore {
				if dist < 0 {
					return nil, false
				}
				dials[i%numDials] = dist
			}
		} else {
			//non first run
			if !ignore {
				if dist < 0 {
					return nil, false
				}
				if dials[i%numDials] != dist {
					return nil, false
				}
			}
		}
	}
	return dials, true
}

func TryMatchingTestsRemovals(line string, tests []string, numDials int, maxRemovals int) (dials []int, success bool) {
	removed := 0
	for i := 0; i < maxRemovals; i++ {
		if len(line) < 1 {
			return nil, false
		}
		if i > 0 {
			splits := strings.SplitN(line, " ", 2)
			if len(splits) != 2 {
				return nil, false
			}
			line = splits[1]
			removed += len(splits[0]) + 1
		}
		dials, success = TryMatchingTests(line, tests, numDials)
		if success {
			move := 4 - removed%numDials
			dials = append(dials[move:], dials[:move]...)
			return
		}
	}
	return nil, false
}

func main() {
	var dial1, dial2, dial3, dial4 int
	flag.IntVar(&dial1, "d1", 0, "decoding dial one")
	flag.IntVar(&dial2, "d2", 0, "decoding dial two")
	flag.IntVar(&dial3, "d3", 0, "decoding dial three")
	flag.IntVar(&dial4, "d4", 0, "decoding dial four")
	find := flag.Bool("find", false, "finds and prints dials")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Please supply non option parameter = input filename")
		os.Exit(1)
	}
	fileName := flag.Args()[0]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fullText := strings.Join(lines, " ")
	if *find {
		dials, success := TryMatchingTestsRemovals(fullText, begins, 4, 10)
		if success {
			fmt.Println(dials)
			fmt.Println(fullText)
			fmt.Println(DecodeLine(fullText, dials))
		}
	} else {
		fmt.Println(fullText)
		fmt.Println(DecodeLine(fullText, []int{dial1, dial2, dial3, dial4}))
	}

}
