package tskv_test

import (
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestNewNoDurationOrFrequency(t *testing.T) {
	_, err := tskv.New(tskv.Config{})
	if err == nil {
		t.Fatal("Err expected, got no error")
	}
}

func TestNewNoDuration(t *testing.T) {
	_, err := tskv.New(tskv.Config{
		Frequency: time.Second,
	})
	if err == nil {
		t.Fatal("Err expected, got no error")
	}
}

func TestNewNoFrequency(t *testing.T) {
	_, err := tskv.New(tskv.Config{
		Duration: time.Second,
	})
	if err == nil {
		t.Fatal("Err expected, got no error")
	}
}

func TestNewFrequencyMoreThanDuration(t *testing.T) {
	_, err := tskv.New(tskv.Config{
		Duration:  time.Second,
		Frequency: time.Second * 5,
	})
	if err == nil {
		t.Fatal("Err expected, got no error")
	}
}
