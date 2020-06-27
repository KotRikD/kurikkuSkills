package main

import (
	"encoding/json"
	"fmt"

	"github.com/KotRikD/kurikkuSkills/helpers"
	"github.com/KotRikD/kurikkuSkills/structs"
)

// StartListenRedis zalupa
func StartListenRedis() {
	pubsub, _ := RD.Subscribe("kr:calc1")

	//Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				fmt.Println(err)
				continue
			}
			var jsonEvent structs.RedisCalc

			err = json.Unmarshal([]byte(msg.Payload), &jsonEvent)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("[I] Found new user score sendedd by LETS score_id:", jsonEvent.ScoreID, "user_id:", jsonEvent.UserID, "mods:", jsonEvent.Mods)
			err = helpers.CalculateScore(jsonEvent)
			if err != nil {
				fmt.Println(err, "<-- calculation result")
				continue
			}
			fmt.Println("[I] Score calculated, starting re-calculate user")

			err = helpers.ReCalculateSkills(jsonEvent.UserID)
			if err != nil {
				fmt.Println("[I] User not calculated user_id:", jsonEvent.UserID)
				continue
			}
			fmt.Println("[I] User re-calculated")
		}
	}()
}
