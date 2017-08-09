package main

import (
  "fmt"
  "html/template"
  
  "github.com/monologid/apollo-8/framework"
  "github.com/labstack/echo"
  echoMiddleware "github.com/labstack/echo/middleware"
)

func InitializeServer() {
  config := ReadBaseConfig()
  
  e := echo.New()
  e.HideBanner = true
  e.Renderer = &framework.TemplateRenderer{
    Templates: template.Must(template.ParseGlob(config.Apollo8.Websrc + "/*.html")),
  }
  e.Static(`/assets`, config.Apollo8.Websrc + `/assets`)
  
  e.Use(echoMiddleware.Secure())
  e.Use(echoMiddleware.Logger())
  e.Use(echoMiddleware.Recover())
  
  // initialize framework
  framework.InitializeRouter(e)
  
  fmt.Println(`-------------------------------`)
  fmt.Println(`## Apollo8 has been launched!`)
  fmt.Println(`## Host => ` + APP_HOST + ` Port => :` + APP_PORT)
  fmt.Println(`-------------------------------`)
  
  e.Logger.Fatal(e.Start(APP_HOST + `:` + APP_PORT))
}