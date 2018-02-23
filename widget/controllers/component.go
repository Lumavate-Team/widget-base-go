package controllers

import (
  "encoding/json"
)

type ComponentOptions struct {
	Categories [] string `json"categories"`
	Components [] Component `json"components"`
}

type Component struct {
	Category string `json:"category"`
	Section string `json:"section"`
	Type string `json:"type"`
	DisplayName string `json:"displayName"`
	Icon string `json:"icon"`
	Label string `json:"label"`
	Properties [] PropertyType `json:"properties"`
}

type PropertyOptionsComponent struct {
	Categories [] string `json:"categories"`
	Components [] Component `json:"components"`
}

type PropertyComponent struct {
	*PropertyBase
	Default Component `json:"default"`
	Options PropertyOptionsComponent `json:"options"`
}

func (p *PropertyComponent) MarshalJSON() (b []byte, e error) {
  type Copy PropertyComponent

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"component",
		(*Copy)(p),
	})
}

type PropertyComponents struct {
	*PropertyBase
	Default [] Component `json:"default"`
	Options PropertyOptionsComponent `json:"options"`
}

func (p *PropertyComponents) MarshalJSON() (b []byte, e error) {
  type Copy PropertyComponents

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"component",
		(*Copy)(p),
	})
}

