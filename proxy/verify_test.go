package proxy

import (
	"fmt"
	conf "github.com/yddeng/webhook/configs/proxy"
	"testing"
)

func TestVerifyAccess(t *testing.T) {
	conf.LoadConfig("../configs/proxy/config.toml")
	fmt.Println(VerifyAccess("10.128.2.123", "ff"))
	fmt.Println(VerifyAccess("10.128.2.123", "token"))
	fmt.Println(VerifyAccess("10.128.2.1", "token"))

}
