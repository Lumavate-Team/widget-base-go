package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  "encoding/json"
  "widget/models"
  "fmt"
  _"os"
)

type MainController struct {
  common_controller.LumavateController
}

func (this *MainController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }

  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["data"] = luma_response.Payload.Data
  this.TplName = "index.tpl"
}

func (this *MainController) Post() {
  luma_response := models.LumavateRequest{}
  json.Unmarshal(this.LumavateGetData(), &luma_response)

  cameraData := models.CameraBase {}

  if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cameraData); err != nil {
    fmt.Println(err)
  }

  fmt.Println(cameraData)

  this.Data["json"] = nil
  this.Ctx.Output.SetStatus(200)
  this.ServeJSON()
}
