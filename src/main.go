package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	PROGRAM_NAME = "gocat"
	VERSION      = "1.0"
)

var (
	stderr = log.New(os.Stderr, fmt.Sprintf("%s: ", PROGRAM_NAME), 0)
)

func main() {
	CLI()
	filenames := flag.Args()

	if len(filenames) == 0 {
		Read(os.Stdin)
		os.Exit(0)
	}

	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			stderr.Fatalln(err.Error())
		}
		defer f.Close()

		Read(f)
	}
}

// Flags
var (
	nFlag       *bool
	versionFlag *bool
)

func CLI() {
	nFlag = flag.Bool("n", false, "Number the output lines, starting at 1.")
	versionFlag = flag.Bool("version", false, "Print version information and exit.")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s -- concatenate and print files (version %v)\n", PROGRAM_NAME, VERSION)
		fmt.Println("github: https://github.com/tuladhar/gocat")
		os.Exit(0)
	}
}

func ReadFile(f io.Reader) {
	Read(f)
}

func ReadStdin() {
	Read(os.Stdin)
}

func Read(f interface{}) {
	scanner := bufio.NewScanner(f.(io.Reader))
	var lineNo int
	for scanner.Scan() {
		line := scanner.Text()
		if *nFlag {
			lineNo++
			fmt.Printf("    %v  %s\n", lineNo, line)
		} else {
			fmt.Println(line)
		}
	}
	err := scanner.Err()
	if err != nil {
		stderr.Fatalln(err.Error())
	}
}
