package calculators

import (
	"fmt"
	"go_lab_4/core"
	"math"
)

func GetCableCalculator() *core.Calculator {
	return &core.Calculator{
		Measurements: []core.MeasurementType{
			{Name: "shortCircuitCurrent", Label: "Струм КЗ", Units: "А"},
			{Name: "shortCircuitTime", Label: "Фіктивний час вимикання струму КЗ", Units: "C"},
			{Name: "expectedLoad", Label: "Розрахункове навантаження", Units: "кВт"},
			{Name: "maxLoadTime", Label: "Час максимального навантаження", Units: "год"},
		},
		Selectors: []core.Selector{
			{Name: "isolationMaterial", Label: "Матеріал ізоляції", Options: map[string]string{
				"pvc":    "ПВХ",
				"rubber": "Резина",
				"paper":  "Папір",
			}},
			{Name: "wireMaterial", Label: "Матеріал провідника", Options: map[string]string{
				"copper":   "Мідь",
				"aluminum": "Алюміній",
			}},
		},
		CalculationLogic: cableLogic,
	}
}

type MaterialKey struct {
	WireMaterial      string
	IsolationMaterial string
}

func cableLogic(i map[string]float64, s map[string]string) []core.MeasurementType {

	economyCurrentDensityForMaterials := map[MaterialKey][3]float64{
		{"copper", "paper"}:    {3.0, 2.5, 2.0},
		{"copper", "rubber"}:   {3.5, 3.1, 2.7},
		{"copper", "pvc"}:      {3.5, 3.1, 2.7},
		{"aluminum", "paper"}:  {1.6, 1.4, 1.2},
		{"aluminum", "rubber"}: {1.9, 1.7, 1.6},
		{"aluminum", "pvc"}:    {1.9, 1.7, 1.6},
	}

	thermalStabilityForMaterials := map[string]float64{
		"paper":   92,
		"plastic": 75,
		"rubber":  65,
	}

	shortCircuitCurrent := i["shortCircuitCurrent"]
	shortCircuitTime := i["shortCircuitTime"]
	expectedLoad := i["expectedLoad"]
	maxLoadTime := i["maxLoadTime"]
	isolationMaterial := s["isolationMaterial"]
	wireMaterial := s["wireMaterial"]

	normalCurrent := expectedLoad / 2 / math.Sqrt(3.0) / 10

	economyCurrentDensityForMaxLoadTimes :=
		economyCurrentDensityForMaterials[MaterialKey{wireMaterial, isolationMaterial}]

	var economyCurrentDensity float64

	if maxLoadTime < 3000 {
		economyCurrentDensity = economyCurrentDensityForMaxLoadTimes[0]
	} else if maxLoadTime < 5000 {
		economyCurrentDensity = economyCurrentDensityForMaxLoadTimes[1]
	} else {
		economyCurrentDensity = economyCurrentDensityForMaxLoadTimes[2]
	}

	economyCrossSection := normalCurrent / economyCurrentDensity

	thermalStableCrossSection := shortCircuitCurrent * math.Sqrt(shortCircuitTime) /
		thermalStabilityForMaterials[isolationMaterial]

	crossSection := math.Max(economyCrossSection, thermalStableCrossSection)

	return []core.MeasurementType{
		{"crossSection", "Поперечний переріз провідника", "кв. мм", fmt.Sprintf("%.2f", crossSection)},
	}
}
