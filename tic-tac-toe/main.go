package main

import (
	"fmt"
	"strings"
)

const (
	empty = iota
	X
	O
)

type Board [3][3]int

func (b *Board) String() string {
	var buf strings.Builder
	for _, row := range b {
		for _, cell := range row {
			switch cell {
			case X:
				buf.WriteString("X")
			case O:
				buf.WriteString("O")
			default:
				buf.WriteString(" ")
			}
			buf.WriteString("|")
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (b *Board) Set(row, col, player int) error {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return fmt.Errorf("invalid position")
	}
	if player != X && player != O {
		return fmt.Errorf("invalid player")
	}
	if b[row][col] != empty {
		return fmt.Errorf("position already occupied")
	}
	b[row][col] = player
	return nil
}

func (b *Board) CheckWin() int {
	for _, row := range b {
		if row[0] != empty && row[0] == row[1] && row[0] == row[2] {
			return row[0]
		}
	}
	for col := 0; col < 3; col++ {
		if b[0][col] != empty && b[0][col] == b[1][col] && b[0][col] == b[2][col] {
			return b[0][col]
		}
	}
	if b[0][0] != empty && b[0][0] == b[1][1] && b[0][0] == b[2][2] {
		return b[0][0]
	}
	if b[0][2] != empty && b[0][2] == b[1][1] && b[0][2] == b[2][0] {
		return b[0][2]
	}
	return empty
}

func main() {
	var board Board
	board.Set(0, 0, X)
	board.Set(0, 1, O)
	board.Set(0, 2, X)
	board.Set(1, 0, O)
	board.Set(1, 1, X)
	board.Set(1, 2, O)
	board.Set(2, 0, X)
	fmt.Println(&board)
	switch board.CheckWin() {
	case X:
		fmt.Println("X wins!")
	case O:
		fmt.Println("O wins!")
	default:
		fmt.Println("No one wins!")
	}
}
