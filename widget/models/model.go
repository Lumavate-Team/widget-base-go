package models

import (
  common "github.com/Lumavate-Team/lumavate-go-common"
  widget "github.com/Lumavate-Team/lumavate-go-common/models"
)

type MainController struct {
  common.LumavateController
}

type LumavateRequest struct {
  Payload struct {
    Data struct {
      widget.CommonWidgetStruct
      FormAction string `json:"formAction"`
    }
  }
}

type CameraImage struct {
  Image string `json:"data"`
}

type PhotoInfo struct {
  Filename string `json:"fileName"`
  SessionId string `json:"sessionId"`
  Image CameraImage `json:"image"`
}

type ImageResponse struct {
  Resources [] ImageData `json:"resources"`
}

type ImageData struct {
  Backup bool `json:"backup"`
  Bytes int `json:"bytes"`
  CreatedAt string `json:"created_at"`
  Format string `json:"format"`
  Height int `json:"height"`
  Public_Id string `json:"public_id"`
  ResourceType string `json:"resource_type"`
  SecureUrl string `json:"secure_url"`
  Type string `json:"type"`
  Url string `json:"url"`
  Version int `json:"version"`
  Width int `json:"width"`
}