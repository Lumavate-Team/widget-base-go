package controllers

type LumavateProperties struct {
}

func (lp *LumavateProperties) GetTextProperty() *PropertyText {
  return &PropertyText{
    &PropertyBase{"sampleText", "General", "General Settings", "Example Text Parm", ""},
    "Hello", PropertyOptionsText{}}
}

func (lp *LumavateProperties) GetNumericProperty() *PropertyNumeric {
  return &PropertyNumeric{
    &PropertyBase{"sampleNumeric", "General", "General Settings", "Example Numeric Parm", ""},
    5, PropertyOptionsNumeric{}}
}

func (lp *LumavateProperties) GetAllProperties() [] PropertyType {
	return [] PropertyType {
		lp.GetTextProperty(),
		lp.GetNumericProperty(),
	}
}
