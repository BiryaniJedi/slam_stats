# MLB Stats API — Everything I Know

> This is compiled from training data, not official documentation. Treat it as a starting point and verify everything by actually hitting the endpoints. Things may have changed.

---

## The Basics

Base URL: `https://statsapi.mlb.com/api/v1`

- No API key required (historically — this may be changing)
- All requests are HTTP GET
- All responses are JSON
- `{ver}` in URL templates = `v1`
- No rate limit is officially documented, but hammering it will get you throttled

---

## URL Pattern

```
https://statsapi.mlb.com/api/v1/{endpoint}/{pathParam}?queryParam=value&queryParam=value
```

Path parameters go directly in the URL. Query parameters go after `?` separated by `&`.

The `hydrate` query parameter is special — it lets you embed related data inline rather than making multiple requests (more on this below).

---

## Player Endpoints

### Get a single player by ID
```
GET /people/{personId}
```
```
https://statsapi.mlb.com/api/v1/people/592450
```
Returns bio: name, position, bats/throws, birthdate, height, weight, current team, MLB debut date.

### Get multiple players by ID
```
GET /people?personIds=592450,660271
```

### Get all players in a season
```
GET /sports/1/players?season=2025
```
`sportId=1` is MLB. Returns every player on a 40-man roster. This is the closest thing to a search — fetch all, filter client-side. Slow but it works.

### Get a player's season stats
```
GET /people/{personId}/stats?stats=season&group=hitting&season=2024
```

`group` options: `hitting`, `pitching`, `fielding`
`stats` options: `season`, `career`, `yearByYear`, `yearByYearAdvanced`, `seasonAdvanced`, `careerAdvanced`

Advanced stats example (OPS+, ERA+, etc.):
```
GET /people/592450/stats?stats=seasonAdvanced&group=hitting&season=2024
```

### Hydrate stats directly onto a player object
```
GET /people/592450?hydrate=stats(group=hitting,type=season,season=2024)
```
Returns the player bio AND their stats in one call.

---

## The `hydrate` Parameter

This is the most powerful and confusing part of the API. It lets you embed related data in a single response instead of chaining calls.

```
?hydrate=stats(group=hitting,type=season,season=2024)
?hydrate=team
?hydrate=currentTeam
?hydrate=rosterEntries
?hydrate=stats(group=hitting,type=season),currentTeam
```

Multiple hydrations are comma-separated. You can nest parameters inside parentheses.

---

## Team Endpoints

### Get all teams
```
GET /teams?sportId=1&season=2025
```

### Get a single team
```
GET /teams/{teamId}
```

### Get a team's roster
```
GET /teams/{teamId}/roster?rosterType=activeRoster&season=2025
```
`rosterType` options: `activeRoster`, `40Man`, `fullRoster`, `depthChart`, `gameday`

### Get a team's stats
```
GET /teams/{teamId}/stats?group=hitting&stats=season&season=2024
```

### Known team IDs (partial list)
| Team | ID |
|---|---|
| Yankees | 147 |
| Red Sox | 111 |
| Dodgers | 119 |
| Giants | 137 |
| Cubs | 112 |
| Cardinals | 138 |
| Astros | 117 |
| Braves | 144 |
| Mets | 121 |
| Phillies | 143 |

Get all IDs with `GET /teams?sportId=1`.

---

## Schedule Endpoints

### Get games on a specific date
```
GET /schedule?sportId=1&date=2024-07-04
```

### Get games in a date range
```
GET /schedule?sportId=1&startDate=2024-07-01&endDate=2024-07-07&teamId=147
```

### Get a specific game (live feed)
```
GET /game/{gamePk}/feed/live
```
This is the richest endpoint in the API. Returns everything about a game: play-by-play, box score, current at-bat, pitch data, fielding positions, weather, attendance. The `gamePk` is the game ID you get from the schedule endpoint.

### Get just the box score
```
GET /game/{gamePk}/boxscore
```

### Get play-by-play
```
GET /game/{gamePk}/playByPlay
```

### Get line score
```
GET /game/{gamePk}/linescore
```

---

## Stats Endpoints

### League-wide stats (leaderboards)
```
GET /stats?stats=season&group=hitting&season=2024&playerPool=All&limit=50
```
`playerPool` options: `All`, `Qualified`, `Rookies`

Sort by a specific stat:
```
GET /stats?stats=season&group=hitting&season=2024&sortStat=homeRuns&order=desc&limit=25
```

### Stats leaders
```
GET /stats/leaders?leaderCategories=homeRuns&season=2024&sportId=1
```
`leaderCategories` can be a comma-separated list: `homeRuns,battingAverage,rbi`

---

## Standings
```
GET /standings?leagueId=103&season=2024    (AL = 103)
GET /standings?leagueId=104&season=2024    (NL = 104)
GET /standings?leagueId=103,104&season=2024
```

---

## Meta Endpoint

Useful for discovering valid values for parameters:
```
GET /meta?type=statTypes
GET /meta?type=statGroups
GET /meta?type=pitchTypes
GET /meta?type=positions
GET /meta?type=gameTypes
GET /meta?type=situationCodes    ← for splits
GET /meta?type=rosterTypes
```

Hit these when you're not sure what valid values are for a parameter. The `situationCodes` one is particularly useful — it lists all the split codes.

---

## Splits (Situational Stats)

This is how you get vs LHP, home/away, etc. Use the `team_stats` endpoint with `stats=statSplits` and `sitCodes`:

```
GET /teams/{teamId}/stats?group=hitting&stats=statSplits&season=2024&sitCodes=vl
```

Or for a player, use the `stats` endpoint:
```
GET /stats?stats=statSplits&group=hitting&season=2024&sitCodes=vl&personId=592450
```

### Common sitCodes
| Code | Meaning |
|---|---|
| `vl` | vs Left-handed pitchers |
| `vr` | vs Right-handed pitchers |
| `h` | Home |
| `a` | Away |
| `risp` | Runners in scoring position |
| `high` | High leverage |
| `low` | Low leverage |
| `ahead` | Ahead in count |
| `behind` | Behind in count |
| `last7` | Last 7 days |
| `last14` | Last 14 days |
| `last30` | Last 30 days |
| `beforeAllStar` | First half |
| `afterAllStar` | Second half |

Get the full list from `GET /meta?type=situationCodes`.

---

## Sport IDs
| Sport | ID |
|---|---|
| MLB | 1 |
| AAA | 11 |
| AA | 12 |
| High-A | 13 |
| Single-A | 14 |

Useful for minor league data — same endpoints work, just change the sportId.

---

## Typical Workflow

For your project, the flow to get a player and their stats looks like this:

1. **Find the player ID** — either hardcode known ones for testing, or fetch all players with `/sports/1/players?season=2025` and filter by name
2. **Get their bio** — `GET /people/{personId}`
3. **Get their stats** — `GET /people/{personId}/stats?stats=season&group=hitting&season=2025`
4. **Or do both in one call** — `GET /people/{personId}?hydrate=stats(group=hitting,type=season,season=2025)`

For splits:
1. Get player ID (same as above)
2. `GET /stats?stats=statSplits&group=hitting&season=2025&sitCodes=vl&personId={id}`

---

## What This API Does NOT Have

- Statcast data (pitch spin, exit velocity, launch angle) — that's Baseball Savant
- xStats (xBA, xSLG, xwOBA) — Baseball Savant
- Pitch-by-pitch Statcast metrics — Baseball Savant
- FIP, WAR, wRC+ — these are calculated stats, not in any raw API. You compute them yourself.

For anything Statcast or advanced sabermetric, you need Baseball Savant's CSV endpoint.

---

## Baseball Savant CSV Endpoint

While we're here — this is the other half of your data pipeline:

```
https://baseballsavant.mlb.com/statcast_search/csv
  ?player_type=batter
  &pitcherid=592450        (use hitterid= for batters)
  &season=2024
  &game_type=R             (R = regular season)
```

Returns a CSV of every pitch for that player in that season. Columns include pitch type, velocity, spin rate, break, plate location, launch speed, launch angle, hit distance, events, and about 80 other fields.

Player IDs are the same as the MLBAM IDs from the stats API, which is convenient.

---

## Known Gotchas

- **Response shape is inconsistent** — some endpoints wrap data in `people[]`, some in `stats[]`, some nest multiple levels deep. Always inspect the raw JSON before writing parse code.
- **Empty stats vs missing stats** — if a player has no data for a season, you may get an empty array or a 200 with no content rather than an error.
- **`season` param is often required** — many endpoints return nothing without it, even if it seems optional.
- **The `hydrate` syntax is fiddly** — nested params use commas inside parens, and the exact syntax varies. Expect trial and error.
- **Statcast coverage starts 2015** — anything before that is pre-Statcast and won't have pitch tracking data.
- **Player IDs are consistent across systems** — the same MLBAM ID works in statsapi.mlb.com AND baseballsavant.mlb.com. Baseball Reference has a crosswalk table if you ever need to match with their IDs.
