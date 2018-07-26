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

type Folder struct {
  Album string `json:"album"`
}

type CameraImage struct {
  Image string `json:"data"`
}

type PhotoInfo struct {
  Filename string `json:"fileName"`
  AlbumName string `json:"album"`
  Image CameraImage `json:"image"`
}

type ImageData struct {
  Backup int `json:"backup_bytes"`
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
  Tags [] string `json:"tags"`
}

type ImageResponse struct {
  AlbumData [] Albums `json:"albumData"`
  AlbumList [] string `json:"albums"`
  Resources [] ImageData `json:"allImages"`
}

type Albums struct {
  AlbumImages [] ImageData `json:"albumImages"`
  AlbumTitle string `json:"albumTitle"`
  PictureCount int `json:"pictureCount"`
}