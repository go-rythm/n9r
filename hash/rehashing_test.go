package hash_test

import (
	"testing"

	"github.com/go-rythm/n9r/hash"
	jsoniter "github.com/json-iterator/go"
)

func TestRehashing(t *testing.T) {
	hashTable := []*hash.ListNode{
		nil,
		{
			Val:  21,
			Next: &hash.ListNode{Val: 9},
		},
		{
			Val: 14,
		},
		nil,
	}

	newTable := hash.Rehashing(hashTable)

	got, _ := jsoniter.MarshalToString(newTable)
	want := `[null,{"Val":9,"Next":null},null,null,null,{"Val":21,"Next":null},{"Val":14,"Next":null},null]`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
