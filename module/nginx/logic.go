package nginx

import (
  "os"
  "github.com/monologid/apollo-8/helper"
)

type Logic struct{}

// GetNginxDirPath is a method to get nginx dir path from config.toml
func (l *Logic) GetNginxDirPath() string {
  configFile := helper.ReadFile(`config.toml`)
  configBase := helper.ReadToml(configFile)
  
  return (((configBase.(map[string]interface{}))["nginx"]).(map[string]interface{}))["dir"].(string)
}

// FindAll is a method to list all existing nginx .conf files
func (l *Logic) FindAll() []string {
  nginxPath := l.GetNginxDirPath()
  
  var filename []string
  files := helper.ReadDir(nginxPath)
  for _, file := range files {
    filename = append(filename, file.Name())
  }
  
  return filename
}

// FindOne is a method to read nginx file by filename
func (l *Logic) FindOne(filename string) string {
  nginxPath := l.GetNginxDirPath()
  filePath := nginxPath + `/` + filename
  return helper.ReadFile(filePath)
}

// Create is a method to create nginx file
func (l *Logic) Create(filename string) error {
  nginxPath := l.GetNginxDirPath()
  filePath := nginxPath + `/` + filename
  
  file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644) 
  if err != nil {
    return err
  }
  
  if err := file.Close(); err != nil {
    return err
  }
  
  return nil
}

// Update is a method to update nginx config file based on filename
func (l *Logic) Update(filename string, content string) error {
  nginxPath := l.GetNginxDirPath()
  filePath := nginxPath + `/` + filename
  
  if err := os.Remove(filePath); err != nil {
    return err
  }
  
  file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
  if err != nil {
    return err
  }
  defer file.Close()
  
  _, errWriteString := file.Write([]byte(content))
  if errWriteString != nil {
    return errWriteString
  }
  
  errFileSync := file.Sync()
  if errFileSync != nil {
    return errFileSync
  }
  
  return nil
}