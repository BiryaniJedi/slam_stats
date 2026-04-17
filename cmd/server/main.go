package main

import (
	"encoding/json"
	"fmt"
	"github.com/BiryaniJedi/slam_stats/players"
	"github.com/BiryaniJedi/slam_stats/responses"
	"github.com/BiryaniJedi/slam_stats/utils"
	"io"
	"net/http"
)

func getPlayers(searchUrl string) ([]players.Player, error) {
	res, err := http.Get(searchUrl)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("search url %s resulted in a bad response", searchUrl)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	apiResp := players.PlayerRespMLBAM{}
	if err := json.Unmarshal(data, &apiResp); err != nil {
		return nil, err
	}
	return apiResp.People, nil
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_ = responses.RespondJSON(w, 200, map[string]bool{"ok": true})
	})

	//Returns first match of searching mlb api with given name
	mux.HandleFunc("/api/players/{fullName}", func(w http.ResponseWriter, req *http.Request) {
		searchName := req.PathValue("fullName")
		searchUrl := fmt.Sprintf("https://statsapi.mlb.com/api/v1/people/search?names=%s", searchName)
		playersList, err := getPlayers(searchUrl)
		if err != nil {
			responses.RespondError(w, 404, fmt.Sprintf("%v", err))
			return
		}
		playerResps, _ := utils.MapSafe(playersList, func(player players.Player) (players.PlayerResponse, error) {
			return player.ToResponse()
		})
		responses.RespondJSON(w, 200, playerResps)
	})

	//Returns the player who's id matches the query
	mux.HandleFunc("/api/player/{id}", func(w http.ResponseWriter, req *http.Request) {
		searchId := req.PathValue("id")
		searchUrl := fmt.Sprintf("https://statsapi.mlb.com/api/v1/people/%s?hydrate=currentTeam", searchId)
		playersList, err := getPlayers(searchUrl)
		if err != nil {
			responses.RespondError(w, 404, fmt.Sprintf("%v", err))
			return
		}
		if len(playersList) == 0 {
			responses.RespondError(w, 404, fmt.Sprintf("No Player with id=%s", searchId))
			return
		}
		playerResp, err := playersList[0].ToResponse()
		if err != nil {
			responses.RespondError(w, 404, fmt.Sprintf("%s %s's data is invalid!", playersList[0].FirstName, playersList[0].LastName))
		}
		responses.RespondJSON(w, 200, playerResp)
	})

	// Serve frontend static files in production
	mux.Handle("/", http.FileServer(http.Dir("./frontend/dist")))

	http.ListenAndServe(":8080", mux)
}
