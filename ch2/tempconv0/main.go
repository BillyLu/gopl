// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// Celsius c
type Celsius float64
// Fahrenheit f
type Fahrenheit float64  

// const temp
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

//CToF : c to f
func CToF(c Celsius) Fahrenheit{
	return Fahrenheit(c*9/5 + 32)
}

// FToC : f to c
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
