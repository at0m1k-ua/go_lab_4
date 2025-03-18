package calculators

import (
	"fmt"
	"go_lab_4/core"
	"math"
)

func GetHpnemCalculator() *core.Calculator {
	return &core.Calculator{
		Measurements: []core.MeasurementType{
			{Name: "resistance", Label: "Активний опір", Units: "Ом"},
			{Name: "reactance", Label: "Реактивний опір", Units: "Ом"},
		},
		CalculationLogic: hpnemLogic,
	}
}

func hpnemLogic(i map[string]float64, _ map[string]string) []core.MeasurementType {
	r := i["resistance"]
	reactance := i["reactance"]

	xt := 11.1 * 115 * 115 / 100 / 6.3
	x := reactance + xt
	z := math.Sqrt(r*r + x*x)
	i3p := 115 * 1000 / math.Sqrt(3.0) / z
	i2p := i3p * math.Sqrt(3.0) / 2
	k := 11.0 * 11 / 115 / 115

	r10kv := r * k
	x10kv := x * k
	z10kv := math.Sqrt(r10kv*r10kv + x10kv*x10kv)
	i3p10kv := 11 * 1000 / math.Sqrt(3.0) / z10kv
	i2p10kv := i3p10kv * math.Sqrt(3.0) / 2

	return []core.MeasurementType{
		{"i3p", "Трифазний струм КЗ на шинах 110 кВ", "A", fmt.Sprintf("%.2f", i3p)},
		{"i2p", "Двофазний струм КЗ на шинах 110 кВ", "A", fmt.Sprintf("%.2f", i2p)},
		{"i3p10kv", "Трифазний струм КЗ на шинах 10 кВ", "A", fmt.Sprintf("%.2f", i3p10kv)},
		{"i2p10kv", "Двофазний струм КЗ на шинах 10 кВ", "A", fmt.Sprintf("%.2f", i2p10kv)},
	}
}
