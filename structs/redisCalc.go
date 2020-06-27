package structs

// RedisCalc struct for handling redis calculations
type RedisCalc struct {
	ScoreID int `json:"score_id"`
	MapID   int `json:"map_id"`
	UserID  int `json:"user_id"`
	Mods    int `json:"mods"`
}
