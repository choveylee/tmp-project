package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/choveylee/tutil"
	"github.com/gin-gonic/gin"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	"dev.choveylee.top/elder-care-backend/internal/lib"
)

func CreateImageAdmin(ctx context.Context, userId string, c *gin.Context, file *multipart.FileHeader, filename, fileExt string) (*data.CreateImageAdminRespData, *terror.Terror) {
	srcFile, err := file.Open()
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Create image admin (user id: %s, filename: %s, file ext: %s) err (open upload file %v).",
			userId, filename, fileExt, err)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("upload file"), constant.ErrorCodeRequestParamInvalid, errMsg)

		return nil, errx
	}

	defer srcFile.Close()

	rawContent, err := io.ReadAll(srcFile)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Create image admin (user id: %s, filename: %s, file ext: %s) err (read all %v).",
			userId, filename, fileExt, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return nil, errx
	}

	objectName := fmt.Sprintf("image/%s.%s", tutil.NewOid().String(), fileExt)

	errx := lib.PutMinioObject(ctx, objectName, rawContent)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create image admin (user id: %s, filename: %s, file ext: %s, object name: %s) err (put minio object %v).",
			userId, filename, fileExt, objectName, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	fileUrl, errx := lib.GenMinioPresignedUrlEx(ctx, objectName, time.Hour)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create image admin (user id: %s, filename: %s, file ext: %s, object name: %s) err (gen minio presigned url ex %v).",
			userId, filename, fileExt, objectName, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	createImageRespData := &data.CreateImageAdminRespData{
		FileKey: objectName,
		FileUrl: fileUrl,
	}

	return createImageRespData, nil
}
