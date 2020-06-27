package main

import (
	"fmt"

	"github.com/KotRikD/kurikkuSkills/helpers"
	"github.com/KotRikD/kurikkuSkills/structs"
)

// SemkiRecalculator э ну чо семки то есть?
func SemkiRecalculator() error {
	rows, err := DB.Query(`SELECT
		scores.id, scores.beatmap_md5, scores.mods, scores.play_mode, scores.userid,
		beatmaps.beatmap_id
	FROM scores
	RIGHT JOIN beatmaps ON beatmaps.beatmap_md5 = scores.beatmap_md5
	RIGHT JOIN users ON users.id = scores.userid
	WHERE scores.completed = 3 AND scores.play_mode = 0 AND (users.privileges & 1 = 1)`)

	if err != nil {
		return err
	}

	var userIDs []int
	for rows.Next() {
		var scoreDB structs.Score
		var beatmapID int
		err := rows.Scan(&scoreDB.ID, &scoreDB.BeatmapMD5, &scoreDB.Mods, &scoreDB.PlayMode, &scoreDB.UserID,
			&beatmapID)
		if err != nil {
			return err
		}

		fmt.Println("[I] Calculating skills for", scoreDB.ID)

		exists, _ := helpers.ExistsInArray(scoreDB.UserID, userIDs)
		if exists != true {
			userIDs = append(userIDs, scoreDB.UserID)
		}
		err = helpers.CalculateScoreByValues(scoreDB.ID, beatmapID, scoreDB.Mods)
		if err != nil {
			fmt.Println("[I] Skills calculation failed", scoreDB.ID, ":", err)
			continue
		}
	}

	fmt.Println("[I] Started calculations for user skills")
	for _, userID := range userIDs {
		err := helpers.ReCalculateSkills(userID)
		if err != nil {
			fmt.Println("[I] Skills of", userID, "can't be calculated", err)
			continue
		}

		fmt.Println("[I] Skills of", userID, "has been re-calculated")
	}

	return nil
}
