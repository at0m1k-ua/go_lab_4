package main

import (
	"fmt"
	"go_lab_4/calculators"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/cable", calculators.GetCableCalculator().GetHandler())
	http.HandleFunc("/hpnem", calculators.GetHpnemCalculator().GetHandler())
	http.HandleFunc("/short_circuit", calculators.GetShortCircuitCalulator().GetHandler())

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
