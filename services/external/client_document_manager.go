package external

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	customError "topic_kafka_to_json_binary/kit/errors"
	"topic_kafka_to_json_binary/model"

	httpLib "github.com/ApexDigitalM5/m5-gommon/http"
)

// UploadFile in the api-document-manager with file in memory or virtual
func UploadFile(ctx context.Context, file http.File, fileByte []byte, requestData model.DocumentRequest) error {
	fmt.Println("method: UploadFile")

	var BaseURL = "http://localhost:8985/api-document-manager"
	URL := fmt.Sprintf("%s/object", BaseURL)

	fileContents, errFile := io.ReadAll(file)
	if errFile != nil {
		return errFile
	}

	fi, errFile := file.Stat()
	if errFile != nil {
		return errFile
	}

	errFile = file.Close()
	if errFile != nil {
		return errFile
	}

	// Create a buffer to store the request body
	bodyBuffer := new(bytes.Buffer)

	// Create a new multipart writer with the buffer
	writer := multipart.NewWriter(bodyBuffer)

	// Create a new form field
	part, errFile := writer.CreateFormFile("file-stream", fi.Name())
	if errFile != nil {
		return errFile
	}

	// Write the contents of the file to the form field
	_, errFile = part.Write(fileContents)
	if errFile != nil {
		return errFile
	}

	// Data request convert to json
	requestDataMarshal, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if _, err = part.Write(requestDataMarshal); err != nil {
		return err
	}

	// added a new form field
	err = writer.WriteField("document", string(requestDataMarshal))
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", URL, bodyBuffer)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		if err = httpResponse.Body.Close(); err != nil {
			fmt.Println(fmt.Errorf("[Body.Close]: %+v ", err))
		}
	}()

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("[io.ReadAll]: %+v ", err))
		return err
	}
	fmt.Printf("[StatusCode]: %v ", httpResponse.StatusCode)
	//fmt.Printf("[Body]: %v ", string(body))

	if httpResponse.StatusCode != http.StatusOK && httpResponse.StatusCode != http.StatusCreated {
		var responseErr model.CommonErrorResponse
		if err = json.Unmarshal(body, &responseErr); err != nil {
			fmt.Println(fmt.Errorf("messages: [json.Unmarshal]: %+v", err))
			return err
		}

		return customError.NewCustomError(responseErr.Message, responseErr.Status)
	}

	//var response model.DocumentResponse
	//if err = json.Unmarshal(body, &response); err != nil {
	//	fmt.Println(fmt.Errorf("messages: [json.Unmarshal]: %+v", err))
	//	return err
	//}

	fmt.Println("Ok")
	return nil
}

// UploadFileV2 in the api-document-manager with file bytes ()
func UploadFileV2(ctx context.Context, fileName string, fileContent []byte, requestData model.DocumentRequest) error {
	fmt.Println("method: UploadFile")

	var BaseURL = "http://localhost:8985/api-document-manager"
	URL := fmt.Sprintf("%s/object", BaseURL)

	// Crear un buffer para almacenar el contenido del multipart/form-data
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Parte del archivo
	part, err := writer.CreateFormFile("file-stream", fileName)
	if err != nil {
		return err
	}

	// Copiar el contenido del archivo al multipart
	fileReader := bytes.NewReader(fileContent)
	if _, err = io.Copy(part, fileReader); err != nil {
		return err
	}

	// Parte del JSON
	jsonPart, err := writer.CreateFormField("document")
	if err != nil {
		return err
	}

	// Convertir el JSON a bytes y escribirlo en el multipart
	jsonBytes, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	if _, err = jsonPart.Write(jsonBytes); err != nil {
		return err
	}

	// Cerrar el writer para finalizar el multipart
	if err = writer.Close(); err != nil {
		return err
	}

	// Crear una nueva petición HTTP
	request, err := http.NewRequest("POST", URL, &b)
	if err != nil {
		return err
	}

	// Establecer el tipo de contenido a multipart/form-data con el boundary correcto
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Enviar la petición
	client := &http.Client{}
	httpResponse, err := client.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		if err = httpResponse.Body.Close(); err != nil {
			fmt.Println(fmt.Errorf("[Body.Close]: %+v ", err))
		}
	}()

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("[io.ReadAll]: %+v ", err))
		return err
	}
	fmt.Printf("[StatusCode]: %v ", httpResponse.StatusCode)
	//fmt.Printf("[Body]: %v ", string(body))

	if httpResponse.StatusCode != http.StatusOK && httpResponse.StatusCode != http.StatusCreated {
		var responseErr model.CommonErrorResponse
		if err = json.Unmarshal(body, &responseErr); err != nil {
			fmt.Println(fmt.Errorf("messages: [json.Unmarshal]: %+v", err))
			return err
		}

		return customError.NewCustomError(responseErr.Message, responseErr.Status)
	}

	//var response model.DocumentResponse
	//if err = json.Unmarshal(body, &response); err != nil {
	//	fmt.Println(fmt.Errorf("messages: [json.Unmarshal]: %+v", err))
	//	return err
	//}

	fmt.Println("Ok")
	return nil
}

// UploadFileV3 in the api-document-manager with file bytes but implement httpLib
func UploadFileV3(ctx context.Context, fileName string, fileContent []byte, requestData model.DocumentRequest) error {
	fmt.Println("method: UploadFile")
	httpClient := httpLib.NewHTTP()

	var BaseURL = "http://localhost:8985/api-document-manager"
	URL := fmt.Sprintf("%s/object", BaseURL)

	// Crear un buffer para almacenar el contenido del multipart/form-data
	var bodyBuffer bytes.Buffer
	writer := multipart.NewWriter(&bodyBuffer)

	// Parte del archivo
	part, err := writer.CreateFormFile("file-stream", fileName)
	if err != nil {
		return err
	}

	// Copiar el contenido del archivo al multipart
	fileReader := bytes.NewReader(fileContent)
	if _, err = io.Copy(part, fileReader); err != nil {
		return err
	}

	// Parte del JSON
	jsonPart, err := writer.CreateFormField("document")
	if err != nil {
		return err
	}

	// Convertir el JSON a bytes y escribirlo en el multipart
	jsonBytes, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	if _, err = jsonPart.Write(jsonBytes); err != nil {
		return err
	}

	// Cerrar el writer para finalizar el multipart
	if err = writer.Close(); err != nil {
		return err
	}

	header := []httpLib.Header{
		{
			Key:   "Content-Type",
			Value: writer.FormDataContentType(),
		},
	}

	options := &httpLib.Option{
		URL:     URL,
		Method:  http.MethodPost,
		Context: ctx,
		Header:  header,
		Body:    &bodyBuffer,
	}

	// Crear una nueva petición HTTP
	httpResponse, err := httpClient.ExecuteRequest(ctx, options)
	if err != nil {
		return err
	}

	defer func() {
		if err = httpResponse.Body.Close(); err != nil {
			fmt.Println(fmt.Errorf("[Body.Close]: %+v ", err))
		}
	}()

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("[io.ReadAll]: %+v ", err))
		return err
	}
	fmt.Printf("[StatusCode]: %v ", httpResponse.StatusCode)
	//fmt.Printf("[Body]: %v ", string(body))

	if httpResponse.StatusCode != http.StatusOK && httpResponse.StatusCode != http.StatusCreated {
		var responseErr model.CommonErrorResponse
		if err = json.Unmarshal(body, &responseErr); err != nil {
			fmt.Println(fmt.Errorf("messages: [json.Unmarshal]: %+v", err))
			return err
		}

		return customError.NewCustomError(responseErr.Message, responseErr.Status)
	}

	//var response model.DocumentResponse
	//if err = json.Unmarshal(body, &response); err != nil {
	//	fmt.Println(fmt.Errorf("messages: [json.Unmarshal]: %+v", err))
	//	return err
	//}

	fmt.Println("Ok")
	return nil
}
