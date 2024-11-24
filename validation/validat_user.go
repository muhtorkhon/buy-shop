package validation

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators(v *validator.Validate) *validator.Validate {
	v.RegisterValidation("phone", ValidatePhoneNumber)
	//v.RegisterValidation("password", ValidatePassword)
	//v.RegisterValidation("email", ValidateEmail)
	return v
}

func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	var phoneRegex = `^(?:\+998|998)?[0-9]{9}$`

	matched, err := regexp.MatchString(phoneRegex, phone)
	if err != nil {
		log.Println("Error in phone validation regex:", err)
		return false
	}
	return matched
}

// func ValidatePassword(fl validator.FieldLevel) bool {
// 	password := fl.Field().String()
// 	log.Println("Validating password:======", password)

// 	var passwordRegex = `^(?=.*[a-zA-Z])(?=.*[0-9]).{8,}$`
// 	matched, err := regexp.MatchString(passwordRegex, password)
// 	if err != nil {
// 		log.Println("Error in password validation regex:", err)
// 		return false
// 	}
// 	return matched
// }

// func ValidateEmail(fl validator.FieldLevel) bool {
// 	email := fl.Field().String()
// 	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
// 	matched, err := regexp.MatchString(emailRegex, email)
// 	if err != nil {
// 		log.Println("invalid email format")
// 		return false
// 	}
// 	return matched
// }

// import (
// 	"regexp"
// 	"fmt"
// )

// func ValidatePassword(password string) error {
// 	regex := regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`)
// 	if !regex.MatchString(password) {
// 		return fmt.Errorf("password must be at least 8 characters long and contain at least one letter and one number")
// 	}
// 	return nil
// }

// func ValidateEmail(email string) error {
// 	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
// 	if !regex.MatchString(email) {
// 		return fmt.Errorf("invalid email format")
// 	}
// 	return nil
// }

// func IsValidPhoneNumber(phone string) bool {
// 	regex := regexp.MustCompile(`^\+998[0-9]{9}$`)
// 	return regex.MatchString(phone)
// }
