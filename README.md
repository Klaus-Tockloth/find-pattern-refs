## find-pattern-refs
### Find pattern references in text files.
```
Usage:
  ./find-pattern-refs [-regex] filelist patternlist

Example:
  ./find-pattern-refs files.dat pattern.dat

Arguments:
  filelist
        list with all (text) files to process (one file per line)
  patternlist
        list with all pattern to reference (one pattern per line)

Options:
  -regex
    	interpret pattern as regular expression (default = false)
```

Output sample:
```
...
--> processing file freizeitkarte-v5.xml ...
alm.svg
    810:      <area src="file:/patterns/alm.svg" symbol-width="32" />
    813:      <area src="file:/patterns/alm.svg" symbol-width="32" stroke="#c6e4b4" stroke-width="0.1" />
feuchtwiese.svg
    1815:      <area src="file:/patterns/feuchtwiese.svg" symbol-width="32" />
...
lines found : pattern (interpreted as text)
2 : alm.svg
5 : feuchtwiese.svg
...
```

### Use cases:
* Find unused resource files, e.g. unused images (svg, png, ...) within your website files (html, css, ...).
* Find all occurrences of one or more pattern (interpreted as text or regex) within a set of one or more files.

### Hints (linux, macOS):
* Use "ls *.html > files.dat" to create a simple file list.
* Use "ls *.svg *.svg > pattern.dat" to create a simple pattern list.
* If you set the "-regex" option, each pattern is interpreted as regex pattern.
* Without the "-regex" option (that's the default), each pattern is interpreted as text.
* Use the text interpretation if you intend to reference (ressource) file names.

### Binaries:
The prebuild binaries for [Linux, MacOS (darwin), Windows](https://github.com/Klaus-Tockloth/find-pattern-refs/releases/latest) have no dependencies ... just download the utility and run it from the command line.

### Technical information:
The master branch is used for development. A stable release is tagged with a string according to semantic versioning specification (e.g. 2.7.1).
