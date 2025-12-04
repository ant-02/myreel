package upyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"myreel/config"
	"time"
)

type UpyunToken struct {
	Policy        string
	Authorization string
	Bucket        string
}

func GeneratePolicyAndSignature(uid int64, saveKey string, notifyPath string, extParam any) (*UpyunToken, error) {
	expire := time.Now().Add(config.Upyun.Expiration * time.Minute).Unix()

	// 上传策略
	p := map[string]interface{}{
		"bucket":               config.Upyun.Bucket,
		"save-key":             saveKey,
		"expiration":           expire,
		"content-length-range": fmt.Sprintf("0,%d", config.Upyun.MaxSize),
	}

	if notifyPath != "" {
		p["notify-url"] = notifyPath
	}

	if extParam != nil {
		p["ext-param"] = extParam
	}

	jsonBytes, _ := json.Marshal(p)
	policy := base64.StdEncoding.EncodeToString(jsonBytes)

	mac := hmac.New(sha1.New, []byte(config.Upyun.Password))
	mac.Write([]byte(fmt.Sprintf("POST&/%s&%s", config.Upyun.Bucket, policy)))

	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return &UpyunToken{
		Policy:        policy,
		Authorization: fmt.Sprintf("UPYUN %s:%s", config.Upyun.Operator, signature),
		Bucket:        config.Upyun.Bucket,
	}, nil
}
