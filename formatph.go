package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func formatPhone(inputPhone string) (string, error) {

	//format 251901040506, 0901040506, +251901040506, 901040506

	if inputPhone == "" {

		return "", errors.New("please enter phone number")

	} else if strings.HasPrefix(inputPhone, "251") ||
		strings.HasPrefix(inputPhone, "+251") ||
		strings.HasPrefix(inputPhone, "09") ||
		strings.HasPrefix(inputPhone, "9") {

		if len(inputPhone) == 9 {
			inputPhone = "251" + inputPhone
			return inputPhone, nil
		} else if len(inputPhone) > 10 {
			return inputPhone[1:], nil
		} else {

			inputPhone = "251" + inputPhone[1:]
			return inputPhone, nil
		}

	} else {
		return "", errors.New("Invalid Phone Number ")
	}

}

func main() {

	phone_number, err := formatPhone("953301736")

	if err != nil {

		log.Fatal(err)
		return
	}

	fmt.Printf("Valid Phone Number :%v\n", phone_number)

}
