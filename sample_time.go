package main

import (
  "fmt"
  "time"
)

func main() {
  // 現在日時取得（Local時間）
  t := time.Now()
  fmt.Println(t)

  // 指定日時取得
  local := time.Date(2001, 2, 3, 4, 5, 6, 0, time.Local)
  fmt.Println(local) // -> 2001-02-03 04:05:06 +0900 JST
  utc := time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  fmt.Println(utc)   // -> 2001-02-03 04:05:06 +0000 UTC

  // UTC <-> Local 変換
  t = time.Date(2001, 2, 3, 4, 5, 6, 0, time.Local)
  fmt.Println(t)     // -> 2001-02-03 04:05:06 +0900 JST
  utc = t.UTC()
  fmt.Println(utc)   // -> 2001-02-02 19:05:06 +0000 UTC
  local = utc.Local()
  fmt.Println(local) // -> 2001-02-03 04:05:06 +0900 JST

  // 年、月、日、時、分、秒、ナノ秒 取得
  t = time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  fmt.Println(t.Year())        // -> 2001
  fmt.Println(t.Month())       // -> February
  fmt.Println(int(t.Month()))  // -> 2
  fmt.Println(t.Day())         // -> 3
  fmt.Println(t.Weekday())     // -> Saturday
  fmt.Println(t.YearDay())     // -> 34
  fmt.Println(t.Hour())        // -> 4
  fmt.Println(t.Minute())      // -> 5
  fmt.Println(t.Second())      // -> 6
  fmt.Println(t.Nanosecond())  // -> 0

  // 年月日を一度に取得
  t = time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  year, month, day := t.Date()
  fmt.Println(year)   // -> 2001
  fmt.Println(month)  // -> February
  fmt.Println(day)    // -> 3

  // 時分秒を一度に取得
  t = time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  hour, min, sec := t.Clock()
  fmt.Println(hour) // -> 4
  fmt.Println(min)  // -> 5
  fmt.Println(sec)  // -> 6

  // UNIX時間 秒、ミリ、ナノ秒
  t = time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  fmt.Println(t.Unix())                               // -> 981173106
  fmt.Println(t.UnixNano() / int64(time.Millisecond)) // -> 981173106000
  fmt.Println(t.UnixNano())                           // -> 981173106000000000

  // 日、時、分、秒の切り捨て
  t = time.Date(2001, 2, 3, 4, 5, 6, 789000000, time.UTC)
  fmt.Println(t)                          // -> 2001-02-03 04:05:06.789 +0000 UTC Truncate前
  fmt.Println(t.Truncate(time.Hour * 24)) // -> 2001-02-03 00:00:00 +0000 UTC
  fmt.Println(t.Truncate(time.Hour))      // -> 2001-02-03 04:00:00 +0000 UTC
  fmt.Println(t.Truncate(time.Minute))    // -> 2001-02-03 04:05:00 +0000 UTC
  fmt.Println(t.Truncate(time.Second))    // -> 2001-02-03 04:05:06 +0000 UTC
  t = time.Date(2001, 2, 3, 4, 5, 6, 789000000, time.Local)
  fmt.Println(t.Truncate(time.Hour * 24)) // -> 2001-02-02 09:00:00 +0900 JST 日の切り捨てはUTCで行われるので注意

  // 日時追加（任意の単位で）
  t = time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  t_add := t.Add(time.Minute) // 追加したい時間を指定（マイナスの指定も可能）
  fmt.Println(t_add) // -> 2001-02-03 04:06:06 +0000 UTC

  // 日時追加（年、月、日の指定）
  t = time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  t_add = t.AddDate(1, 2, 3) // 引数は 年、月、日
  fmt.Println(t_add) // -> 2002-04-06 04:05:06 +0000 UTC

  // 前後比較
  t1 := time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  t2 := time.Date(2001, 2, 3, 4, 5, 7, 0, time.UTC)
  fmt.Println(t1.Before(t2)) // -> true  :t1 < t2
  fmt.Println(t2.Before(t1)) // -> false
  fmt.Println(t1.After(t2))  // -> false
  fmt.Println(t2.After(t1))  // -> true  :t2 > t1

  // 同値比較
  t1 = time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  t2 = time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC)
  fmt.Println(t1.Equal(t2)) // -> true :t1 == t2

  // time -> string 変換 アメリカ式時刻 1月2日午後3時4分5秒2006年 を並び変える
  t = time.Date(2001, 2, 3, 14, 5, 6, 0, time.UTC)
  fmt.Println(t.Format("2006/1/2 3:4:5"))       // -> 2001/2/3 2:5:6
  fmt.Println(t.Format("2006/01/02 03:04:05"))  // -> 2001/02/03 02:05:06
  fmt.Println(t.Format("2006/01/02 15:04:05"))  // -> 2001/02/03 14:05:06
  fmt.Println(t.Format("20060102150405"))       // -> 20010203140506
  // 時間 3指定 1-12, 時間 15指定 0-23

  // string -> time 変換 指定するフォーマットはFormatで使用するものと同じ
  t, _ = time.Parse("2006/01/02 15:04:05", "2001/02/03 04:05:06") // 引数1:フォーマット,引数2:パース対象文字列
  fmt.Println(t)  // -> 2001-02-03 04:05:06 +0000 UTC
}
