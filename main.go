package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultBanner  = "shadow.txt"
	charHeight     = 8
	startAscii     = 32
	printableCount = 95
)

func loadBanner(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	glyphs := make([][]string, 0, printableCount)

	i := 0
	for len(glyphs) < printableCount && i < len(lines) {
		//skip any empty leading lines in glyphs
		if lines[i] == "" {
			i++
			continue
		}
		//ensure there are enough lines remaining for one glyph
		if i+charHeight > len(lines) {
			return nil, errors.New("Banner file malformed, unexpected end while reading glyph")
		}

		glyph := make([]string, charHeight)
		for h := 0; h < charHeight; h++ {
			glyph[h] = lines[i+h]
		}
		//after a glyph there is a normal blank separotor line, skip if present
		if i < len(lines) && lines[i] == "" {
			i++
		}
	}
	if len(glyphs) < printableCount {
		return nil, fmt.Errorf("Banner file malformed: found %d glyphs, expected %d", len(glyphs), printableCount)
	}
	return glyphs, nil
}

// sanitizeInput, joins CLI args, removes a singke surroundding pair of braces or quotes,
// and converts literal "\n" into real newline
func sanitizeInput(args []string) string {
	row := strings.Join(args, " ")

	//remove a single pair of matching surrounding braces or quotes, if present
	if len(row) >= 2 {
		first := row[0]
		last := row[len(row)-1]

		if (first == '{' && last == '}') || (first == '"' && last == '"') || (first == '\'' && last == '\'') {
			row = row[1 : len(row)-1]

		}
	}

	//convert "\n" into actual new line
	row = strings.ReplaceAll(row, `\n`, "\n")
	return row
}

func glyphForRune(glyphs [][]string, r rune) []string {
	if r < startAscii && r >= startAscii+printableCount {
		return glyphs[0]
	}
	return glyphs[int(r)-startAscii]

}

func renderLine(glyphs [][]string, line string) {
	rows := make([]strings.Builder, charHeight)
	for _, r := range line {
		g := glyphForRune(glyphs, r)
		for h := 0; h < charHeight; h++ {
			rows[h].WriteString(g[h])
		}
	}
	//print the assembled rows
	for h := 0; h < charHeight; h++ {
		fmt.Println(rows[h].String())
	}
}

func main() {
	if len(os.Args) < 2 {
		// per spec, if no argument provided, print usage to stderr and exit with non-zero code
		fmt.Fprintln(os.Stderr, "usage: go run . \"{Hello There}\"")
		os.Exit(1)
	}

	// Optional: second arg can specify which banner file to use
	// Usage possibilities:
	//  go run . "{Hello}"           -> uses default 'standard' file in current dir
	//  go run . "{Hello}" shadow    -> uses 'shadow' file in current dir
	args := os.Args[1:]
	bannerFile := defaultBanner
	if len(args) >= 2 {
		// if last argument matches a banner file name (shadow/standard/thinkertoy) we treat it as banner name
		possibleBanner := args[len(args)-1]
		// If user explicitly provided a banner file name, use it and remove from text args.
		// We consider it provided only if it exactly matches a file in current dir.
		if _, err := os.Stat(possibleBanner); err == nil {
			bannerFile = possibleBanner
			args = args[:len(args)-1]
		}
	}

	absBannerPath, _ := filepath.Abs(bannerFile)
	glyphs, err := loadBanner(absBannerPath)
	if err != nil {
		// if banner not found in provided path, try current directory file name
		glyphs, err = loadBanner(bannerFile)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load banner file %q: %v\n", bannerFile, err)
		os.Exit(1)
	}

	input := sanitizeInput(args)

	// If the input is empty string, per examples we output nothing.
	if input == "" {
		return
	}

	// Split the input into lines by real newline characters
	lines := strings.Split(input, "\n")

	for idx, ln := range lines {
		// empty input line -> print a single blank line
		if ln == "" {
			fmt.Println()
			continue
		}
		// render the ascii art for this line
		renderLine(glyphs, ln)
		// After a non-empty rendered line, print two blank lines (matches examples)
		fmt.Println()
		fmt.Println()

		// (not required, but keeps behaviour identical for last line)
		_ = idx
	}
}
