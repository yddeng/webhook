package core

import "github.com/yddeng/webhook/conf"

func VerifyAccess(ip, token string) bool {
	access := conf.GetConfig().Access

	if access.AccessToken != "" && token != access.AccessToken {
		return false
	}

	if len(access.AccessIP) != 0 {
		for _, ip_ := range access.AccessIP {
			if ip_ != "" && ip == ip_ {
				return true
			}
		}
		return false
	}

	return true
}

func VerifyCommand(cmd string) bool {
	cmds := conf.GetConfig().Command

	for _, c := range cmds {
		if c == cmd {
			return true
		}
	}
	return false
}
