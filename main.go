package main

import (
  "fmt"
  "strconv"
)

const word = "exit"
var trys int = 7
var found string = "____"

func main() {
  intro()
  playGame()
}

func intro() {
  fmt.Println("welcome to hang man")
}

func playGame() {
  if(trys == 0) {
    fmt.Printf("you lost\n");
    return;
  }
  if(word == found) {
    fmt.Printf("you win\n");
    return;
  }
  fmt.Printf("trys: " + strconv.Itoa(trys) + " found: " + found)
  fmt.Print("\nenter a letter: ")
  var input string
  fmt.Scanln(&input)
  if(valideLetter(input)) {
    trys -= 1;
    matchWord(input)
    playGame()
  } else {
    fmt.Println("that is not a valid letter")
    playGame()
  }
}

func valideLetter(letter string) bool {
  var alphabet string = "abcdefghijklmnoprstuvwxyz";

  for _, l := range alphabet {
    if(letter == string(l)) {
      return true
    }
  }
  return false
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

func matchWord(guess string) {
  for j, c := range word {
    for _, i := range guess {
      if(i == c) {
        found = replaceAtIndex(found, i, j)
      }
    }
  }
}
