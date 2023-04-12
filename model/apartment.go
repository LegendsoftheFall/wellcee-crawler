package model

type ApartmentDetail struct {
	ID       string `json:"id"`
	Url      string `json:"url"`
	Floor    string `json:"floor"`
	Type     string `json:"type"`
	Deposit  string `json:"deposit"`
	Room     string `json:"room"`
	Area     string `json:"area"`
	Subway   string `json:"subway"`
	Status   string `json:"status"`
	Price    string `json:"price"`
	Intro    string `json:"intro"`
	Location string `json:"location"`
	//Tag      []string
	//Facility []string
	//Image    []string
	//Video    []string
}
