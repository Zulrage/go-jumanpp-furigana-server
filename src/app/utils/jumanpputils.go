package utils

import (
  "strings"
)

func GetOriginalWriting(line string) string {
  return strings.Split(line, " ")[0]
}

func GetFuriganaWriting(line string) string {
  return strings.Split(line, " ")[1]
}

func IsEndProcess(line string) bool {
  return strings.Contains(line, "EOS")
}

func IsAlias(line string) bool {
  return GetOriginalWriting(line) == "@" && GetOriginalWriting(line) != GetFuriganaWriting(line)
}

var notCompletingTypes = []string{"特殊", "助詞"}
func IsAcceptableCharacter(e string) bool {
    for _, a := range notCompletingTypes {
        if strings.Contains(e, a) {
            return false
        }
    }
    return true
}

func IsSpace(e string) bool {
    return strings.HasPrefix(e, "　")
}
