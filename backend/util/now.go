package util

import (
	"time"

	"github.com/jorgejr568/salary-go-api/cfg"
)

func Now() time.Time {
	loc, _ := time.LoadLocation(cfg.CfgDefaultTimezone())
	return time.Now().In(loc)
}
