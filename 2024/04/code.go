package main

import (
  "aoc-in-go/utils"
	"bufio"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func parseInput(input string) utils.TwoDArray {
  board := [][]rune{}
  r := strings.NewReader(input)
  scanner := bufio.NewScanner(r)

  for scanner.Scan() {
    line := scanner.Text()
    board = append(board, []rune(line))
  }
  arr := utils.TwoDArray{ Board: board }
  return arr
}

func run(part2 bool, input string) any {
  xmas_count := 0
  board := parseInput(input)
  
	if part2 {
    for i := 1; i < len(board.Board) - 1; i++ {
      for j := 1; j < len(board.Board[i]) - 1; j++ {
        if board.Board[i][j] == 'A' {
          mas_count := 0
          // Forward Diagonal 
          if (board.Board[i-1][j-1] == 'M' && board.Board[i+1][j+1] == 'S') ||
          (board.Board[i-1][j-1] == 'S' && board.Board[i+1][j+1] == 'M') {
            mas_count++
          }

          // Backwards diagonal
          if (board.Board[i+1][j-1] == 'M' && board.Board[i-1][j+1] == 'S') ||
          (board.Board[i+1][j-1] == 'S' && board.Board[i-1][j+1] == 'M') {
            mas_count++
          }

          if(mas_count == 2) {
            xmas_count++;
          }
        }
      }
    }
    return xmas_count;
	}

  for i, row := range(board.Board) {
    for j := range(row) {
      for k := 0; k < 8; k++ {
        if board.FindWord("XMAS", i, j, utils.Direction(k)) {
          xmas_count++;
        }
      }
    }
  }

  return xmas_count
}
