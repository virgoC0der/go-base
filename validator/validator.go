package validator

import "github.com/go-playground/validator/v10"

var valid *validator.Validate

func init() {
	valid = validator.New()
	RegisterValidator(valid)
}

func RegisterValidator(v *validator.Validate) {
	v.RegisterValidation("phone", validPhone)
	v.RegisterValidation("email", validEmail)
	v.RegisterValidation("password", validPassword)
	v.RegisterValidation("username", validUsername)
	v.RegisterValidation("timestamp", validTimestamp)
}

func validPhone(fl validator.FieldLevel) bool {
	return PhoneRegex.MatchString(fl.Field().String())
}

func validUsername(fl validator.FieldLevel) bool {
	return UsernameRegex.MatchString(fl.Field().String())
}

func validPassword(fl validator.FieldLevel) bool {
	return PasswordRegex.MatchString(fl.Field().String())
}

func validEmail(fl validator.FieldLevel) bool {
	return EmailRegex.MatchString(fl.Field().String())
}

func validTimestamp(fl validator.FieldLevel) bool {
	return !(0 <= fl.Field().Int() && fl.Field().Int() <= 253370736000)
}
