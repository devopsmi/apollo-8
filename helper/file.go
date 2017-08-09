package helper

import (
  "io/ioutil"
  "log"
  "os"
  
  "github.com/BurntSushi/toml"
)

// ReadDir is a method to read directory
func ReadDir (path string) []os.FileInfo {
  files, err := ioutil.ReadDir(path)
  if err != nil {
    log.Fatal(err)
  }
  
  return files
}

// ReadFile is a method to read a file
func ReadFile (path string) string {
  file, err := ioutil.ReadFile(path)
  if err != nil {
    log.Fatal(err)
  }
  
  return string(file)
}

// ReadToml is a method to read a .toml file
func ReadToml (file string) interface{} {
  var schema interface{}
  if _, err := toml.Decode(file, &schema); err != nil {
    log.Fatal(err)
  }
  return schema
}