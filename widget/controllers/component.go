package controllers

type ComponentOptions struct {
	Categories [] string `json"categories"`
	Components [] Component `json"components"`
}

type Component struct {
	Category string `json:"category"`
	Section string `json:"section"`
	Type string `json:"type"`
	DisplayName string `json:"displayName"`
	Icon string `json:"icon"`
	Label string `json:"label"`
	Properties [] PropertyType `json:"properties"`
}
