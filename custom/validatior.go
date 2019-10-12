package custom

// import "gopkg.in/go-playground/validator.v9"
import (
	"github.com/asaskevich/govalidator"
)

// Validator custom validator
type Validator struct{}

// Validate implementing echo validate interface
func (v *Validator) Validate(i interface{}) error {
	ok, err := govalidator.ValidateStruct(i)
	if ok {
		return nil
	}
	if e, ok := err.(govalidator.Errors); ok {
		errors := e.Errors()
		return errors[0]
	}
	return err
}
