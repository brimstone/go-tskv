package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestInsert(t *testing.T) {
	kv, err := tskv.New(&tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = kv.Insert(time.Now(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertNow(t *testing.T) {
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
}
