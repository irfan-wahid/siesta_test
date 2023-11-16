package models

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	NISN   string `json:"nisn"`
	Class  string `json:"class"`
	School string `json:"school"`
}

type Data struct {
	Classes         int       `json:"classes"`
	Students        []Student `json:"students"`
	Sicks           []Student `json:"sicks"`
	PermittedLeaves []Student `json:"permitted_leaves"`
}
