package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestGetSince(t *testing.T) {
	kv, err := tskv.New()
	if err != nil {
		t.Fatal(err)
	}

	err = kv.InsertNow(1)
	if err != nil {
		t.Fatal(err)
	}

	err = kv.Insert(time.Now().Add(-time.Hour), 2)
	if err != nil {
		t.Fatal(err)
	}

	elements := kv.GetSince(time.Now().Add(-time.Minute))
	/*
		length := len(elements)
		if length != 1 {
			t.Fatal("Expected 1 element in slice, got ", length)
		}
	*/
	value := elements.Oldest().Int()
	if value != 1 {
		t.Fatal("Expected oldest to be 1, got ", value)
	}
}