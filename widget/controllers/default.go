package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  common_models "github.com/Lumavate-Team/lumavate-go-common/models"
  "encoding/json"
  "widget/models"
  "fmt"
  "os"
  "time"
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

  imageData := models.ImageResponse{}
  if err := json.Unmarshal(resp, &imageData); err != nil {
    fmt.Println(err)
  }

  fmt.Printf("%+v", imageData)

  if ((len(imageData.Resources)) == 0) && ((len(imageData.AlbumList)) == 0) {
    this.Data["visible"] = false
  } else {
    this.Data["visible"] = true
  }

  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["images"] = imageData.Resources
  this.Data["albums"] = imageData.AlbumData
  this.Data["WidgetInstancePrefix"] = os.Getenv("WIDGET_URL_PREFIX") + this.Ctx.Input.Param(":wid")
  this.Data["data"] = luma_response.Payload.Data
  this.TplName = "directory.tpl"
}

func (this *MainController) Post() {
  luma_response := models.LumavateRequest{}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }

  albumData := models.Folder{}
  albumData.Album = this.GetString("albumTitle")
  this.Data["json"] = nil
  this.Ctx.Output.SetStatus(200)

  q := fmt.Sprintf("%v/cloudinary/upload/album",
    luma_response.Payload.Data.FormAction)

  b, _ := json.Marshal(&albumData)
  resp, status := this.LumavatePost(q, b, true)


  if status != "200" {
    var error_response common_models.ErrorResponse
    json.Unmarshal(resp, &error_response)
    fmt.Println(error_response)
  }

  time.Sleep(1 * time.Second)
  this.Ctx.Redirect(302, this.GetRedirectUrl("/"))
}