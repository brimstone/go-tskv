package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestSubscribe(t *testing.T) {
	kv, err := tskv.New(&tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	channel := kv.Subscribe()

	err = kv.Inc(time.Now(), 42)
	if err != nil {
		t.Fatal(err)
	}

	element := <-channel

	value := element.Last().Int()
	if value != 42 {
		t.Fatal("Expected 42, got", value)
	}
}
