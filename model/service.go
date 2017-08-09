package model

// ServiceSchema is a schema for service object
type ServiceSchema struct {
  Name string `json:"name"`
  Path string `json:"path"`
  Port int32 `json:"port"`
}