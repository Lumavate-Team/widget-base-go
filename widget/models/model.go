package models

import (
  common "github.com/Lumavate-Team/lumavate-go-common"
  widget "github.com/Lumavate-Team/lumavate-go-common/models"
  component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type MainController struct {
  common.LumavateController
}

type LumavateRequest struct {
  Payload struct {
    Data struct {
      widget.CommonWidgetStruct
      Quote component_data.QuoteStruct
      SampleText string
    }
  }
}

type CameraBase struct {
  CameraData string `json:"cameraData"`
}