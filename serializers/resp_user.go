package serializers

// UserLoginResp : struct
type UserLoginResp struct {
	Token   string `json:"Token"`
	Expired string `json:"Expired"`
}
