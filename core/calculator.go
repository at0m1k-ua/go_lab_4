package core

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Calculator struct {
	Measurements     []MeasurementType
	Selectors        []Selector
	Output           []MeasurementType
	CalculationLogic func(i map[string]float64, s map[string]string) []MeasurementType

	i map[string]float64
	s map[string]string
}

func (c *Calculator) GetHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/calculator.html"))

		m := TemplateModel{}
		m.Measurements = c.Measurements
		m.Selectors = c.Selectors

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

		for i := range m.Selectors {
			m.Selectors[i].Value = r.FormValue(m.Selectors[i].Name)
		}

		if err := c.Calculate(); err != nil {
			m.Error = err
		} else {
			m.CalcResult = c.Output
		}

		_ = tmpl.Execute(w, m)
	}
}

func (c *Calculator) Calculate() error {
	c.i = make(map[string]float64)
	c.s = make(map[string]string)

	for _, m := range c.Measurements {
		if m.Value == "" {
			return fmt.Errorf("поле \"%s\" не заповнене", m.Label)
		}

		if val, err := strconv.ParseFloat(m.Value, 64); err != nil {
			return fmt.Errorf("поле \"%s\" містить невірне значення", m.Label)
		} else {
			c.i[m.Name] = val
		}
	}

	for _, item := range c.Selectors {
		if item.Value == "" {
			return fmt.Errorf("селектор \"%s\" не обрано", item.Label)
		}

		if _, ok := item.Options[item.Value]; !ok {
			return fmt.Errorf("селектор \"%s\" має некоректне значення", item.Label)
		}

		c.s[item.Name] = item.Value
	}

	c.Output = c.CalculationLogic(c.i, c.s)
	return nil
}
