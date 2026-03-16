package lark

import "encoding/json"

// Element is the interface that all card components implement.
// The cardTag method returns the JSON "tag" value for the component.
type Element interface {
	cardTag() string
}

// Elements is a slice of Element that handles JSON marshaling
// by injecting the "tag" field from each element's cardTag() method.
type Elements []Element

// MarshalJSON implements json.Marshaler, injecting the "tag" field into each element.
func (e Elements) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte("null"), nil
	}
	result := make([]json.RawMessage, 0, len(e))
	for _, el := range e {
		data, err := json.Marshal(el)
		if err != nil {
			return nil, err
		}
		var m map[string]json.RawMessage
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, err
		}
		tagBytes, _ := json.Marshal(el.cardTag())
		m["tag"] = tagBytes
		injected, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		result = append(result, injected)
	}
	return json.Marshal(result)
}
