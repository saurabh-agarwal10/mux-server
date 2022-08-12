package dto

type UserDetailsReq struct {
	UserID string `json:"user_id"`
}

type UserDetailsRes struct {
	Code     int    `json:"http_code"`
	Message  string `json:"message"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Services string `json:"services"`
}
