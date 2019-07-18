package network

import "encoding/json"

// Response ...
type Response struct {
	Body []byte
}

// Decode ...
func (r Response) Decode(dest interface{}) error {
	return json.Unmarshal(r.Body, dest)
}
