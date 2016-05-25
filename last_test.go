package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestLast(t *testing.T) {
	kv, err := tskv.New(&tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = kv.InsertNow(1)
	if err != nil {
		t.Fatal(err)
	}
	last := kv.Last()
	if last.Int() != 1 {
		t.Fatal("Expected last element to be 1")
	}
}
