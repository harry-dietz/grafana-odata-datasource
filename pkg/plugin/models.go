package plugin

type queryModel struct {
	EntitySet        entitySet         `json:"entitySet"`
	TimeProperty     property          `json:"timeProperty"`
	Properties       []property        `json:"properties"`
	From             string            `json:"from"`
	To               string            `json:"to"`
	FilterConditions []filterCondition `json:"filterConditions"`
}

type entitySet struct {
	Name       string `json:"name"`
	EntityType string `json:"entityType"`
}

type property struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type filterCondition struct {
	Property property `json:"property"`
	Operator string   `json:"operator"`
	Value    string   `json:"value"`
}
