// Package lib initializes shared library integrations required by the service.
package lib

import (
	"context"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
)

// InitLib initializes shared library dependencies required by the service.
func InitLib(ctx context.Context) *terror.Terror {
	errx := initMinio(ctx)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Init lib err (init minio %v)",
			errx.Error())
		errx.AttachErrMsg(errMsg)

		return errx
	}

	errx = initJwt(ctx)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Init lib err (init jwt %v)",
			errx.Error())
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}
