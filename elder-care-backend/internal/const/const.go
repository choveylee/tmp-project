// Package constant defines shared constants used across the service.
package constant

import (
	"regexp"

	"github.com/dlclark/regexp2"
)

const (
	RunModeDebug   = "debug"
	RunModeRelease = "release"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

const (
	MaxPageNum  = 100
	MaxPageSize = 100
)

const (
	SortSeqAsc  = "asc"
	SortSeqDesc = "des"
)

var (
	ChnMobileReg = regexp.MustCompile(`^((\+86)|(86))?(1)\d{10}$`)
	ChnPhoneReg  = regexp.MustCompile(`^[+]?[0-9]{0,3}-?(13|14|15|16|17|18|19)[0-9]{9}$|^0\d{2,3}-\d{7,8}$|^0\d{2,3}-\d{7,8}-\d{1,4}$`)

	PasswordReg = regexp2.MustCompile(`^(?=.*\d)(?=.*[a-zA-Z])(?=.*[^\da-zA-Z\s]).{1,9}$`, 0)
)
