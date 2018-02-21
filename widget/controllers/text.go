package controllers

import (
  "encoding/json"
)

type PropertyOptionsText struct {
	ReadOnly bool
	Rows int
}

type PropertyText struct {
	*PropertyBase
	Default string
  Options PropertyOptionsText
}

func (p *PropertyText) MarshalJSON() (b []byte, e error) {
  type Copy PropertyText

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"text",
		(*Copy)(p),
	})
}

type PropertyTranslatedText struct {
	*PropertyBase
	Default string
  Options PropertyOptionsText
}

func (p *PropertyTranslatedText) MarshalJSON() (b []byte, e error) {
  type Copy PropertyTranslatedText

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"translated-text",
		(*Copy)(p),
	})
}

