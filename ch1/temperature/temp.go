package main

import (
	"fmt"
	"os"
	"strconv"
)

func toCelcius(fahren float64) float64 {
	return (fahren - 32) * 5/9 
}

func toFahrenheit(celcius float64) float64 {
	return celcius * 9 / 5 + 32
}

func usage() int {
	var s, sep string
	var lenargs int = len(os.Args)
	
	for i := 1; i < lenargs; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	if lenargs > 1 {
		fmt.Println("CMD: " + s)
	} else {
		println("Usage: <prog> <tof|toc> <value>")
		println(" e.g.: <prog> toc 32")
	}
	return lenargs
}

func main() {
	if 3 > usage() {
		return
	}
	cmd := os.Args[1]
	arg := os.Args[2]
	const bitSize = 64 
	// Don't think about it to much. It's just 64 bits.
	floatNum, err := strconv.ParseFloat(arg, bitSize)
	if nil != err {
		println("Error converting " + arg + ", " + err.Error())
	}

	if cmd == "toc" {
		result := toCelcius(floatNum)
		fmt.Printf("Input %s, Result: %0.2f\n", 
			arg, result)	
	} else if cmd == "tof" {
		result := toFahrenheit(floatNum)
		fmt.Printf("Input %s, Result: %0.2f\n", 
			arg, result)	
	} else {
		println("Unknown cmd " + cmd + ", (expected: toc, tof)")	
	}
	println("Completed.")
}
