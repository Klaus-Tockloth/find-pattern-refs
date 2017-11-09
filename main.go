/*
Purpose:
- find pattern references in text files

Description:
- Define a pattern set (text or regex) and find all occurrences within a file set.

Releases:
- 1.0.0 - 2017/11/08 : initial release

Author:
- Klaus Tockloth

Copyright and license:
- Copyright (c) 2017 Klaus Tockloth
- MIT license

Permission is hereby granted, free of charge, to any person obtaining a copy of this software
and associated documentation files (the Software), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or
substantial portions of the Software.

The software is provided 'as is', without warranty of any kind, express or implied, including
but not limited to the warranties of merchantability, fitness for a particular purpose and
noninfringement. In no event shall the authors or copyright holders be liable for any claim,
damages or other liability, whether in an action of contract, tort or otherwise, arising from,
out of or in connection with the software or the use or other dealings in the software.

Contact (eMail):
- freizeitkarte@googlemail.com

Remarks:
- NN

Links:
- NN
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// general program info
var (
	progName    = os.Args[0]
	progVersion = "1.0.0"
	progDate    = "2017/11/08"
	progPurpose = "find pattern references in text files"
	progInfo    = "Define a pattern set (text or regex) and find all occurrences within a file set."
)

// Pair : a data structure to hold a key/value pair
type Pair struct {
	Key   string
	Value int
}

// PairList : a slice of Pairs that implements sort.Interface to sort by Value
type PairList []Pair

// command line flags
var regexFlag = flag.Bool("regex", false, "interpret pattern as regular expression (default = false)")

/*
main starts this program
*/
func main() {

	fmt.Printf("\nProgram:\n")
	fmt.Printf("  Name    : %s\n", progName)
	fmt.Printf("  Release : %s - %s\n", progVersion, progDate)
	fmt.Printf("  Purpose : %s\n", progPurpose)
	fmt.Printf("  Info    : %s\n", progInfo)
	fmt.Printf("\n")

	flag.Usage = printUsage
	flag.Parse()

	// verify that sufficient file arguments have been provided
	if len(flag.Args()) != 2 {
		fmt.Printf("Error:\n  Number of arguments insufficient.\n")
		printUsage()
	}

	// read in file list
	filelist, err := slurpFile(flag.Arg(0))
	if err != nil {
		fmt.Printf("error <%v> at slurpFile(); file = <%v>\n", err, flag.Arg(0))
		os.Exit(1)
	}

	// read in pattern list
	patternlist, err := slurpFile(flag.Arg(1))
	if err != nil {
		fmt.Printf("error <%v> at slurpFile(); file = <%v>\n", err, flag.Arg(1))
		os.Exit(1)
	}

	fmt.Printf("Start options and arguments:\n")
	fmt.Printf("  -regex      : %t\n", *regexFlag)
	fmt.Printf("  filelist    : %s\n", flag.Arg(0))
	fmt.Printf("  patternlist : %s\n", flag.Arg(1))
	fmt.Printf("\n")

	verifyReferences(filelist, patternlist)

	fmt.Printf("\n")
}

/*
printUsage prints the usage of this program
*/
func printUsage() {

	fmt.Printf("\nUsage:\n")
	fmt.Printf("  %s [-regex] filelist patternlist\n", os.Args[0])

	fmt.Printf("\nExample:\n")
	fmt.Printf("  %s files.dat pattern.dat\n", os.Args[0])

	fmt.Printf("\nArguments:\n")
	fmt.Printf("  filelist\n")
	fmt.Printf("        list with all (text) files to process (one file per line)\n")
	fmt.Printf("  patternlist\n")
	fmt.Printf("        list with all pattern to reference (one pattern per line)\n")

	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	fmt.Printf("\n")

	os.Exit(1)
}

/*
verifyReferences verifies if text references are found in the file
*/
func verifyReferences(filelist []string, patternlist []string) {

	// build reference map
	refmap := make(map[string]int)
	for _, text := range patternlist {
		refmap[text] = 0
	}

	// build regex list from patternlist (compile only once)
	var regexlist []*regexp.Regexp
	if *regexFlag {
		for _, text := range patternlist {
			regexlist = append(regexlist, regexp.MustCompile(text))
		}
	}

	// for all files
	for _, file := range filelist {
		fmt.Printf("--> processing file %s ...\n", file)
		lines, err := slurpFile(file)
		if err != nil {
			fmt.Printf("error <%v> at slurpFile(); file = <%v>\n", err, file)
			continue
		}
		// for all texts
		for i, text := range patternlist {
			fmt.Printf("%s\n", text)
			for linenumber, line := range lines {
				found := false
				if *regexFlag {
					regexFound := regexlist[i].MatchString(line)
					if regexFound {
						found = true
					}
				} else {
					stringFound := strings.Index(line, text)
					if stringFound != -1 {
						found = true
					}
				}
				if found {
					fmt.Printf("    %d:%s\n", (linenumber + 1), line)
					refmap[text]++
				}
			}
		}
	}

	// print reference map (sort by value)
	pl := make([]Pair, len(refmap))
	i := 0
	for key, value := range refmap {
		pl[i] = Pair{key, value}
		i++
	}
	sort.Slice(pl, func(i, j int) bool { return pl[i].Value < pl[j].Value })
	if *regexFlag {
		fmt.Printf("\nlines found : pattern (interpreted as regex)\n")
	} else {
		fmt.Printf("\nlines found : pattern (interpreted as text)\n")
	}
	for _, kvPair := range pl {
		fmt.Printf("%v : %v\n", refmap[kvPair.Key], kvPair.Key)
	}
}

/*
slurpFile slurps all lines of a text file into a slice of strings
*/
func slurpFile(filename string) ([]string, error) {

	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}
