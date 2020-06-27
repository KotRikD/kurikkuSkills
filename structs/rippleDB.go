package structs

// Beatmap beatmap struct db
type Beatmap struct {
	BeatmapID          int     `json:"beatmap_id"`
	BeatmapsetID       int     `json:"beatmapset_id"`
	BeatmapMD5         string  `json:"beatmap_md5"`
	SongName           string  `json:"song_name"`
	AR                 float32 `json:"ar"`
	OD                 float32 `json:"od"`
	DIFFSTD            float64 `json:"std"`
	DIIFFTaiko         float64 `json:"taiko"`
	DIFFCTB            float64 `json:"ctb"`
	DIFFMania          float64 `json:"mania"`
	MaxCombo           int     `json:"max_combo"`
	HitLength          int     `json:"hit_length"`
	Ranked             int     `json:"ranked"`
	RankedStatusFrozen int     `json:"ranked_status_frozen"`
	LatestUpdate       int64   `json:"latest_update"`
}

// Score score struct db
type Score struct {
	ID         int    `json:"id"`
	BeatmapMD5 string `json:"beatmap_md5"`
	UserID     int
	Score      int64   `json:"score"`
	MaxCombo   int     `json:"max_combo"`
	FullCombo  bool    `json:"full_combo"`
	Mods       int     `json:"mods"`
	Count300   int     `json:"count_300"`
	Count100   int     `json:"count_100"`
	Count50    int     `json:"count_50"`
	CountKatu  int     `json:"count_katu"`
	CountGeki  int     `json:"count_geki"`
	CountMiss  int     `json:"count_miss"`
	Time       int64   `json:"time"`
	PlayMode   int     `json:"play_mode"`
	Completed  int     `json:"completed"`
	Accuracy   float64 `json:"accuracy"`
	PP         float32 `json:"pp"`
}

// SkillsScores struct for DB
type SkillsScores struct {
	ID             int
	Mods           int
	Stamina        float64
	Tenacity       float64
	Agility        float64
	Precision      float64
	Reading        float64
	Memory         float64
	Accuracy       float64
	Reaction       float64
	SliderSpinners int
	Circles        int
}

// UsersSkills mini-struct for Ripple DB users
type UsersSkills struct {
	ID        int
	Stamina   float64
	Tenacity  float64
	Agility   float64
	Precision float64
	Reading   float64
	Memory    float64
	Accuracy  float64
	Reaction  float64
}
