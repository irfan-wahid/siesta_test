package main

import (
	"q4/controllers"
	"q4/models"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	users := []models.User{
		{Name: "John", Gender: "F"},
		{Name: "Jack", Gender: "F"},
		{Name: "May", Gender: "M"},
		{Name: "Sonia", Gender: "M"},
	}

	for i := range users {
		wg.Add(1)
		go controllers.ChangeGender(&wg, &users[i])
	}

	wg.Wait()
}
