// Package service implements service-layer logic exposed through HTTP handlers and scheduled jobs.
package service

import (
	"context"

	"github.com/choveylee/tcfg"
	"github.com/choveylee/terror"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

var (
	runMode string
)

// InitService initializes service-layer dependencies.
func InitService(ctx context.Context) *terror.Terror {
	runMode = tcfg.DefaultString(tcfg.LocalKey("RUN_MODE"), constant.RunModeDebug)

	return nil
}
