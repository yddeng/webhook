package verify

import (
	"fmt"
	"github.com/yddeng/webhook/conf"
	"testing"
)

func TestVerifyAccess(t *testing.T) {
	conf.LoadConfig("../../config.toml")
	fmt.Println(VerifyAccess("10.128.2.123", "ff"))
	fmt.Println(VerifyAccess("10.128.2.123", "token"))
	fmt.Println(VerifyAccess("10.128.2.1", "token"))

}
