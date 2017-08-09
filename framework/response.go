package framework

import (
  "net/http"
)

// JsonResponse is a schema for JSON response
type JsonResponse struct {
  StatusCode int32 `json:"code"`
  Message string `json:"message"`
  Err string `json:"error"`
  Data interface{} `json:"data"`
} 

// SuccessResponse is a method to generate
// success response for router
// Parameters:
// @statusCode int32
// @obj[0] interface{} - it should contain json data
// @obj[1] string - message (optional)
func SuccessResponse(statusCode int32, obj ...interface{}) JsonResponse {
  resp := JsonResponse{}
  resp.StatusCode = statusCode
  resp.Data = obj[0]
  
  if len(obj) > 1 {
    resp.Message = obj[1].(string)
  }
  
  return resp
}

// ErrorResponse is a method to generate
// error response for router
// Parameters:
// @statusCode int32
// @obj[0] interface{} - it should contain json data
// @obj[1] string - message (leave it "" or nil)
// @obj[2] string - error message
func ErrorResponse(statusCode int32, obj ...interface{}) JsonResponse {
  resp := JsonResponse{}
  resp.StatusCode = statusCode
  resp.Data = obj[0]
  
  if len(obj) > 1 {
    resp.Message = obj[1].(string)
  }
  
  if len(obj) > 2 {
    resp.Err = obj[2].(string)
  }
  
  return resp
}

// Response is a method to generate final response
func Response (schema interface{}, err error) (int, interface{}) {
  if err != nil {
    return http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, schema, `error`, err.Error()) 
  } else {
    return http.StatusOK, SuccessResponse(http.StatusOK, schema, `success`)
  }
}