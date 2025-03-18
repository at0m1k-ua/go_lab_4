package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type MeasurementType struct {
	Name  string
	Label string
	Units string
	Value string
}

type TemplateModel struct {
	Measurements []MeasurementType
	CalcResult   []MeasurementType
	Error        error
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")

	m := TemplateModel{}
	m.Measurements = []MeasurementType{
		{Name: "power", Label: "Середньодобова потужність", Units: "МВт"},
		{Name: "deviation", Label: "Середньоквадратичне відхилення", Units: "МВт"},
	}

	if r.Method != "POST" {
		_ = tmpl.Execute(w, m)
		return
	}

	err := r.ParseForm()
	if err != nil {
		return
	}

	for i := range m.Measurements {
		m.Measurements[i].Value = r.FormValue(m.Measurements[i].Name)
	}

	if res, err := Calculate(m.Measurements); err != nil {
		m.Error = err
	} else {
		m.CalcResult = res
	}

	_ = tmpl.Execute(w, m)
}

func Calculate(measurements []MeasurementType) ([]MeasurementType, error) {
	i := make(map[string]float64)
	var o []MeasurementType

	for _, m := range measurements {
		if m.Value == "" {
			return o, fmt.Errorf("поле \"%s\" не заповнене", m.Label)
		}

		if val, err := strconv.ParseFloat(m.Value, 64); err != nil {
			return o, fmt.Errorf("поле \"%s\" містить невірне значення", m.Label)
		} else {
			i[m.Name] = val
		}

	}

	//powerAmt := i["power"]
	//deviationAmt := i["deviation"]
	//const cost = 7.0
	//
	//normalDist := distuv.Normal{
	//	Mu:    powerAmt,
	//	Sigma: deviationAmt,
	//}
	//
	//maxDeviation := powerAmt * 0.05
	//lowerLimit := powerAmt - maxDeviation
	//upperLimit := powerAmt + maxDeviation
	//
	//cdfLower := normalDist.CDF(lowerLimit)
	//cdfUpper := normalDist.CDF(upperLimit)
	//integralValue := cdfUpper - cdfLower
	//
	//profit := cost * 24 * powerAmt * integralValue
	//loss := cost * 24 * powerAmt * (1 - integralValue)
	//total := profit - loss
	//
	//o = []MeasurementType{
	//	{"total", "Добовий дохід електростанції", "тис. грн", fmt.Sprintf("%.2f", total)},
	//}

	return o, nil
}

func main() {
	http.HandleFunc("/", IndexHandler)

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
