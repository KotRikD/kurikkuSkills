package main

import (
	"fmt"

	"github.com/KotRikD/kurikkuSkills/api"
	"github.com/KotRikD/kurikkuSkills/helpers"
	"github.com/KotRikD/kurikkuSkills/osuSkills"
	"github.com/KotRikD/kurikkuSkills/structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/thehowl/conf"
	"gopkg.in/redis.v5"
)

var (
	// Config application
	Config structs.Config
	// RD workable variable for Redis
	RD *redis.Client
	// DB variable for sqlx
	DB *sqlx.DB
)

func main() {
	osuSkills.LoadVars()

	err := conf.Load(&Config, "application.conf")
	switch err {
	case nil:
		// carry on
	case conf.ErrNoFile:
		conf.Export(Config, "application.conf")
		fmt.Println("The configuration file was not found. We created one for you.")
		return
	default:
		panic(err)
	}

	DB, err = sqlx.Open("mysql", Config.SQLDsn+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	RD = redis.NewClient(&redis.Options{
		Addr:     Config.RedisAddress,
		Password: Config.RedisPassword,
		DB:       Config.RedisDB,
	})

	_, err = RD.Ping().Result()
	if err != nil {
		panic(err)
	}

	if Config.NeedSemki {
		err := SemkiRecalculator()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	api.Config = Config
	api.DB = DB
	api.RD = RD
	helpers.DB = DB
	helpers.Config = Config

	StartListenRedis()
	api.InitServer()
	// skills, err := osuSkills.CalculateSkills("f:/osu!new/Songs/1031162 Parry Gripp - Guinea Pig Olympics/Parry Gripp - Guinea Pig Olympics (Sotarks) [Expert].osu", 16)
	// if err != nil {
	// 	fmt.Println("aah?")
	// 	return
	// }
	// fmt.Println(skills.Accuracy)
}
