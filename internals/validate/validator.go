package validate

import "github.com/go-playground/validator/v10"

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func Struct(s interface{}) error {
	return GetValidator().Struct(s)
}

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		// registering alias so we can see the differences between
		// map key, value validation errors
		validate.RegisterAlias("tenMax", "max=10")
		validate.RegisterAlias("twentyMax", "max=20")
	}
	return validate
}
