package dto

type PeopleAgeClientDTO struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type PeopleGenderClientDTO struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      *string `json:"gender"`
	Probability float64 `json:"probability"`
}

type PeopleNationalizeClientDTO struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryId   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}
