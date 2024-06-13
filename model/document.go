package model

import (
	"time"
)

type Document struct {
	ContainerPath  string     `json:"container_path"`
	FileName       string     `json:"file_name"`
	FileUniqueName string     `json:"file_unic_name"`
	ContentType    string     `json:"conten_type"`
	Properties     Properties `json:"properties"`
	Metadata       Metadata   `json:"metadata"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type DocumentUploadedInfo struct {
	Size         int64
	CreationTime time.Time
	LastModified time.Time
}

type DocumentResponse struct {
	IDObject   string     `json:"id_object"`
	Properties Properties `json:"properties"`
	Metadata   Metadata   `json:"metadata"`
}

type DocumentB64Response struct {
	Base64 string `json:"base64"`
	Format string `json:"format"`
}

type DocumentRequest struct {
	Country  string `json:"country"`
	Channel  string `json:"channel"`
	User     string `json:"user"`
	Process  string `json:"process"`
	Customer string `json:"customer"`
	Bucket   string `json:"bucket"`
}

type DownloadResponse struct {
	Name        string
	Size        int64
	Buffer      []byte
	ContentType string
}

type MetadataRequest struct {
	IDObject string   `json:"id_object"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Channel  string `json:"channel"`
	User     string `json:"user"`
	Process  string `json:"process"`
	Customer string `json:"customer"`
}

type Properties struct {
	URL          string    `json:"url" bson:"url"`
	CreationTime time.Time `json:"creation_time"`
	LastModified time.Time `json:"last_modified"`
	Type         string    `json:"type"`
	Size         int64     `json:"size"`
}
