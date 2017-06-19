package main

import (
  "fmt"
  "strings"
  "bytes"

  "utils"
  "command"
)

func main() {
  //input := "けものフレンズ、アイカツとアイカツ、アイカツ、アイカツはＦｕｊｉ＆ｇｕｍｉＧａｍｅｓなので"
  //input := "太郎は花子が読んでいる本を次郎に渡した"
  //input := "【商品】ブルーレイ＆ＤＶＤ「ゆるゆりｖｏｌ．１」（ポニーキャニオン）、【商品】「けものフレンズＢＤ付オフィシャルガイドブック第２巻」、【商品】ブルーレイ＆ＤＶＤ「アイカツスターズ！」（ハピネット）を紹介。"
  input := "【ゲスト】津田美波【資料・協力】なもり／一迅社・七森中ごらく部、セガ／ヒーローバンクプロジェクト、けものフレンズプロジェクトＡ、ＢＮＰ／ＢＡＮＤＡＩ，ＤＥＮＴＳＵ，ＴＶＴＯＫＹＯ、ＢＮＥＩ／ＰＲＯＪＥＣＴＣＩＮＤＥＲＥＬＬＡ、ＢＡＮＤＡＩＮＡＭＣＯＥｎｔｅｒｔａｉｎｍｅｎｔＩｎｃ．、セガ、２０１６コーエーテクモゲームスＡｌｌｒｉｇｈｔｓｒｅｓｅｒｖｅｄ．、Ｆｕｊｉ＆ｇｕｍｉＧａｍｅｓ※複製・転用等を禁止します。"
  //input := "何らかの成長を世界が認めた際に、まれに獲得出来る事があるのが\"スキル\"らしい。\"進化\"など、それこそ普通の人には縁のないものなのだそうだ。"
  //input := "から揚げ万象だと？これは全ての知識が労せず手に入ったのか！？と思ったのだが…実際には、俺が触れた情報に対して、俺の知りえる事柄に対してのみ情報開示が可能との事。"

  lines := command.GetKnpCommand(input)
  var buffer bytes.Buffer

  for _, line := range lines {
    fmt.Println(line)
    if(!utils.IsEndProcess(line)) {
      if(utils.IsAcceptableCharacter(line)) {
        if(!utils.IsAlias(line)) {
          buffer.WriteString(processLine(line))
        }
      } else {
        buffer.WriteString(utils.GetOriginalWriting(line))
      }
    }
  }
  fmt.Println(buffer.String())
}

func processLine(line string) string {
  result := utils.GetOriginalWriting(line)
  if(!utils.IsOnlyKana(result)) {
    var buffer bytes.Buffer
    m := make([]string, 0)
    // For japanese character random access handling
    runeForm := []rune(result)
    var firstState bool = utils.IsKana(string(runeForm[0]))
    // Default to non first character to create in array
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
    furigana := utils.GetFuriganaWriting(line)
    mapState := firstState
    for i, part := range m {
      if(mapState) {
        buffer.WriteString(part)
        if(i != len(m)-1) {
          furigana = strings.Split(furigana, part)[1]
        }
      } else {
        if(i == len(m)-1) {
          buffer.WriteString("<" + part + "/" + furigana + ">")
        } else {
          splittedFurigana := strings.Split(furigana, m[i+1])
          buffer.WriteString("<" + part + "/" + splittedFurigana[0] + ">")
          furigana = strings.Split(furigana, splittedFurigana[0])[1]
        }
      }
      mapState = !mapState
    }
    result = buffer.String()
  }
  return result
}
