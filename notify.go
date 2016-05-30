package tskv

import "time"

func (t *Tskv) notify(when time.Time, value interface{}) {
	if len(t.subscribers) == 0 {
		return
	}
	kv, _ := New(Config{
		Duration:  t.duration,
		Frequency: t.frequency,
	})
	kv.Set(when, value)
	for _, c := range t.subscribers {
		go func(channel chan *Tskv) {
			channel <- kv
		}(c)
	}
}
