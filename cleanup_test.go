package tskv_test

import (
	"fmt"
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
	fmt.Println(elements)
	/*
		if value != 1 {
			t.Fatal("Expected oldest to be 1, got ", value)
		}
	*/
}
