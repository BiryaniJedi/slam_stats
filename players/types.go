package players

import (
	"fmt"
	"strconv"
	"strings"
)

type PlayerRespMLBAM struct {
	People []Player `json:"people"`
}

type Player struct {
	Id         int          `json:"id"` // pkey
	Active     *bool        `json:"active"`
	FirstName  *string      `json:"firstName"`
	LastName   *string      `json:"lastName"`
	BirthDate  *string      `json:"birthDate"`
	Height     *string      `json:"height"`
	Weight     *float64     `json:"weight"`
	PrimaryPos *FieldPos    `json:"primaryPosition"`
	BatSide    *BatterSide  `json:"batSide"`
	PitchHand  *PitcherHand `json:"pitchHand"`
	CurTeam    *Team        `json:"currentTeam"`
}

func (p Player) ToResponse() (PlayerResponse, error) {
	var heightInches *float64 = nil
	if p.Height != nil {
		height_in, err := heightStrToInches(*p.Height)
		if err != nil {
			return PlayerResponse{}, fmt.Errorf("Error converting Player MLBAM to Player Response: %w", err)
		}
		heightInches = new(float64(height_in))
	}
	var batSideCode *string = nil
	if p.BatSide != nil {
		batSideCode = &(*p.BatSide).Code
	}
	var pitchHandCode *string = nil
	if p.PitchHand != nil {
		pitchHandCode = &(*p.PitchHand).Code
	}
	return PlayerResponse{
		Id:         p.Id,
		Active:     p.Active,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		BirthDate:  p.BirthDate,
		Height_in:  heightInches,
		Weight_lbs: p.Weight,
		PrimaryPos: p.PrimaryPos,
		BatSide:    batSideCode,
		PitchHand:  pitchHandCode,
		CurTeam:    p.CurTeam,
	}, nil
}

type PlayerResponse struct {
	Id         int       `json:"id"`
	Active     *bool     `json:"active"`
	FirstName  *string   `json:"firstName"`
	LastName   *string   `json:"lastName"`
	BirthDate  *string   `json:"birthDate"`
	Height_in  *float64  `json:"height_in"`
	Weight_lbs *float64  `json:"weight_lbs"`
	PrimaryPos *FieldPos `json:"primaryPosition"`
	BatSide    *string   `json:"batSide"`
	PitchHand  *string   `json:"pitchHand"`
	CurTeam    *Team     `json:"currentTeam"`
}

type Team struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	ApiLink string `json:"link"`
}

func heightStrToInches(heightStr string) (int, error) {
	if len(heightStr) < 5 {
		return 0, fmt.Errorf("Height string is not long enough")
	}
	split := strings.Split(strings.TrimSuffix(heightStr, "\""), "' ")
	if len(split) != 2 {
		return 0, fmt.Errorf("Height string has too many values? got %d values, should be 2", len(split))
	}
	var sum int = 0
	for idx, sec := range split {
		intVal, err := strconv.Atoi(sec)
		if err != nil {
			return 0, fmt.Errorf("Error parsing int, got=%s", sec)
		}
		if idx == 0 {
			sum += 12 * intVal
			continue
		}
		sum += intVal
	}
	return sum, nil
}

type FieldPos struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Abbrv string `json:"abbreviation"`
}

type BatterSide struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type PitcherHand struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
