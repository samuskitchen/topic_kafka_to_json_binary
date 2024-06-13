package wrapper

import "time"

// Result contiene la duración de la ejecución y cualquier error que pueda ocurrir
type Result struct {
	Duration time.Duration
	Err      error
}

// MeasureTime envuelve la llamada a una función y mide su tiempo de ejecución
func MeasureTime(resultChan chan Result, f error) {
	start := time.Now()
	err := f
	duration := time.Since(start)
	resultChan <- Result{Duration: duration, Err: err}
}
