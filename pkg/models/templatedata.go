package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRSFToken string //Cross request forgery token
	Flash     string //flash message
	Warning   string
	Error     string
}
