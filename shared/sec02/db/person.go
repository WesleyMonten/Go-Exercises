package db

type person struct {
	ID        uint     `json:"id,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Gender    string   `json:"gender,omitempty"`
	Income    Currency `json:"income,omitempty"`
	Zip       string   `json:"zip,omitempty"`
}
