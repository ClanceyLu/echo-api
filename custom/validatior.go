package custom

// import "gopkg.in/go-playground/validator.v9"
import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

// Validator custom validator
type Validator struct{}

// Validate implementing echo validate interface
func (v *Validator) Validate(i interface{}) error {
	_, err := govalidator.ValidateStruct(i)
	if errs, ok := err.(govalidator.Errors); ok {
		return fmt.Errorf("%s", errs[0])
	}
	return err
}
