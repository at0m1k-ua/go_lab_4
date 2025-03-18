package calculators

import (
	"fmt"
	"go_lab_4/core"
	"math"
)

func GetShortCircuitCalulator() *core.Calculator {
	return &core.Calculator{
		Measurements: []core.MeasurementType{
			{Name: "shortCircuitPower", Label: "Потужність КЗ", Units: "МВт"},
		},
		CalculationLogic: shortCircuitLogic,
	}
}

func shortCircuitLogic(i map[string]float64, _ map[string]string) []core.MeasurementType {
	shortCircuitPower := i["shortCircuitPower"]

	const u1 = 10.5
	const u2 = 6.3

	xc := u1 * u1 / shortCircuitPower
	xt := u1 * u1 * u1 / 100 / u2

	sumX := xc + xt
	initialCurrent := u1 / math.Sqrt(3.0) / sumX

	return []core.MeasurementType{
		{Name: "initialCurrent", Label: "Струм КЗ", Units: "кА", Value: fmt.Sprintf("%.2f", initialCurrent)},
	}
}
