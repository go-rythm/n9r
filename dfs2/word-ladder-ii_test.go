package dfs2_test

import (
	"testing"

	"github.com/go-rythm/n9r/dfs2"
)

func TestFindLadders(t *testing.T) {
	dict := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
	}
	t.Log(dfs2.FindLadders("a", "c", dict))
}
