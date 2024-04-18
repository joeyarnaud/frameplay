package validators

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

// validateCredentials checks if the provided email and password are valid.
func ValidateCredentials(email, password string) error {
    // Email validation
    if err := validateEmail(email); err != nil {
        return err
    }

    // Password validation
    if err := validatePassword(password); err != nil {
        return err
    }

    return nil
}

// validateEmail checks if the provided email is in a valid format.
func validateEmail(email string) error {
    // Regular expression for validating an email
    regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
    if !regex.MatchString(email) {
        return fmt.Errorf("invalid email format")
    }
    return nil
}


// validatePassword checks if the password meets the defined criteria.
func validatePassword(password string) error {
    var (
        hasMinLen  = false
        hasUpper   = false
        hasLower   = false
        hasNumber  = false
        hasSpecial = false
    )

    // Check for minimum length
    if len(password) >= 8 {
        hasMinLen = true
    }

    // Check for other requirements
    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsDigit(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }

    // Check if all conditions are met
    if !hasMinLen {
        return errors.New("password must be at least 8 characters long")
    }
    if !hasUpper {
        return errors.New("password must include at least one uppercase letter")
    }
    if !hasLower {
        return errors.New("password must include at least one lowercase letter")
    }
    if !hasNumber {
        return errors.New("password must include at least one digit")
    }
    if !hasSpecial {
        return errors.New("password must include at least one special character")
    }

    return nil
}



