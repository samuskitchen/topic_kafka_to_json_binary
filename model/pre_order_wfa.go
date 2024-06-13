package model

import (
	"time"
)

//easyjson:json
type PreOrderWorkForceApp struct {
	MessageUniqueID  string            `json:"message_unique_id" validate:"required"`
	InvoiceID        string            `json:"invoice_id"`
	ClientID         string            `json:"client_id" validate:"required,max=10,min=1"`
	CreatedBy        string            `json:"created_by"`
	FingerPrint      string            `json:"finger_print"`
	Lat              float64           `json:"lat"`
	Lng              float64           `json:"lng"`
	Channel          string            `json:"channel" validate:"required"`
	PurchaseDate     string            `json:"purchase_date" validate:"required"`
	Route            int               `json:"route" validate:"required"`
	BlockReason      string            `json:"block_reason"`
	PaymentCondition string            `json:"payment_condition" validate:"required,max=4,min=1"`
	PaymentMethod    string            `json:"payment_method"`
	PaymentType      string            `json:"payment_type"`
	TransactionType  string            `json:"transaction_type" validate:"required,max=4,min=1"`
	OrderType        int               `json:"order_type" validate:"required"`
	DeliveryDate     string            `json:"delivery_date"`
	CustomerPhone    string            `json:"customer_phone"`
	Status           string            `json:"status" validate:"required"`
	Retry            int               `json:"retry,omitempty"`
	CreateWfa        time.Time         `json:"createWfa,omitempty"`
	Items            []OrderWfaItems   `json:"items" validate:"required,dive"`
	Bonuses          []BonusesWfaItems `json:"bonuses" validate:"required"`
}

//easyjson:json
type OrderWfaItems struct {
	Material             string `json:"material" validate:"required,max=18,min=1"`
	PromotionCode        string `json:"promotion_code" validate:"required,max=8,min=1"`
	Quantity             int    `json:"quantity"`
	SalesUnit            string `json:"sales_unit,omitempty"`
	DeliveryPriority     int    `json:"delivery_priority"`
	Usage                string `json:"usage" validate:"max=4"`
	SuggestedOrder       bool   `json:"suggested_order"`
	SuggestedOrderOrigen string `json:"suggested_order_origen"`
	PaymentType          string `json:"payment_type"`
	InvoiceID            string `json:"invoice_id"`
}

//easyjson:json
type BonusesWfaItems struct {
	Material    string `json:"material" validate:"required,max=18,min=1"`
	Quantity    int    `json:"quantity"`
	UnitMeasure string `json:"unit_measure,omitempty"`
	Usage       string `json:"usage" validate:"max=4"`
	BonusGroup  string `json:"group_bonus"`
}

//easyjson:json
type InfoOrderType struct {
	MessageUniqueID string   `json:"messages_uniq_id"`
	DocType         string   `json:"doc_type"`
	PaymentType     string   `json:"payment_type"`
	Items           []string `json:"items"`
}

//easyjson:json
type MessageUniqueAndType struct {
	MessageUniqueID string `json:"messages_uniq_id"`
	DocType         string `json:"doc_type"`
	PaymentType     string `json:"payment_type"`
}
