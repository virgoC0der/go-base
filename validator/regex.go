package validator

import "regexp"

var (
	PhoneRegex    = regexp.MustCompile(`^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`)
	UsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_/-]{4,16}$`)
	PasswordRegex = regexp.MustCompile(`^[a-zA-Z0-9_/-@!#%&*()=+]{6,16}$`)
	EmailRegex    = regexp.MustCompile(`^[a-zA-Z0-9_/-]+@[a-zA-Z0-9_/-]+(\.[a-zA-Z0-9_/-]+)+$`)
)
