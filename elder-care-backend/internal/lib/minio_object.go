/**
 * @Author: lidonglin
 * @Description:
 * @File:  minio_object.go
 * @Version: 1.0.0
 * @Date: 2025/11/6 22:38:52
 */

package lib

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/s3utils"
	"github.com/minio/minio-go/v7/pkg/signer"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

func PutMinioObject(ctx context.Context, objectName string, data []byte) *terror.Terror {
	reader := bytes.NewReader(data)

	info, err := minioClient.PutObject(ctx, minioBucket, objectName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Put minio object (object name: %s) err (put object %v)",
			objectName, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return errx
	}

	fmt.Printf("内存数据上传成功: %s (大小: %d bytes)\n", objectName, info.Size)

	return nil
}

func PutMinioExcel(ctx context.Context, objectName string, reader *bytes.Buffer) error {
	info, err := minioClient.PutObject(ctx, minioBucket, objectName, reader, int64(reader.Len()), minio.PutObjectOptions{
		ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Put minio excel (object name: %s) err (put object %v)",
			objectName, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return errx
	}

	fmt.Printf("内存数据上传成功: %s (大小: %d bytes)\n", objectName, info.Size)

	return nil
}

func GetMinioObject(ctx context.Context, objectName string) ([]byte, *terror.Terror) {
	object, err := minioClient.GetObject(ctx, minioBucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Get minio object (object name: %s) err (get object %v)",
			objectName, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return nil, errx
	}
	defer object.Close()

	// 3. 读取对象内容
	objectBytes, err := io.ReadAll(object)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Get minio object (object name: %s) err (read all %v)",
			objectName, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return nil, errx
	}

	return objectBytes, nil
}

func GenMinioPresignedUrl(ctx context.Context, objectName string, expiredTime time.Duration) (string, *terror.Terror) {
	presignedURL, err := minioClient.PresignedGetObject(ctx, minioBucket, objectName, expiredTime, nil)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Gen minio presigned url (object name: %s, expired time: %v) err (presigned get object %v)",
			objectName, expiredTime, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return "", errx
	}

	return presignedURL.String(), nil
}

// GenMinioPresignedUrlEx 使用 MINIO_ENDPOINT2 作为下载链接的 host（path-style：/{bucket}/{object}），
// 凭证与 bucket 与 minioClient 一致，签名算法为 AWS Signature Version 4（与 S3 预签名 GET 一致）。
func GenMinioPresignedUrlEx(ctx context.Context, objectName string, expiredTime time.Duration) (string, *terror.Terror) {
	err := s3utils.CheckValidObjectName(objectName)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Gen minio presigned url ex (object name: %s, expired time: %v) err (check valid object name %v)",
			objectName, expiredTime, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return "", errx
	}

	expireSeconds := int64(expiredTime / time.Second)

	location, err := minioClient.GetBucketLocation(ctx, minioBucket)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Gen minio presigned url ex (object name: %s, expired time: %v) err (get bucket location %v)",
			objectName, expiredTime, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return "", errx
	}

	if location == "" {
		location = "us-east-1"
	}

	pathUrl := minioEndpoint2 + "/" + minioBucket + "/" + s3utils.EncodePath(objectName)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, pathUrl, nil)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Gen minio presigned url ex (object name: %s, expired time: %v) err (new request %v)",
			objectName, expiredTime, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return "", errx
	}

	signedUrl := signer.PreSignV4(*request, minioAccessKey, minioSecretKey, "", location, expireSeconds)

	return signedUrl.URL.String(), nil
}
