package util

import (
	"github.com/holanda-lucas/gotasks-back/database"
	"github.com/robfig/cron/v3"
)

func StartTokenDisposer() {
	c := cron.New()

	c.AddFunc("@hourly", database.DisposeExpiredTokens)

	c.Start()
}