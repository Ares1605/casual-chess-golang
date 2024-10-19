package apiresps

import (
  "encoding/json"
  "errors"
)

// baseResp contains the common fields for all API responses
type baseResp struct {
  Success bool   `json:"success"`
  Message string `json:"message,omitempty"`
  Error   *struct {
    Message string `json:"message"`
    Type    string `json:"type"`
  } `json:"error,omitempty"`
}

// Resp is a generic struct that can handle any data type
type Resp[T any] struct {
  baseResp
  Data T `json:"data,omitempty"`
}
type JunkResp struct {
  Resp[any]
}

// UnmarshalJSON is a generic function to unmarshal any APIResponse
func UnmarshalResp[T any](data []byte) (*Resp[T], error) {
  var response Resp[T]
  err := json.Unmarshal(data, &response)
  if err != nil {
    return nil, err
  }

  // Validate the structure based on the Success field
  if response.Success {
    if response.Error != nil {
      return nil, errors.New("success response cannot have an error field")
    }
  } else {
    if response.Error == nil {
      return nil, errors.New("error response must have an error field")
    }
  }
  return &response, nil
}
