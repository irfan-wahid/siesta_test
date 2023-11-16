package main

import (
	"encoding/json"
	"fmt"
	"q5/models"
)

func main() {
	// Data JSON
	dataJSON := `{
		"classes": 14,
		"students": [
			{
				"id": 12,
				"name": "Fikri Naufal",
				"nisn": "43567200883",
				"class": "VII",
				"school": "MTs Ali Maksum"
			},
			{
				"id": 13,
				"name": "Moh Rasya",
				"nisn": "4356720042",
				"class": "X",
				"school": "MA Ali Maksum"
			},
			{
				"id": 30,
				"name": "Naufal Kurniawan",
				"nisn": "63567200868",
				"class": "XI",
				"school": "MA Ali Maksum"
			},
			{
				"id": 39,
				"name": "Aminuddin Abdallah",
				"nisn": "43567200827",
				"class": "XII",
				"school": "MA Ali Maksum"
			}
		],
		"sicks": [
			{
				"id": 39,
				"name": "Aminuddin Abdallah",
				"nisn": "43567200827",
				"class": "XII",
				"school": "MA Ali Maksum"
			},
			{
				"id": 39,
				"name": "Aminuddin Abdallah",
				"nisn": "43567200827",
				"class": "XII",
				"school": "MA Ali Maksum"
			},
			{
				"id": 13,
				"name": "Moh Rasya",
				"nisn": "4356720042",
				"class": "X",
				"school": "MA Ali Maksum"
			}
		],
		"permitted_leaves": [
			{
				"id": 12,
				"name": "Fikri Naufal",
				"nisn": "43567200883",
				"class": "VII",
				"school": "MTs Ali Maksum"
			}
		]
	}`

	// Unmarshal data JSON
	var data models.Data
	err := json.Unmarshal([]byte(dataJSON), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

}
