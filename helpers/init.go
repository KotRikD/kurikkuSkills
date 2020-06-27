package helpers

import (
	"github.com/KotRikD/kurikkuSkills/structs"
	"github.com/jmoiron/sqlx"
)

var (
	// Config a
	Config structs.Config
	// DB a
	DB *sqlx.DB
)
