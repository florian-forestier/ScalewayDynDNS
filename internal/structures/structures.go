package structures

type Connect struct {
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	ChangeType string    `json:"changeType"`
	Records    []Records `json:"records"`
}

type Records struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Priority int    `json:"priority"`
	TTL      int    `json:"ttl"`
	Data     string `json:"data"`
}
