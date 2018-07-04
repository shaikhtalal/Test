package shared

import "time"

type UserSignup struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Messages struct {
	Message string `json:"message"`
}

type GeoIP struct {
	// The right side is the name of the JSON variable
	IP          string    `json:"ip"`
	Type        string    `json:"type"`
	CountryCode string    `json:"country_code"`
	CountryName string    `json:"country_name"`
	RegionCode  string    `json:"region_code"`
	RegionName  string    `json:"region_name"`
	City        string    `json:"city"`
	Zipcode     string    `json:"zip"`
	Lat         float32   `json:"latitude"`
	Lon         float32   `json:"longitude"`
	Time        time.Time `json:"time"`
}
