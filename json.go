package telego

import "github.com/chococola/telego/internal/json"

// SetJSONMarshal set JSON marshal func used in Telego
//
// Warning: Panics if passed func is nil
//
// Warning: This method is not concurrently-safe, do not call if bot is already running
func SetJSONMarshal(marshal func(v any) ([]byte, error)) {
	if marshal == nil {
		panic("Telego: nil marshal func not allowed")
	}
	json.Marshal = marshal
}

// SetJSONUnmarshal set JSON unmarshal func used in Telego
// Note: Unmarshal func should support unmarshalling into interface types if the struct field is populated with
// the correct type, not all libraries support this
//
// Warning: Panics if passed func is nil
//
// Warning: This method is not concurrently-safe, do not call if bot is already running
func SetJSONUnmarshal(unmarshal func(data []byte, v any) error) {
	if unmarshal == nil {
		panic("Telego: nil unmarshal func not allowed")
	}
	json.Unmarshal = unmarshal
}
