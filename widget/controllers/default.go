package controllers

import (
  "github.com/astaxie/beego"
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type MainController struct {
  beego.Controller
}

type LumavateRequest struct {
  Payload struct {
    Data struct {
      SampleText string
    }
  }
}

func (this *MainController) Get() {
  pwa_jwt := this.Ctx.GetCookie("pwa_jwt")

  no_auth_redirect_url := fmt.Sprintf("https://%s?u=https://%s/%s/%s/%s",
    this.Ctx.Input.Host(),
    this.Ctx.Input.Host(),
    this.Ctx.Input.Param(":ic"),
    this.Ctx.Input.Param(":url_ref"),
    this.Ctx.Input.Param(":wid"),
    )

  widget_data_url := fmt.Sprintf("/pwa/v1/widget-instances/%s",
    this.Ctx.Input.Param(":wid"),
    )

  s := Signer{}
  signed_widget_data_url := fmt.Sprintf("%s%s",
    os.Getenv("BASE_URL"),
    s.GetSignature("get", widget_data_url, []byte{}))

  req, _ := http.NewRequest("GET", signed_widget_data_url, nil)
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer " + pwa_jwt)

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  if res.StatusCode == 401 {
    this.Ctx.Redirect(302, no_auth_redirect_url)
  }

  luma_response := LumavateRequest{}
  json.Unmarshal(body, &luma_response)
  this.Data["data"] = luma_response.Payload.Data
  this.TplName = "index.tpl"
}
