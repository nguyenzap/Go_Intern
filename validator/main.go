package main

import (
	"log"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `validate:"required"`
	Email string `validate:"required,email"`
	Age int `validate:"gte=18"`
}

func main() {
	u := User{"Zap", "wrongmail", 17}

	validate := validator.New()

	err := validate.Struct(u)

	if err != nil {
		log.Println("validation failed")
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
		}
	} else {
		log.Println("Validation Pass")
	}
}