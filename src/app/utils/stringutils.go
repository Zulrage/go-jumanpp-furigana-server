package utils

import (
  "fmt"
  "bufio"
  "os"
  "strings"

  "unicode"
)

func IsOnlyKana(line string) bool {
  for _, c := range line {
      if(!IsKana(string(c))) {
        return false
      }
  }
  return true
}

func StringToLines(s string) []string {
      var lines []string
      scanner := bufio.NewScanner(strings.NewReader(s))
      for scanner.Scan() {
              lines = append(lines, scanner.Text())
      }
      if err := scanner.Err(); err != nil {
          fmt.Fprintln(os.Stderr, "reading standard input:", err)
      }

      return lines
}

func IsChar(s string, rangeTable []*unicode.RangeTable) bool {
	runeForm := []rune(s)
	for _, r := range runeForm {
		if !unicode.IsOneOf(rangeTable, r) {
			return false
		}
	}
	return true
}

func IsKana(s string) bool {
	return !IsKanji(s)
}

func IsKanji(s string) bool {
	return IsChar(s, []*unicode.RangeTable{unicode.Ideographic})
}
