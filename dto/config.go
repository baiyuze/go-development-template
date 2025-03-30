package dto

type RequestBody struct {
	WidgetName string                            `json:"widgetName"`
	WidgetId   string                            `json:"widgetId"`
	Menu       []map[string]interface{}          `json:"menu"`
	MenuMap    map[string]map[string]interface{} `json:"menuMap"`
	Type       int                               `json:"type"`
}

type MenuConfig struct {
	Menu    []map[string]interface{}
	MenuMap map[string]map[string]interface{}
}

type ErrorType struct {
	string
	int
	error
}
