package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// This fizzbuzz uses concurrency to create a "high throughput" fizzbuzz
// application.

// Here is a simplified diagram of what's happening:

//                    +--> fizzer ----+
//                    |               |
// for first..last ---+--> buzzer ----+--> results
//                    |               |
//                    +--> numberer --+

func main() {

	first, last := limitsFromArgs()

	fIn, fOut := piper(fizzer)
	bIn, bOut := piper(buzzer)
	nIn, nOut := piper(numberer)

	feed := tee(fIn, bIn, nIn)
	results := cat(fOut, bOut, nOut)

	go func() {
		for s := range results {
			fmt.Println(s)
		}
	}()

	for i := first; i <= last; i++ {
		feed <- i
	}

	closeAll(fIn, bIn, nIn)
}

// cat reads a set of channels and concatenates the strings from each (in the
// order provided) into a single string. Returns a channel that can be read to
// consume the concatenated results.
func cat(chans ...chan string) chan string {

	out := make(chan string)

	go func() {
		defer close(out)

		for {
			sb := strings.Builder{}
			for _, c := range chans {
				s, ok := <-c
				if !ok {
					return
				}
				sb.WriteString(s)
			}
			out <- sb.String()
		}
	}()

	return out
}

// tee produces a single channel, from which it reads. It will copy each input
// into all of the provided output channels.
func tee(chans ...chan int) chan int {

	in := make(chan int)

	go func() {
		defer close(in)

		for i := range in {
			for _, c := range chans {
				c <- i
			}
		}
	}()

	return in
}

// piper creates a pair of channels around a function that takes an int and
// returns a string, creating a pipeline that reads ints and produces strings.
// The caller must close the returned input channel.
func piper(f func(int) string) (chan int, chan string) {

	in := make(chan int)
	out := make(chan string)

	go func() {
		defer close(out)

		for n := range in {
			out <- f(n)
		}
	}()

	return in, out
}

// fizzer checks if an input int is fizzable. Returns either "fizz" or "".
func fizzer(n int) string {
	if n%3 == 0 {
		return "fizz"
	}

	return ""
}

// buzzer checks if an input int is buzzable. Returns either "buzz" or "".
func buzzer(n int) string {
	if n%5 == 0 {
		return "buzz"
	}

	return ""
}

// numberer checks if an input int is neither fizzable, nor buzzable. Returns
// either the int as a string, or "".
func numberer(n int) string {
	if n%3 == 0 || n%5 == 0 {
		return ""
	}

	return strconv.Itoa(n)
}

// limitsFromArgs parses the command line args to get first, last numbers.
func limitsFromArgs() (int, int) {
	if len(os.Args) != 3 {
		log.Fatalf("usage example: fizzbuzz 1 20")
	}

	first, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("failed to convert input: %v", err)
	}

	last, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("failed to convert input: %v", err)
	}

	if first > last {
		return last, first
	}

	return first, last
}

// closeAll closes all the input channels.
func closeAll(chans ...chan int) {
	for _, c := range chans {
		close(c)
	}
}
