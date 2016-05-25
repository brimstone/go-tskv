package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestCleanup(t *testing.T) {
	kv, err := tskv.New(&tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = kv.Insert(time.Now().Add(-time.Hour), 1)
	if err != nil {
		t.Fatal(err)
	}

	elements := kv.GetInt()
	for _, v := range elements {
		if v != 0 {
			t.Fatal("Expected value to be 0, got ", v)
		}
	}
}
