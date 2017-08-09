package framework

import (
  "github.com/monologid/apollo-8/module/nginx"
  "github.com/labstack/echo"
)

var nullMap = map[string]interface{}{}

// InitializeRouter is a method to initialize main router.
// All the main routes will be defined here
func InitializeRouter(e *echo.Echo) {
  DashboardHandler(e)
  NginxHandler(e)
}

// DashboardHandler is router handler for Dashboard
func DashboardHandler(e *echo.Echo) {
  e.GET(`/`, dashboardRender)
}

func dashboardRender(c echo.Context) error {
  return c.Render(200, `index.html`, nullMap)
}

// NginxHandler is router handler for Nginx
func NginxHandler(e *echo.Echo) {
  e.GET(`/api/v1/nginx`, nginxGetAll)
  e.GET(`/api/v1/nginx/:filename`, nginxGetByFilename)
  e.POST(`/api/v1/nginx/:filename`, nginxCreateConfig)
  e.PUT(`/api/v1/nginx/:filename`, nginxUpdateConfig)
}

func nginxGetAll(c echo.Context) error {
  ng := new(nginx.Controller)
  data, err := ng.GetAllConfigFilename()
  code, json := Response(data, err)
  return c.JSON(code, json)
}

func nginxGetByFilename(c echo.Context) error {
  ng := new(nginx.Controller)
  data, err := ng.ReadConfigFile(c.Param(`filename`))
  code, json := Response(data, err)
  return c.JSON(code, json)
}

func nginxCreateConfig(c echo.Context) error {
  ng := new(nginx.Controller)
  data, err := ng.CreateConfigFile(c.Param(`filename`))
  code, json := Response(data, err)
  return c.JSON(code, json)
}

func nginxUpdateConfig(c echo.Context) error {
  ng := new(nginx.Controller)
  m := new(nginx.Model)
  
  if err := c.Bind(m); err != nil {
    code, json := Response(map[string]interface{}{}, err)
    return c.JSON(code, json)
  }
  
  data, err := ng.UpdateConfigFile(c.Param(`filename`), m.Content)
  code, json := Response(data, err)
  return c.JSON(code, json)
}