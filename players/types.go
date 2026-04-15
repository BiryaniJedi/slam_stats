package players

type PlayerApiRes struct {
	People []Player `json:"people"`
}

type Player struct {
	Id         int      `json:"id"`
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	CurAge     int      `json:"currentAge"`
	PrimaryPos FieldPos `json:"primaryPosition"`
	Active     bool     `json:"active"`
}

type FieldPos struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Abbrv string `json:"abbreviation"`
}
