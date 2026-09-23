package tasscommon

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	playground "github.com/go-playground/validator/v10"
)

const datetimeRegex string = `[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]{3}`
const dateRegex string = `[0-9]{4}-[0-9]{2}-[0-9]{2}`

const dateFormat string = `2006-01-02`
const dateTimeFormat string = `2006-01-02T03:04:05.000`

var validator = newStructValidator()

func Validate(value any) error {
	err := validator.Struct(value)
	if err == nil {
		return nil
	}

	validationErrors, ok := errors.AsType[playground.ValidationErrors](err)
	if !ok || len(validationErrors) == 0 {
		return err
	}
	return fieldError{validationErrors[0]}
}

func newStructValidator() *playground.Validate {
	validate := playground.New(playground.WithRequiredStructEnabled())
	if err := validate.RegisterValidation("datebeforetoday", ValidateDateBeforeToday); err != nil {
		panic(err)
	}
	return validate
}

func ValidateDateBeforeToday(field playground.FieldLevel) bool {
	now := time.Now()
	date, err := time.Parse(field.Param(), field.Field().String())
	if err != nil {
		return false
	}
	if date.Year() != now.Year() ||
		date.Month() != now.Month() ||
		date.Day() != now.Day() {
		return false
	}
	return true
}

func ValidateDateAfterDate(field playground.FieldLevel) bool {
	fieldDate, err := parseDate(field.Field().String())
	if err != nil {
		return false
	}
	validationDate, err := parseDate(field.Field().String())
	if err != nil {
		return false
	}
	if fieldDate.Before(validationDate) {
		return false
	}
	return true
}

func parseDate(date string) (time.Time, error) {
	match, err := regexp.MatchString(datetimeRegex, date)
	if err != nil {
		return time.Time{}, err
	}
	if match {
		parsedDate, err := time.Parse(dateTimeFormat, date)
		if err != nil {
			return time.Time{}, err
		}
		return parsedDate, nil
	}

	match, err = regexp.MatchString(dateRegex, date)
	if err != nil {
		return time.Time{}, err
	}
	if match {
		parsedDate, err := time.Parse(dateFormat, date)
		if err != nil {
			return time.Time{}, err
		}
		return parsedDate, nil
	}
	return time.Time{}, fmt.Errorf("date %s does not match any configured formats", date)
}

type fieldError struct {
	playground.FieldError
}

func (err fieldError) Error() string {
	field := err.Field()
	switch err.Tag() {
	case "max":
		return fmt.Sprintf("%s must be at most %s", field, err.Param())
	case "datebeforetoday":
		return field + " must be a date and must be before today"
	case "datetime":
		return fmt.Sprintf("%s must be a date in the format %s", field, err.Param())
	case "nefield":
		return fmt.Sprintf("%s cannot be the same as %s", field, err.Param())
	case "dateafterdate":
		return fmt.Sprintf("%s must be after %s", field, err.Param())
	default:
		return field + " is invalid"
	}
}
