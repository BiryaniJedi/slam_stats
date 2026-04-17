-- +goose Up
CREATE TABLE players (
  id INTEGER PRIMARY KEY,
  fname TEXT,
  lname TEXT,
  bdate TEXT,
  height_in INTEGER,
  weight_lbs FLOAT,
  pos_code TEXT,
  pos_name TEXT,
  pos_type TEXT,
  pos_abbrv TEXT,
  bat_side_code TEXT,
  bat_side_desc TEXT,
  pitch_hand_code TEXT,
  pitch_hand_desc TEXT,
  cur_team_id INTEGER,
  cur_team_name TEXT,
  cur_team_link TEXT
);

-- +goose Down
DROP TABLE players;
