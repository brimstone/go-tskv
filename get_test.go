package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestGetSince(t *testing.T) {
	kv, err := tskv.New(&tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = kv.SetNow(1)
	if err != nil {
		t.Fatal(err)
	}

	err = kv.Set(time.Now().Add(-time.Hour), 2)
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

func TestGetInt(t *testing.T) {
	kv, err := tskv.New(&tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = kv.SetNow(1)
	if err != nil {
		t.Fatal(err)
	}
	ints := kv.GetInt()
	if ints[0] != 1 {
		t.Fatal("Expected 1, got", ints[0])
	}
	for _, v := range ints[1:] {
		if v != 0 {
			t.Fatal("Expected 0, got", v)
		}
	}
}
