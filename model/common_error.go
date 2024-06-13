package model

import (
	"strconv"
	"time"
)

type CommonResponse struct {
	Status    string    `json:"status,omitempty"`
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func NewCommonResponse(status int, message string) *CommonResponse {
	return &CommonResponse{Status: strconv.Itoa(status), Message: message, Timestamp: time.Now()}
}

type CommonErrorResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewCommonErrorResponse(status int, message string) *CommonErrorResponse {
	return &CommonErrorResponse{Status: status, Message: message}
}
