package dfs2_test

import (
	"testing"

	"github.com/go-rythm/n9r/dfs2"
)

func TestWordSearchII(t *testing.T) {
	board := [][]byte{
		[]byte("doaf"),
		[]byte("agai"),
		[]byte("dcan"),
	}
	words := []string{"dog", "dad", "dgdg", "can", "again"}
	t.Log(dfs2.WordSearchII(board, words))
}
