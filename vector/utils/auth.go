package utils

import (
	"net/http"

	"github.com/volcengine/volc-sdk-golang/base"
)

func SignRequest(req *http.Request, ak string, sk string) *http.Request {
	credential := base.Credentials{
					AccessKeyID:     ak,
					SecretAccessKey: sk,
					Service:         "vikingdb",
					Region:          "cn-north-1",
	}
	req = credential.Sign(req)
	return req
}