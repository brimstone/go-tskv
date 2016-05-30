package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestSet(t *testing.T) {
	kv, err := tskv.New(tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = kv.Set(time.Now(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetNow(t *testing.T) {
	kv, err := tskv.New(tskv.Config{
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
}
