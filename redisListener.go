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

			fmt.Println(msg.Payload)

			err = json.Unmarshal([]byte(msg.Payload), &jsonEvent)
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = helpers.CalculateScore(jsonEvent)
			if err != nil {
				fmt.Println(err, "<-- calculation result")
			}
			fmt.Println(jsonEvent)

		}
	}()
}
