package main

import (
	"encoding/json"
	"fmt"
	"github.com/BiryaniJedi/slam_stats/players"
	"github.com/BiryaniJedi/slam_stats/responses"
	"io"
	// "log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_ = responses.RespondJSON(w, 200, map[string]bool{"ok": true})
	})

	//Returns first match of searching mlb api with given name
	mux.HandleFunc("/api/players/{fullName}", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL)
		searchName := req.PathValue("fullName")
		searchUrl := fmt.Sprintf("https://statsapi.mlb.com/api/v1/people/search?names=%s", searchName)
		res, err := http.Get(searchUrl)
		if err != nil {
			responses.RespondError(w, 404, fmt.Sprintf("%v", err))
		}
		if res.StatusCode != 200 {
			responses.RespondError(w, 404, "Bad response from fetching player data")
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			responses.RespondError(w, 404, fmt.Sprintf("%v", err))
		}
		apiResp := players.PlayerApiRes{}
		if err := json.Unmarshal(data, &apiResp); err != nil {
			responses.RespondError(w, 404, fmt.Sprintf("%v", err))
		}

		if len(apiResp.People) == 0 {
			responses.RespondError(w, 404, "No players found maching that name")
		}
		_ = responses.RespondJSON(w, 200, apiResp.People[0])
	})

	http.ListenAndServe(":8080", mux)
}
