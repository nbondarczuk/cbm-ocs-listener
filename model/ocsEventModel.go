package model

type (
	Parameter struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	Parameters struct {
		Parameter []Parameter `json:"parameter"`
	}

	OcsEvent struct {
		EventId          string     `json:"eventId"`
		EventType        string     `json:"eventType"`
		EventDescription string     `json:"eventDescription"`
		SourceSystem     string     `json:"sourceSystem"`
		SourceDate       string     `json:"sourceDate"`
		Parameters       Parameters `json:"parameters"`
	}
)
