package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	tests := []struct {
		name    string
		row     int
		col     int
		player  int
		wantErr bool
	}{
		{"valid move", 0, 0, X, false},
		{"invalid position - negative row", -1, 0, X, true},
		{"invalid position - negative column", 0, -1, X, true},
		{"invalid position - row out of range", 3, 0, X, true},
		{"invalid position - column out of range", 0, 3, X, true},
		{"invalid player", 0, 0, 3, true},
		{"position already occupied", 0, 0, O, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var board Board
			if tt.name == "position already occupied" {
				board.Set(0, 0, X)
			}
			err := board.Set(tt.row, tt.col, tt.player)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckWin(t *testing.T) {
	tests := []struct {
		name string
		board Board
		want int
	}{
		{
			"win by first row",
			Board{
				{X, X, X},
				{O, empty, O},
				{empty, O, empty},
			},
			X,
		},
		{
			"win by first column",
			Board{
				{O, X, X},
				{O, empty, O},
				{O, O, empty},
			},
			O,
		},
		{
			"win by diagonal",
			Board{
				{X, O, X},
				{O, X, O},
				{empty, O, X},
			},
			X,
		},
		{
			"no win",
			Board{
				{X, O, X},
				{O, X, O},
				{O, X, O},
			},
			empty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.CheckWin(); got != tt.want {
				t.Errorf("CheckWin() = %v, want %v", got, tt.want)
			}
		})
	}
}