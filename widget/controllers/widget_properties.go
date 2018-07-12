package controllers

import (
  properties "github.com/Lumavate-Team/lumavate-go-common/properties"
  components "github.com/Lumavate-Team/lumavate-go-common/components"
)

type LumavateProperties struct {
}

/*
 * Returns all properties for the widget
 */
func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
  return [] properties.PropertyType {
    &properties.PropertyText{
      &properties.PropertyBase{"formAction", "Actions", "Microservices", "Photo Storage URI", ""}, "", properties.PropertyOptionsText{}},
    components.GetNavBarProperty(),
    components.GetNavBarItemsProperty(),
  }
}

/*
 * Returns all components for the widget
 */
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
    components.GetNavBarComponent(),
    components.GetNavBarItemComponent(),
  }
}
