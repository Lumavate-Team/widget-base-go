package controllers

import (
  common_models "github.com/Lumavate-Team/lumavate-go-common/models"
  "encoding/json"
  "widget/models"
  "fmt"
  "time"
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

func (this *CameraController) Post() {
  fmt.Println("STEP 1")
  luma_response := models.LumavateRequest{}
  json.Unmarshal(this.LumavateGetData(), &luma_response)

  fmt.Println("IN PHOTO SUBMIT")

  cameraData := models.CameraImage {}
  if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cameraData); err != nil {
    fmt.Println(err)
  }

  fmt.Println("GETTING CAMERA DATA")
  fmt.Println(cameraData)

  filename := time.Now()

  dataToSend := models.PhotoInfo{filename.String(), "", cameraData}
  this.Data["json"] = nil
  this.Ctx.Output.SetStatus(200)

  q := fmt.Sprintf("%v/cloudinary/upload",
    luma_response.Payload.Data.FormAction)

  fmt.Println("URL TO SEND DATA")
  fmt.Println(q)

  b, _ := json.Marshal(&dataToSend)
  fmt.Println("THESE ARE THE BYTES")
  fmt.Println(string(b[:1000]))
  resp, status := this.LumavatePost(q, b, true)

  fmt.Println(status)


  if status != "200" {
    var error_response common_models.ErrorResponse
    json.Unmarshal(resp, &error_response)
    fmt.Println(error_response)
  }
}
