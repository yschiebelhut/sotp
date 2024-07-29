/*
Copyright © 2024 Yannik Schiebelhut <yannik.schiebelhut@gmail.com>
*/
package otp

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func GenerateOTP(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())
}
