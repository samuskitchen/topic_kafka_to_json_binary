package main

import (
	"fmt"
	"topic_kafka_to_json_binary/kit/wrapper"
	"topic_kafka_to_json_binary/services"
)

func main() {
	// Canal para recibir el resultado de la ejecución
	resultChan := make(chan wrapper.Result)

	for i := 0; i < 100; i++ {
		// Iniciar la medición de tiempo en una goroutine
		go wrapper.MeasureTime(resultChan, services.ExampleByteJsonV3())
	}

	// Obtener el resultado desde el canal
	result := <-resultChan
	if result.Err != nil {
		fmt.Printf("Error: %v\n", result.Err)
	} else {
		fmt.Printf("Duración de la ejecución: %v\n", result.Duration)
	}

}
