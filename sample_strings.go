package main

import (
  "fmt"
  "strings"
  "strconv"
)

func main() {
  // 文字列の結合
  s := "ABCDEFG"
  s = s + "hijklmn"
  fmt.Println(s)
  // -> "ABCDEFhijklmn"

  // 文字列の 大文字、小文字変換
  s = "ABCDEFGhijklmn"
  fmt.Println(strings.ToUpper(s))  // -> "ABCDEFHIJKLMN"
  fmt.Println(strings.ToLower(s))  // -> "abcdefhijklmn"

  // 文字列の部分取得
  s = "ABCDEFGhijklmn"
  fmt.Println(s[4:10]) // -> "EFGhij"
  fmt.Println(s[4:])   // -> "EFGhijklmn"
  fmt.Println(s[:10])  // -> "ABCDEFGhij"
  fmt.Println(s[:])    // -> "ABCDEFGhijklmn"

  // 両端トリム
  s = "    123456    "
  fmt.Printf("[%s]\n", strings.TrimSpace(s)) // -> "[123456]"

  // 両端トリム（トリム対象指定）
  s = "    123456    "
  fmt.Printf("[%s]\n", strings.Trim(s, " ")) // -> "[123456]"

  // 左側トリム、右側トリム
  s = "    123456    "
  fmt.Printf("[%s]\n", strings.TrimLeft(s, " "))  // -> "[123456    ]"
  fmt.Printf("[%s]\n", strings.TrimRight(s, " ")) // -> "[    123456]"

  // 文字列中に指定文字列が含まれているか
  s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"
  fmt.Println(strings.Contains(s, "Delta")) // -> true

  // 文字列中の指定文字列の出現位置
  s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"
  fmt.Println(strings.Index(s, "Delta")) // -> 19
  fmt.Println(strings.Index(s, "Hotel")) // -> -1

  // 文字列中の指定文字列の最後の出現位置
  s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"
  fmt.Println(strings.LastIndex(s, "a")) // -> 23

  // 文字列中の指定文字列のカウント
  s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"
  fmt.Println(strings.Count(s, "a")) // -> 4

  // 文字列の先頭一致
  s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"
  fmt.Println(strings.HasPrefix(s, "Alfa")) // -> true

  // 文字列の先頭一致
  s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"
  fmt.Println(strings.HasSuffix(s, "Golf")) // -> true

  // 文字列置換
  s = "hogehogehogehoge"
  fmt.Println(strings.Replace(s, "ge", "ji", 2)) // -> "hojihojihogehoge" 最後の引数は回数
  fmt.Println(strings.Replace(s, "ge", "ji", 0)) // -> "hojihojihojihoji" ０で無制限

  // 文字列繰り返し
  s = "abc"
  fmt.Println(strings.Repeat(s, 3)) // -> "abcabcabc"

  // 文字列のタイトルケースへの変換
  s = "alfa bravo charlie delta echo foxtrot golf"
  fmt.Println(strings.Title(s)) // -> "Alfa Bravo Charlie Delta Echo Foxtrot Golf"

  // 文字列の分割
  s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"
  slice := strings.Split(s, " ")
  for _, str := range slice {
    fmt.Printf("[%s]", str) // -> [Alfa][Bravo][Charlie][Delta][Echo][Foxtrot][Golf]
  }
  fmt.Println("")

  // 文字列結合
  slice = []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf"}
  fmt.Println(strings.Join(slice, ",")) // -> "Alfa,Bravo,Charlie,Delta,Echo,Foxtrot,Golf"

  // 文字列長カウント
  s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"
  fmt.Println(len(s)) // -> 42

  // 文字列長カウント（マルチバイト）
  s = "Go言語"
  fmt.Println(len(s))         // -> 8
  fmt.Println(len([]rune(s))) // -> 4

    // 文字列の整数変換
    s = "123"
    si, _ := strconv.Atoi(s)
    si = si + 1
    fmt.Println(si) // -> 124

    // 文字列の浮動小数変換
    s = "123.456"
    sf, _ := strconv.ParseFloat(s, 64)
    sf = sf + 0.1
    fmt.Println(sf) // -> 123.556

}
