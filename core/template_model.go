package core

type TemplateModel struct {
	Measurements []MeasurementType
	Selectors    []Selector
	CalcResult   []MeasurementType
	Error        error
}
