package main

import (
	"time"

	"github.com/go-playground/validator"
)

type DateString struct {
	Value string `validate:"datetime=2006-01-02"`
}

func dateTimeValidator(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	layout := fl.Param()
	_, err := time.Parse(layout, dateStr)
	return err == nil
}

func IsValidDate(dateStr string) bool {
	validate := validator.New()
	validate.RegisterValidation("datetime", dateTimeValidator)

	date := DateString{Value: dateStr}
	err := validate.Struct(date)
	return err == nil
}
