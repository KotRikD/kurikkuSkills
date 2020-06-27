package api

import (
	"log"
	"net/http"

	"github.com/KotRikD/kurikkuSkills/structs"
	"github.com/jmoiron/sqlx"
	"gopkg.in/macaron.v1"
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

// InitServer initing server
func InitServer() {
	m := macaron.Classic()
	m.Get("/", myHandler)

	log.Println("[Macaroon <3] Macaroon started")
	log.Println(http.ListenAndServe(Config.IPPort, m))
}

func myHandler(ctx *macaron.Context) string {
	return "hi, unknown dude"
}
