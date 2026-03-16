package lark

import "encoding/json"

// cardJSON is the wire format with schema injected.
type cardJSON struct {
	Schema   string    `json:"schema"`
	Config   *Config   `json:"config,omitempty"`
	CardLink *CardLink `json:"card_link,omitempty"`
	Header   Header    `json:"header,omitempty"`
	Body     Body      `json:"body,omitempty"`
}

// JSON marshals the card to JSON bytes with schema "2.0" injected.
func (c Card) JSON() ([]byte, error) {
	return json.Marshal(cardJSON{
		Schema:   "2.0",
		Config:   c.Config,
		CardLink: c.CardLink,
		Header:   c.Header,
		Body:     c.Body,
	})
}

// MustJSON is like JSON but panics on error.
func (c Card) MustJSON() []byte {
	data, err := c.JSON()
	if err != nil {
		panic(err)
	}
	return data
}
