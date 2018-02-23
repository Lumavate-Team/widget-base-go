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
    5, PropertyOptionsNumeric{0, 1000}}
}

func (lp *LumavateProperties) GetAllProperties() [] PropertyType {
	return [] PropertyType {
		lp.GetTextProperty(),
		lp.GetNumericProperty(),
	}
}

func (lp *LumavateProperties) GetAllComponents() [] Component {
	props := [] PropertyType {
		&PropertyText{
			&PropertyBase{"sampleText2", "General", "General Settings", "Example Text Parm", ""},
			"Hello", PropertyOptionsText{},
	  },
	}

	return [] Component {
		Component{"animal", "General", "dog", "Dog", "a", "Dog", props},
	}
}
