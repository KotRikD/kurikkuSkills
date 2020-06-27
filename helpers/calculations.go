package helpers

import (
	"math"
	"strconv"

	"github.com/KotRikD/kurikkuSkills/osuSkills"
	"github.com/KotRikD/kurikkuSkills/structs"
)

// CalculateScore calculate and update single score
func CalculateScore(score structs.RedisCalc) error {
	skills, err := osuSkills.CalculateSkills(Config.BeatmapsPath+strconv.Itoa(score.MapID)+".osu", score.Mods)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`insert into scores_skills values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		score.ScoreID, score.Mods, skills.Stamina, skills.Tenacity, skills.Agility, skills.Precision, skills.Reading, skills.Memory, skills.Accuracy, skills.Reaction, skills.SliderSpinners, skills.Circles)
	if err != nil {
		return err
	}
	return nil
}

// CalculateScoreByValues calculate values meh
func CalculateScoreByValues(scoreID int, mapID int, mods int) error {
	skills, err := osuSkills.CalculateSkills(Config.BeatmapsPath+strconv.Itoa(mapID)+".osu", mods)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`insert into scores_skills values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		scoreID, mods, skills.Stamina, skills.Tenacity, skills.Agility, skills.Precision, skills.Reading, skills.Memory, skills.Accuracy, skills.Reaction, skills.SliderSpinners, skills.Circles)
	if err != nil {
		return err
	}
	return nil
}

// ReCalculateSkills Recalculate All users scores
func ReCalculateSkills(UserID int) error {
	rows, err := DB.Query(`SELECT 
		scores.id, scores.userid, scores.completed, scores.beatmap_md5, scores.mods, scores.max_combo, scores.misses_count, scores.accuracy AS accuracy,
		scores_skills.stamina, scores_skills.tenacity, scores_skills.agility, scores_skills.pprecision, scores_skills.reading, scores_skills.memory, scores_skills.accuracy, 
		scores_skills.reaction, 
		scores_skills.sliderspinners, scores_skills.circles
		beatmaps.max_combo, beatmaps.hit_length
	FROM scores
	RIGHT JOIN scores_skills ON scores_skills.id = scores.id
	RIGHT JOIN beatmaps ON beatmaps.beatmap_md5 = scores.beatmap_md5
	WHERE scores.completed = 3 AND scores.play_mode = 0 AND scores.userid = ?`, UserID)

	if err != nil {
		return err
	}

	var gainedSkills []structs.Skills
	for rows.Next() {
		var scorepart structs.Score
		var scoreskillspart structs.SkillsScores
		var beatmappart structs.Beatmap

		err := rows.Scan(&scorepart.ID, &scorepart.UserID, &scorepart.Completed, &scorepart.BeatmapMD5, &scorepart.Mods, &scorepart.MaxCombo, &scorepart.CountMiss, &scorepart.Accuracy,
			&scoreskillspart.Stamina, &scoreskillspart.Tenacity, &scoreskillspart.Agility, &scoreskillspart.Precision, &scoreskillspart.Reading, &scoreskillspart.Memory,
			&scoreskillspart.Accuracy, &scoreskillspart.Reaction, &scoreskillspart.SliderSpinners, &scoreskillspart.Circles,
			&beatmappart.MaxCombo, &beatmappart.HitLength,
		)
		if err != nil {
			return err
		}

		newSkills := structs.Skills{
			Stamina:   scoreskillspart.Stamina,
			Tenacity:  scoreskillspart.Tenacity,
			Agility:   scoreskillspart.Agility,
			Precision: scoreskillspart.Precision,
			Reading:   scoreskillspart.Reading,
			Memory:    scoreskillspart.Memory,
			Accuracy:  scoreskillspart.Accuracy,
			Reaction:  scoreskillspart.Reaction,
		}
		if scorepart.MaxCombo == beatmappart.MaxCombo {
			gainedSkills = append(gainedSkills, newSkills)
		} else {
			combomult := math.Min(math.Pow(float64(scorepart.MaxCombo), 0.8)/math.Pow(float64(beatmappart.MaxCombo), 0.8), 1.)

			newSkills.Stamina *= math.Pow(0.97, float64(scorepart.CountMiss))
			newSkills.Tenacity *= math.Pow(0.97, float64(scorepart.CountMiss))
			newSkills.Agility *= math.Pow(0.97, float64(scorepart.CountMiss)) * combomult
			newSkills.Precision *= math.Pow(0.97, float64(scorepart.CountMiss)) * combomult
			newSkills.Reading *= math.Pow(0.97, float64(scorepart.CountMiss)) * combomult
			newSkills.Memory *= math.Pow(0.97, float64(scorepart.CountMiss)) * combomult
			newSkills.Reaction *= math.Pow(0.97, float64(scorepart.CountMiss)) * combomult
		}

		if scoreskillspart.Circles == 0 {
			newSkills.Accuracy = 0
		} else {
			cacc := scorepart.Accuracy - (float64(scoreskillspart.SliderSpinners/scoreskillspart.Circles))*(1-scorepart.Accuracy)
			if cacc < 0.36 {
				newSkills.Accuracy = 0
			}

			// double v = spline3dcalc(spline, cacc, bData[index].circles, GetODWithMods(plays[index].mods, bData[index].od));
			// gained[index].skills.accuracy /= pow(v, 1.3);
			// idk what is spline3dcalc, not enough c++ skill
		}

		gainedSkills = append(gainedSkills, newSkills)

		// now we need get weighted value of these calculations of any skill...
		var st []float64
		var ten []float64
		var agi []float64
		var pre []float64
		var read []float64
		var mem []float64
		var acc []float64
		var reac []float64
		for _, x := range gainedSkills {
			st = append(st, x.Stamina)
			ten = append(ten, x.Tenacity)
			agi = append(agi, x.Agility)
			pre = append(pre, x.Precision)
			read = append(read, x.Reading)
			mem = append(mem, x.Memory)
			acc = append(acc, x.Accuracy)
			reac = append(reac, x.Reaction)
		}

		valStamina := GetWeightedValue(st, 0.95)
		valTenacity := GetWeightedValue(ten, 0.95)
		valAgility := GetWeightedValue(agi, 0.95)
		valPrecision := GetWeightedValue(pre, 0.95)
		valReading := GetWeightedValue(read, 0.95)
		valMemory := GetWeightedValue(mem, 0.95)
		valAccuracy := GetWeightedValue(acc, 0.95)
		valReaction := GetWeightedValue(reac, 0.95)

		DB.Exec(`UPDATE users_stats
			SET skill_stamina = ?, skill_tenacity = ?, skill_agility = ?, skill_precision = ?, skill_reading = ?, skill_memory = ?, skill_accuracy = ?, skill_reaction = ? WHERE id = ?`,
			valStamina, valTenacity, valAgility, valPrecision, valReading, valMemory, valAccuracy, valReaction, UserID)
	}
	return nil
}
