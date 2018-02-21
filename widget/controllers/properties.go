package controllers

import (
  "github.com/astaxie/beego"
  "encoding/json"
)

const (
	PropertyTranslatedtText string = "translated-text"
	PropertyColor string = "color"
	PropertyComponent string = "component"
	PropertyComponents string = "components"
	PropertyImage string = "image-upload"
	PropertyCheckbox string = "checkbox"
	PropertyToggle string = "toggle"
	PropertyDropdown string = "dropdown"
	PropertyMultiselect string = "multiselect"
	PropertyPageLink string = "page-link"
)

///////////////////////////////////////////////////
// Generic Interface
///////////////////////////////////////////////////
type PropertyType interface {
}

type PropertyBase struct {
  Name string
  Classification string
  Section string
  Label string
  HelpText string
}

///////////////////////////////////////////////////
// Text Property
///////////////////////////////////////////////////
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
		"numeric",
		(*Copy)(p),
	})
}


///////////////////////////////////////////////////
// Numeric Property
///////////////////////////////////////////////////
type PropertyOptionsNumeric struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type PropertyNumeric struct {
	*PropertyBase
	Default int `json:"default"`
  Options PropertyOptionsNumeric `json:"options"`
}

func (p *PropertyNumeric) MarshalJSON() (b []byte, e error) {
  type Copy PropertyNumeric

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"text",
		(*Copy)(p),
	})
}

///////////////////////////////////////////////////
// Numeric Property
///////////////////////////////////////////////////
type ComponentOptions struct {
	Categories [] string `json"categories"`
	Components [] Component `json"components"`
}

type DropdownOptions struct {
	Options map[string]string
}

type MultiselectOption struct {
	Value string `json:"value"`
	DisplayValue string `json:"displayValue"`
}

type MultiselectOptions struct {
	Options [] MultiselectOption
}

type Component struct {
	Category string `json:"category"`
	Section string `json:"section"`
	Type string `json:"type"`
	DisplayName string `json:"displayName"`
	Icon string `json:"icon"`
	Label string `json:"label"`
	//Properties [] Property `json:"properties"`
}
	
type PropertyController struct {
  beego.Controller
}

func (this *PropertyController) Get() {
  textWidget := PropertyText{
		&PropertyBase{"sampleText", "General", "General Settings", "Example Text Parm", ""},
		"Hello", PropertyOptionsText{}}

  numericWidget := PropertyNumeric{
		&PropertyBase{"sampleNumeric", "General", "General Settings", "Example Numeric Parm", ""},
		5, PropertyOptionsNumeric{}}

  this.Data["json"] = [] PropertyType { &textWidget, &numericWidget }
  this.ServeJSON()
}
