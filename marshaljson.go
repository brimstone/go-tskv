package tskv

import (
	"encoding/json"
	"time"
)

// MarshalJSON provides an interface for json.Marshal
func (t *Tskv) MarshalJSON() ([]byte, error) {
	all := make(map[string]interface{})
	for k, v := range t.GetMap() {
		all[k.Format(time.RFC3339)] = v.Value()
	}
	return json.Marshal(all)
}
