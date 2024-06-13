package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"topic_kafka_to_json_binary/model"
	"topic_kafka_to_json_binary/services/external"

	jsoniter "github.com/json-iterator/go"
)

var jsonString = "{\n  \"message_unique_id\": \"6654afa030dcaf6d79849be3\",\n  \"client_id\": \"0001018188\",\n  \"created_by\": \"122235\",\n  \"finger_print\": \"\",\n  \"lat\": 0,\n  \"lng\": 0,\n  \"channel\": \"CAP\",\n  \"route\": 101405,\n  \"purchase_date\": \"2024-05-27\",\n  \"block_reason\": \"\",\n  \"payment_condition\": \"CP00\",\n  \"payment_method\": \"\",\n  \"transaction_type\": \"01\",\n  \"order_type\": 111,\n  \"delivery_date\": \"\",\n  \"customer_phone\": \"+50245195899\",\n  \"status\": \"NEW\",\n  \"items\": [\n    {\n      \"quantity\": 6,\n      \"material\": \"BA017975\",\n      \"sales_unit\": \"ST\",\n      \"delivery_priority\": 0,\n      \"usage\": \"0000\",\n      \"suggested_order\": true,\n      \"suggested_order_origen\": \"PULL\",\n      \"promotion_code\": \"00\",\n      \"payment_type\": \"CASH\",\n      \"invoice_id\": \"1\"\n    },\n    {\n      \"quantity\": 6,\n      \"material\": \"AA685001\",\n      \"sales_unit\": \"ST\",\n      \"delivery_priority\": 0,\n      \"usage\": \"0000\",\n      \"suggested_order\": true,\n      \"suggested_order_origen\": \"ESTRATEGICO\",\n      \"promotion_code\": \"00\",\n      \"payment_type\": \"CASH\",\n      \"invoice_id\": \"1\"\n    },\n    {\n      \"quantity\": 1,\n      \"material\": \"BA022398\",\n      \"sales_unit\": \"CS\",\n      \"delivery_priority\": 0,\n      \"usage\": \"0000\",\n      \"suggested_order\": false,\n      \"suggested_order_origen\": \"NULL\",\n      \"promotion_code\": \"00\",\n      \"payment_type\": \"CASH\",\n      \"invoice_id\": \"1\"\n    },\n    {\n      \"quantity\": 3,\n      \"material\": \"AA015001\",\n      \"sales_unit\": \"ST\",\n      \"delivery_priority\": 0,\n      \"usage\": \"0000\",\n      \"suggested_order\": false,\n      \"suggested_order_origen\": \"NULL\",\n      \"promotion_code\": \"00\",\n      \"payment_type\": \"CASH\",\n      \"invoice_id\": \"1\"\n    }\n  ],\n  \"totals\": {\n    \"currency\": \"Q\",\n    \"discount\": 0,\n    \"subTotal\": 1235036,\n    \"taxe\": 148202,\n    \"total\": 1383240,\n    \"totalQuantityBottle\": 15,\n    \"totalQuantityBox\": 1,\n    \"totalWithDiscount\": 1383240\n  },\n  \"bonuses\": [\n    {\n      \"quantity\": 1,\n      \"material\": \"BA022398\",\n      \"unit_measure\": \"ST\",\n      \"group_bonus\": \"0201338\"\n    }\n  ],\n  \"country\": \"GT\"\n}"

// ExampleByteJson generating the mapper bytes with the easyjson library and create file in memory
func ExampleByteJson() error {
	var wg sync.WaitGroup
	errorChan := make(chan error, 2)

	data := model.PreOrderWorkForceApp{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		//invalid character '\'' looking for beginning of object key string
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		fmt.Println(err.Error())
	}

	myFile := &model.MyFile{
		Reader: bytes.NewReader(dataBytes),
		Mif: model.MyFileInfo{
			NameInfo: "workforce-pre-order-created.json",
			Data:     dataBytes,
		},
	}

	var file http.File = myFile

	requestData := model.DocumentRequest{
		Channel: "CAP",
		Country: "GT",
		Process: "data_integration",
		Bucket:  "workforce-pre-order-created",
	}

	// Llamadas asincrónicas a funciones
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		err = external.UploadFile(context.Background(), file, dataBytes, requestData)
		duration := time.Since(start)
		fmt.Printf("Duración de UploadFile: %v\n", duration)
		if err != nil {
			errorChan <- err
		}
	}()

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	close(errorChan)

	// Comprobar si hubo errores
	for err := range errorChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// ExampleByteJsonV2 generating the mapper bytes with the easyjson library
func ExampleByteJsonV2() error {
	var wg sync.WaitGroup
	errorChan := make(chan error, 2)

	data := model.PreOrderWorkForceApp{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		//invalid character '\'' looking for beginning of object key string
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		fmt.Println(err.Error())
	}

	requestData := model.DocumentRequest{
		Channel: "CAP",
		Country: "GT",
		Process: "data_integration",
		Bucket:  "workforce-pre-order-created",
	}

	// Llamadas asincrónicas a funciones
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		err = external.UploadFileV2(context.Background(), "workforce-pre-order-created.json", dataBytes, requestData)
		duration := time.Since(start)
		fmt.Printf("Duración de UploadFileV2: %v\n", duration)
		if err != nil {
			errorChan <- err
		}
	}()

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	close(errorChan)

	// Comprobar si hubo errores
	for err := range errorChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// ExampleByteJsonV25 generating the mapper bytes with the jsoniter library
func ExampleByteJsonV25() error {
	var wg sync.WaitGroup
	errorChan := make(chan error, 2)

	data := model.PreOrderWorkForceApp{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		//invalid character '\'' looking for beginning of object key string
	}

	var jsonLibrary = jsoniter.ConfigCompatibleWithStandardLibrary
	dataBytes, err := jsonLibrary.Marshal(&data)
	if err != nil {
		return err
	}

	requestData := model.DocumentRequest{
		Channel: "CAP",
		Country: "GT",
		Process: "data_integration",
		Bucket:  "workforce-pre-order-created",
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		err = external.UploadFileV2(context.Background(), "workforce-pre-order-created.json", dataBytes, requestData)
		duration := time.Since(start)
		fmt.Printf("Duración de UploadFileV2: %v\n", duration)
		if err != nil {
			errorChan <- err
		}
	}()

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	close(errorChan)

	// Comprobar si hubo errores
	for err := range errorChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func ExampleByteJsonV3() error {
	var wg sync.WaitGroup
	errorChan := make(chan error, 2)

	data := model.PreOrderWorkForceApp{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		//invalid character '\'' looking for beginning of object key string
	}

	var jsonLibrary = jsoniter.ConfigCompatibleWithStandardLibrary
	dataBytes, err := jsonLibrary.Marshal(&data)
	if err != nil {
		return err
	}

	requestData := model.DocumentRequest{
		Channel: "CAP",
		Country: "GT",
		Process: "data_integration",
		Bucket:  "workforce-pre-order-created",
	}

	// Llamadas asincrónicas a funciones
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		err = external.UploadFileV3(context.Background(), "workforce-pre-order-created.json", dataBytes, requestData)
		duration := time.Since(start)
		fmt.Printf("Duración de UploadFileV3: %v\n", duration)
		if err != nil {
			errorChan <- err
		}
	}()

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	close(errorChan)

	// Comprobar si hubo errores
	for err := range errorChan {
		if err != nil {
			return err
		}
	}

	return nil
}
