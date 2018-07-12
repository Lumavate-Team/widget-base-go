package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  "encoding/json"
  "widget/models"
  "fmt"
  "os"
)

type MainController struct {
  common_controller.LumavateController
}

func (this *MainController) GetRedirectUrl(controller string) string {
  instance_id := this.Ctx.Input.Param(":wid")
  fmt.Println("Instance: ", instance_id)
  url := fmt.Sprintf("%s%s%s%s%s",
    os.Getenv("PROTO"),
    this.Ctx.Input.Host(),
    this.Data["WidgetUrlPrefix"],
    instance_id,
    controller)
  fmt.Println(url)
  return url
}

func (this *MainController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }

  q := fmt.Sprintf("%v/cloudinary/images",
  luma_response.Payload.Data.FormAction)

  resp, _ := this.LumavateGet(q, true)

  // fmt.Println(string(resp[:1500]))

  imageData := models.ImageResponse {}
  if err := json.Unmarshal(resp, &imageData); err != nil {
    fmt.Println(err)
  }

  length := 0
  for _, x := range imageData.Resources {
    length = length + 1
    fmt.Println(x)
  }

  this.Data["images"] = imageData.Resources
  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["WidgetInstancePrefix"] = os.Getenv("WIDGET_URL_PREFIX") + this.Ctx.Input.Param(":wid")
  this.Data["data"] = luma_response.Payload.Data
  this.Data["album"] = true
  this.Data["length"] = length
  this.TplName = "directory.tpl"
}

func (this *MainController) Post() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }






}






