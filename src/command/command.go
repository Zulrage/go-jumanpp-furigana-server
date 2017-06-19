package command

import (
	"os/exec"
  "io"
  "bytes"

	"utils"
)

func GetKnpCommand(input string) []string {
  return utils.StringToLines(getKnpString(input))
}

func getKnpString(input string) string {
  c1 := exec.Command("echo", input)
  c2 := exec.Command("jumanpp")

  jumanReader, echoWriter := io.Pipe()
  c1.Stdout = echoWriter
  c2.Stdin = jumanReader

  var b2 bytes.Buffer
  c2.Stdout = &b2

  c1.Start()
  c2.Start()
  c1.Wait()
  echoWriter.Close()
  c2.Wait()
  str := b2.String()

  return string(str)
}
