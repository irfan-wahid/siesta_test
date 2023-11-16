package controllers

import (
	"fmt"
	"q4/models"
	"sync"
)

func ChangeGender(wg *sync.WaitGroup, user *models.User) {
	defer wg.Done()

	if user.Gender == "F" {
		user.Gender = "Female"
	} else if user.Gender == "M" {
		user.Gender = "Male"
	}

	fmt.Printf("%s's gender is %s\n", user.Name, user.Gender)
}
