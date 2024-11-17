package main

import (
  "fmt"
  "strings"
  "unicode"
  //"strconv"
)

func main() {
    fmt.Println("Hello, World!")
    fen := "8/5k2/3p4/1p1Pp2p/pP2Pp1P/P4P1K/8/8 b - - 99 50"
    game := Game{fen}
    game.Draw()
  }


//♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜
//♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟
//… … … … … … … …
//… … … … … … … …
//… … … … … … … …
//… … … … … … … …
//♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙
//♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖


func fen_to_board() {
  fmt.Println("")
}

type Game struct {
   Fen string // https://www.chess.com/terms/fen-chess
}

func (g Game) Draw() float64 {
  board := strings.Split(g.Fen, " ")
  lines := strings.Split(board[0], "/")

  var divider strings.Builder
  divider.WriteString("   ")
  divider.WriteString(strings.Repeat("+---", 8))
  divider.WriteString("+")

  charmap := map[string]string {
    "R":"♜",
    "N":"♞",
    "B":"♝",
    "Q":"♛",
    "K":"♚",
    "P":"♟",

    "r":"♖",
    "n":"♘",
    "b":"♗",
    "q":"♕",
    "k":"♔",
    "p":"♙",
  }


  for index, line := range lines {

    
    fmt.Println(divider.String())
  

    var real_line strings.Builder
    for _, char := range line {
        //fmt.Println(string(char))


        if unicode.IsDigit(char) {
          //fmt.Println(char, int(char), string(char), int(char)-'0')
          real_line.WriteString(strings.Repeat("|   ", int(char)-'0'))
          
          //fmt.Println(strings.Repeat("x", int(char)))
        } else {
          str_char := string(char)

          real_line.WriteString("| ")
          if real_char, ok := charmap[str_char]; ok {
            real_line.WriteString(real_char)
          }

          real_line.WriteString(" ")
        }
    }

    real_line.WriteString("|")
    fmt.Println("", index + 1, real_line.String())



  }
  fmt.Println(divider.String())
  fmt.Println("     a   b   c   d   e   f   g   h")
  return 1
}


//func isNumeric(s rune) bool {
//    _, err := strconv.ParseFloat(s, 64)
//    return err == nil
//}




