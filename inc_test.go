package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestInc(t *testing.T) {
	kv, err := tskv.New(&tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = kv.Inc(time.Now(), 1)
	if err != nil {
		t.Fatal(err)
	}
	err = kv.Inc(time.Now(), 1)
	if err != nil {
		t.Fatal(err)
	}
	value := kv.Last().Int()
	if value != 2 {
		t.Fatal("Expected 2, got", value)
	}
}

func TestIncNow(t *testing.T) {
	kv, err := tskv.New(&tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = kv.IncNow(1)
	if err != nil {
		t.Fatal(err)
	}
	err = kv.IncNow(1)
	if err != nil {
		t.Fatal(err)
	}
	value := kv.Last().Int()
	if value != 2 {
		t.Fatal("Expected 2, got", value)
	}
}
