package controllers

import (
  "encoding/json"
  "widget/models"
  "fmt"
)

type AddController struct {
  MainController
}

func (this *AddController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }

  q := fmt.Sprintf("%v/cloudinary/images",
  luma_response.Payload.Data.FormAction)

  resp, _ := this.LumavateGet(q, true)

  imageData := models.ImageResponse{}
  if err := json.Unmarshal(resp, &imageData); err != nil {
    fmt.Println(err)
  }

  fmt.Sprintf("%+v", imageData)

  if ((len(imageData.Resources)) == 0) {
    this.Data["visible"] = false
  } else {
    this.Data["visible"] = true
  }

  fmt.Printf("%+v",imageData.Resources)

  this.Data["taskURL"] = luma_response.Payload.Data.FormAction
  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["data"] = luma_response.Payload.Data
  this.Data["images"] = imageData.Resources
  this.Data["wid"] = this.Ctx.Input.Param(":wid")
  this.TplName = "add.tpl"
}