/**
 * @Author: lidonglin
 * @Description:
 * @File:  minio.go
 * @Version: 1.0.0
 * @Date: 2025/11/6 22:38:45
 */

package lib

import (
	"context"
	"time"

	"github.com/choveylee/tcfg"
	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client

	minioEndpoint2 string

	minioBucket string

	minioAccessKey string
	minioSecretKey string
)

func initMinio(ctx context.Context) *terror.Terror {
	minioEndpoint := tcfg.DefaultString(tcfg.LocalKey("MINIO_ENDPOINT"), "")
	if minioEndpoint == "" {
		errMsg := tlog.E(ctx).Msgf("init minio err (minio endpoint illegal).")

		errx := terror.NewRawTerror(ctx, terror.ErrConfInvalid("minio endpoint"), errMsg)

		return errx
	}

	minioEndpoint2 = tcfg.DefaultString(tcfg.LocalKey("MINIO_ENDPOINT2"), "")
	if minioEndpoint2 == "" {
		errMsg := tlog.E(ctx).Msgf("init minio err (minio endpoint2 illegal).")

		errx := terror.NewRawTerror(ctx, terror.ErrConfInvalid("minio endpoint2"), errMsg)

		return errx
	}

	minioAccessKey = tcfg.DefaultString(tcfg.LocalKey("MINIO_ACCESS_KEY"), "")
	if minioAccessKey == "" {
		errMsg := tlog.E(ctx).Msgf("init minio err (minio access key illegal).")

		errx := terror.NewRawTerror(ctx, terror.ErrConfInvalid("minio access key"), errMsg)

		return errx
	}

	minioSecretKey = tcfg.DefaultString(tcfg.LocalKey("MINIO_SECRET_KEY"), "")
	if minioSecretKey == "" {
		errMsg := tlog.E(ctx).Msgf("init minio err (minio secret key illegal).")

		errx := terror.NewRawTerror(ctx, terror.ErrConfInvalid("minio secret key"), errMsg)

		return errx
	}

	minioBucket = tcfg.DefaultString(tcfg.LocalKey("MINIO_BUCKET"), "")
	if minioBucket == "" {
		errMsg := tlog.E(ctx).Msgf("init minio err (minio bucket illegal).")

		errx := terror.NewRawTerror(ctx, terror.ErrConfInvalid("minio bucket"), errMsg)

		return errx
	}

	minioUseSsl := tcfg.DefaultBool(tcfg.LocalKey("MINIO_USE_SSL"), false)

	var err error

	minioClient, err = minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: minioUseSsl,
	})
	if err != nil {
		errMsg := tlog.E(ctx).Msgf("init minio err (%s, %s, %s) err (new minio %v).",
			minioEndpoint, minioAccessKey, minioSecretKey, err)

		errx := terror.NewRawTerror(ctx, terror.ErrSvcAbnormal("minio"), errMsg)

		return errx
	}

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = minioClient.ListBuckets(ctx)
	if err != nil {
		errMsg := tlog.E(ctx).Msgf("init minio err (%s, %s, %s) err (list bucket %v).",
			minioEndpoint, minioAccessKey, minioSecretKey, err)

		errx := terror.NewRawTerror(ctx, terror.ErrSvcAbnormal("minio"), errMsg)

		return errx
	}

	return nil
}
