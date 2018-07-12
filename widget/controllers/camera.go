package controllers

import (
  common_models "github.com/Lumavate-Team/lumavate-go-common/models"
  "encoding/json"
  "widget/models"
  "fmt"
  _"math/rand"
  _"strconv"
)

type CameraController struct {
  MainController
}

func (this *CameraController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }

  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["data"] = luma_response.Payload.Data
  this.TplName = "camera.tpl"
}

func (this *MainController) Post() {
  luma_response := models.LumavateRequest{}
  json.Unmarshal(this.LumavateGetData(), &luma_response)
  identity := this.Ctx.Request.Header.Get("PWA-S")

  cameraData := models.CameraImage {}
  if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cameraData); err != nil {
    fmt.Println(err)
  }

  dataToSend := models.PhotoInfo{"peace", identity, cameraData}
  this.Data["json"] = nil
  this.Ctx.Output.SetStatus(200)

  q := fmt.Sprintf("%v/cloudinary/upload",
    luma_response.Payload.Data.FormAction)

  b, _ := json.Marshal(&dataToSend)
  resp, status := this.LumavatePost(q, b, true)


  if status != "200" {
    var error_response common_models.ErrorResponse
    json.Unmarshal(resp, &error_response)
    fmt.Println(error_response)
  }

  this.Ctx.Redirect(302, "/")
}