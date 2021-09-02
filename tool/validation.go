package tool

import (
	"gopkg.in/asaskevich/govalidator.v4"
)

// ValidateRequest incoming request
func ValidateRequest(data interface{}) error {
	if valid, err := govalidator.ValidateStruct(data); valid == false {
		return err
	}

	return nil
}
