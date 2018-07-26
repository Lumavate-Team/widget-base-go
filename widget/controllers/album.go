package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  "encoding/json"
  "widget/models"
  _"fmt"
  "os"
)

type AlbumController struct {
  common_controller.LumavateController
}

func (this *AlbumController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }

  this.Data["taskURL"] = luma_response.Payload.Data.FormAction
  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["WidgetInstancePrefix"] = os.Getenv("WIDGET_URL_PREFIX") + this.Ctx.Input.Param(":wid")
  this.Data["data"] = luma_response.Payload.Data
  this.TplName = "album.tpl"
}