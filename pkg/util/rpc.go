package util

import (
	"myreel/kitex_gen/model"
	"myreel/pkg/errno"
)

func IsSuccess(baseResp *model.BaseResp) bool {
	return baseResp.Code == errno.SuccessCode
}
