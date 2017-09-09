package jumanpp

import (
  "strings"
  "bytes"

  "app/utils"
  "app/command"
)

func ToFuriganaText(input string) string {
  lines := utils.StringToLines(input)
  var buffer bytes.Buffer
  for _, line := range lines {
    buffer.WriteString("\n")
    jumanLines := command.GetKnpCommand(line)
    buffer.WriteString(processText(jumanLines))
  }

  return buffer.String()
}

func processText(lines []string) string {
  var buffer bytes.Buffer

  for _, line := range lines {
    if(!utils.IsEndProcess(line) && !utils.IsAlias(line)) {
      if(utils.IsAcceptableCharacter(line)) {
        buffer.WriteString(processLine(line))
      } else {
        buffer.WriteString(utils.GetOriginalWriting(line))
      }
    }
  }
  return buffer.String()
}

func processLine(line string) string {
  result := utils.GetOriginalWriting(line)
  if(!utils.IsOnlyKana(result)) {
    m, firstState := separateKanaAndKanjis(result)
    result = getFormattedStringFromSeparation(line, m, firstState)
  }
  return result
}

func separateKanaAndKanjis(result string) ([]string, bool) {
  m := make([]string, 0)
  // For japanese character random access handling
  runeForm := []rune(result)
  var firstState bool = utils.IsKana(string(runeForm[0]))
  // Default opposite of first state
  var previousState bool = !firstState
  for _, c := range runeForm {
      state := utils.IsKana(string(c))
      if (previousState != state) {
        previousState = state
        m = append(m, string(c))
      } else {
        m[len(m)-1] += string(c)
      }
  }
  return m, firstState
}

func getFormattedStringFromSeparation(line string,m []string, firstState bool) string {
  var buffer bytes.Buffer
  furigana := utils.GetFuriganaWriting(line)
  mapState := firstState
  for i, part := range m {
    if(mapState) {
      buffer.WriteString(part)
      if(i != len(m)-1) {
        furigana = strings.Split(furigana, part)[1]
      }
    } else {
      buffer.WriteString("<")
      buffer.WriteString(part)
      buffer.WriteString("/")
      if(i == len(m)-1) {
        buffer.WriteString(furigana)
      } else {
        splittedFurigana := strings.Split(furigana, m[i+1])
        buffer.WriteString(splittedFurigana[0])
        furigana = strings.Split(furigana, splittedFurigana[0])[1]
      }
      buffer.WriteString(">")
    }
    mapState = !mapState
  }
  return buffer.String()
}
