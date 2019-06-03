package serializers

// UserLoginReq : struct
type UserLoginReq struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// UserRegisterReq : struct
type UserRegisterReq struct {
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}

// UserSubscribeReq : struct
type UserSubscribeReq struct {
	Email   string `json:"Email"`
	Name    string `json:"Name"`
	Company string `json:"Company"`
}
