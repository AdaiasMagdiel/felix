// Command felix — A recursive cat
//
// Felix is a command-line utility written in Go that recursively
// reads files matching wildcard patterns (like *.md, *.txt, **/*.log)
// and prints their contents with the full file path.
//
// Features:
//   • Recursive glob support (**/*.ext)
//   • Skips binary files automatically
//   • Optional output redirection via -o flag
//
// Example:
//   felix -o all.txt **/*.md *.txt
//
// Author: Adaias Magdiel
// License: GPLv3

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

// isBinary checks if a file is likely binary by scanning the first bytes.
func isBinary(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return true
	}
	defer f.Close()

	buf := make([]byte, 800)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return true
	}
	for i := 0; i < n; i++ {
		if buf[i] == 0 {
			return true // contains null byte — likely binary
		}
	}
	return false
}

// printFile prints the file header and contents.
func printFile(w io.Writer, path string) {
	fmt.Fprintf(w, "\n===== %s =====\n", path)

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Fprintln(w, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
	}
}

func main() {
	// Define output flag
	outputPath := flag.String("o", "", "Redirect output to a file instead of stdout")
	flag.Parse()

	patterns := flag.Args()
	if len(patterns) == 0 {
		fmt.Println("Usage: felix [options] <pattern1> [<pattern2> ...]")
		fmt.Println("Example: felix -o all.txt **/*.md *.txt")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Prepare writer
	var writer io.Writer = os.Stdout
	if *outputPath != "" {
		f, err := os.Create(*outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", *outputPath, err)
			os.Exit(1)
		}
		defer f.Close()
		writer = f
	}

	for _, pattern := range patterns {
		matches, err := doublestar.Glob(os.DirFS("."), pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error matching pattern %s: %v\n", pattern, err)
			continue
		}

		for _, m := range matches {
			path := filepath.ToSlash(m)
			info, err := os.Stat(path)
			if err != nil || info.IsDir() {
				continue
			}
			if isBinary(path) {
				fmt.Fprintf(os.Stderr, "Skipping binary file: %s\n", path)
				continue
			}
			printFile(writer, path)
		}
	}
}
