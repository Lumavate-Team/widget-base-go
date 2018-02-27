package controllers

import (
  properties "github.com/Lumavate-Team/go-properties"
	_ "os"
)

type LumavateProperties struct {
}

func (lp *LumavateProperties) GetTextProperty() *properties.PropertyText {
  return &properties.PropertyText{
    &properties.PropertyBase{"sampleText", "General", "General Settings", "Example Text Parm", ""},
    "Hello", properties.PropertyOptionsText{}}
}

func (lp *LumavateProperties) GetNumericProperty() *properties.PropertyNumeric {
  return &properties.PropertyNumeric{
    &properties.PropertyBase{"sampleNumeric", "General", "General Settings", "Example Numeric Parm", ""},
    5, properties.PropertyOptionsNumeric{0, 1000}}
}

func (lp *LumavateProperties) GetComponentProperty() *properties.PropertyComponent {
  return &properties.PropertyComponent{
    &properties.PropertyBase{"quote", "Quote", "Settings", "Example Component Parm", ""},
    lp.GetQuoteComponent(), properties.PropertyOptionsComponent{[] string {"all"}, [] *properties.Component {lp.GetQuoteComponent()} },
	}
}

func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
	return [] properties.PropertyType {
		lp.GetTextProperty(),
		lp.GetNumericProperty(),
		lp.GetComponentProperty(),
	}
}

func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
	return [] *properties.Component {
		lp.GetQuoteComponent(),
	}
}

func (lp *LumavateProperties) GetQuoteComponent() *properties.Component {
	//return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
	return properties.LoadComponent("https://experience.john.labelnexusdev.com", "1.0.0", "quote")
}
