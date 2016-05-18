package tskv_test

import (
	"testing"

	"github.com/brimstone/go-tskv"
)

func TestNew(t *testing.T) {
	_, err := tskv.New()
	if err != nil {
		t.Fatal(err)
	}
}
