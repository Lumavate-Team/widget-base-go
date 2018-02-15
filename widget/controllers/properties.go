package controllers

import (
  "github.com/astaxie/beego"
)

type PropertyOptions struct {
}

type TextProperty struct {
  Name string `json:"name"`
  Classification string `json:"classification"`
  Section string `json:"section"`
  Label string `json:"label"`
  Type string `json:"type"`
  Options PropertyOptions `json:"options"`
  HelpText string `json:"helpText"`
}

type PropertyController struct {
  beego.Controller
}

func (this *PropertyController) Get() {
  name := TextProperty{}
  name.Name = "sampleText"
  name.Classification = "General"
  name.Section = "General Settings"
  name.Label = "Example Go Text Parm"
  name.Type = "text"
  this.Data["json"] = [] TextProperty { name }
  this.ServeJSON()
}
