package validator

import "gopkg.in/go-playground/validator.v9"

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

type Validator struct {
	validator *validator.Validate
}

func customValidation(vd validator.StructLevel) {

}
func (v *Validator) Validate(i interface{}) error {
	//v.validator.RegisterStructValidation()

	return v.validator.Struct(i)
}
