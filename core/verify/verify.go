package verify

import (
	"fmt"
	"github.com/yddeng/webhook/conf"
)

func VerifyAccess(ip, token string) error {
	access := conf.GetConfig()

	if access.AccessToken != "" && token != access.AccessToken {
		return fmt.Errorf("token is failed\n")
	}

	if len(access.AccessIP) != 0 {
		for _, ip_ := range access.AccessIP {
			if ip_ != "" && ip == ip_ {
				return nil
			}
		}
		return fmt.Errorf("ip is failed\n")
	}

	return nil
}
