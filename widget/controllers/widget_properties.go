package controllers

import (
  properties "github.com/Lumavate-Team/go-properties"
  _ "os"
)

type LumavateProperties struct {
}

/*
 * Defines a simple, text property
 */
func (lp *LumavateProperties) GetTextProperty() *properties.PropertyText {
  return &properties.PropertyText{
    &properties.PropertyBase{"sampleText", "Quote", "Simple Settings", "Hello World Text", ""},
    "Hello, World!", properties.PropertyOptionsText{}}
}

/*
 * Gets a property that allows entry for 'Quote' data
 */
func (lp *LumavateProperties) GetComponentProperty() *properties.PropertyComponent {
  return &properties.PropertyComponent{
    &properties.PropertyBase{"quote", "Quote", "Quote Settings", "Quote Data", ""},
    lp.GetQuoteComponent(), properties.PropertyOptionsComponent{[] string {"all"}, [] *properties.Component {lp.GetQuoteComponent()} },
  }
}

/*
 * Gets a description for the 'Quote' component.  This is defined in a central place
 */
func (lp *LumavateProperties) GetQuoteComponent() *properties.Component {
  //return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
  return properties.LoadComponent("https://experience.john.labelnexusdev.com", "1.0.0", "quote")
}

/*
 * Gets a property that allows entry for 'NavBar' data
 */
func (lp *LumavateProperties) GetNavBarProperty() *properties.PropertyComponent {
  return &properties.PropertyComponent{
    &properties.PropertyBase{"navBar", "Nav Bar", "Nav Bar Properties", "Nav Bar", ""},
    lp.GetNavBarComponent(), properties.PropertyOptionsComponent{[] string {"navBar"}, [] *properties.Component {lp.GetNavBarComponent()} },
  }
}

/*
 * Gets a property that allows entry for 'NavBarItems' data
 */
func (lp *LumavateProperties) GetNavBarItemsProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents{
    &properties.PropertyBase{"navBarItems", "Nav Bar", "Nav Bar Items", "Nav Bar", ""},
    [] *properties.Component {}, properties.PropertyOptionsComponent{[] string {"navBarItem"}, [] *properties.Component {lp.GetNavBarItemComponent()} },
  }
}

/*
 * Gets a description for the 'NavBar' component.  This is defined in a central place
 */
func (lp *LumavateProperties) GetNavBarComponent() *properties.Component {
  //return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
        comp := properties.LoadComponent("https://experience.john.labelnexusdev.com", "1.0.0", "navBar")
        comp.Category = "navBar"
        return comp
}

/*
 * Gets a description for the 'NavBarItem' component.  This is defined in a central place
 */
func (lp *LumavateProperties) GetNavBarItemComponent() *properties.Component {
  //return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
        comp :=properties.LoadComponent("https://experience.john.labelnexusdev.com", "1.0.0", "navBarItem")
        comp.Category = "navBarItem"
        return comp
}

/*
 * Returns all properties for the widget
 */
func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
  return [] properties.PropertyType {
    lp.GetTextProperty(),
    lp.GetNavBarProperty(),
    lp.GetNavBarItemsProperty(),
    lp.GetComponentProperty(),
  }
}

/*
 * Returns all components for the widget
 */
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
    lp.GetQuoteComponent(),
    lp.GetNavBarComponent(),
    lp.GetNavBarItemComponent(),
  }
}
