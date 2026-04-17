package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeightToInches(t *testing.T) {
	result, err := heightStrToInches("6' 4\"")
	assert.NoError(t, err)
	assert.Equal(t, 76, result)

	_, err = heightStrToInches("")
	assert.EqualError(t, err, "Height string is not long enough")

	result, err = heightStrToInches("5' 11\"")
	assert.NoError(t, err)
	assert.Equal(t, 71, result)

	result, err = heightStrToInches("5' 11\" 6' 11\"")
	assert.EqualError(t, err, "Height string has too many values? got 3 values, should be 2")

	result, err = heightStrToInches("5'    ")
	assert.EqualError(t, err, "Error parsing int, got=   ")
}

func createTestPlayer() Player {
	return Player{
		Active:     true,
		Id:         123,
		FirstName:  "Test",
		LastName:   "Test",
		BirthDate:  "2000-01-01",
		Height:     "6' 0\"",
		Weight:     200,
		PrimaryPos: FieldPos{Code: "Test", Name: "Test", Type: "Test", Abbrv: "Test"},
		BatSide:    BatterSide{Code: "Test", Description: "Test"},
		PitchHand:  PitcherHand{Code: "Test", Description: "Test"},
		CurTeam:    Team{Id: 123, Name: "Test"},
	}
}

func TestToResponse(t *testing.T) {
	bad_player := Player{}
	_, err := bad_player.ToResponse()
	assert.EqualError(t, err, "Error converting Player MLBAM to Player Response: Height string is not long enough")

	good_player := createTestPlayer()
	good_resp, err := good_player.ToResponse()
	assert.NoError(t, err)
	assert.Equal(t, 72.0, good_resp.Height_in)
	assert.Equal(t, 200.0, good_resp.Weight_lbs)
	assert.Equal(t, 123, good_resp.CurTeam.Id)
}
