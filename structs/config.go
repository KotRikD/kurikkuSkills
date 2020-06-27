package structs

// Config application
type Config struct {
	SQLDsn        string `description:"SQL url to connect MySQL db"`
	RedisAddress  string `description:"Redis address"`
	RedisPassword string `description:"Redis password"`
	RedisDB       int    `description:"Redis Database Index"`
	BeatmapsPath  string `description:"It should look like: /var/lets/.data/maps/"`
	IPPort        string `description:"where i should be located"`
	NeedSemki     bool   `description:"start as re-calculator"`
}
