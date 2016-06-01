package tskv_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/brimstone/go-tskv"
)

func TestMarshal(t *testing.T) {
	kv, err := tskv.New(tskv.Config{
		Duration:  time.Second * 5,
		Frequency: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	err = kv.Set(now, 1)
	if err != nil {
		t.Fatal(err)
	}

	jsondata, err := json.Marshal(kv)
	if err != nil {
		t.Fatal(err)
	}

	now = now.Truncate(time.Second).Add(-4 * time.Second)
	expected := "{\"" + now.Format(time.RFC3339) + "\":null"

	now = now.Add(time.Second)
	expected += ",\"" + now.Format(time.RFC3339) + "\":null"

	now = now.Add(time.Second)
	expected += ",\"" + now.Format(time.RFC3339) + "\":null"

	now = now.Add(time.Second)
	expected += ",\"" + now.Format(time.RFC3339) + "\":null"

	now = now.Add(time.Second)
	expected += ",\"" + now.Format(time.RFC3339) + "\":1}"

	if string(jsondata) != expected {
		t.Fatal("Expected", expected, "received", string(jsondata))
	}
}
