package common

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("'%v' is required", fe.Field())
	case "min":
		return fmt.Sprintf("'%v' must be at least %v", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("'%v' must be at most %v", fe.Field(), fe.Param())
	case "eq":
		return fmt.Sprintf("'%v' must be equal to %v", fe.Field(), fe.Param())
	case "ne":
		return fmt.Sprintf("'%v' must not be equal to %v", fe.Field(), fe.Param())
	case "lt":
		return fmt.Sprintf("'%v' must be less than %v", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("'%v' must be less than or equal to %v", fe.Field(), fe.Param())
	case "gt":
		return fmt.Sprintf("'%v' must be greater than %v", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("'%v' must be greater than or equal to %v", fe.Field(), fe.Param())
	case "alpha":
		return fmt.Sprintf("'%v' must contain only alphabetic characters", fe.Field())
	case "alphanum":
		return fmt.Sprintf("'%v' must contain only alphanumeric characters", fe.Field())
	case "numeric":
		return fmt.Sprintf("'%v' must be a numeric value", fe.Field())
	case "email":
		return fmt.Sprintf("'%v' must be a valid email address", fe.Field())
	case "uuid":
		return fmt.Sprintf("'%v' must be a valid UUID", fe.Field())
	case "url":
		return fmt.Sprintf("'%v' must be a valid URL", fe.Field())
	case "datetime":
		return fmt.Sprintf("'%v' must be a valid date/time", fe.Field())
	case "len":
		return fmt.Sprintf("'%v' must be %v characters long", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("Unknown error for field %s", fe.Field())
	}
}

func HandleBindingErr(ctx *gin.Context, err error) {
	if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
		panic(NewBadRequestApiErrorWithCause(
			err,
			"",
			fmt.Sprintf("'%s' must be %s", unmarshalErr.Field, unmarshalErr.Type.Name()),
		))
	}

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errsMsg := make([]string, len(validationErrs))

		for idx, err := range validationErrs {
			errsMsg[idx] = getErrorMsg(err)
		}

		panic(NewBadRequestApiErrorWithCause(err, "", errsMsg))
	}

	panic(NewBadRequestApiError(err, ""))
}
