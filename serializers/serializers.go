package serializers

// Resp : struct
type Resp struct {
	Result interface{} `json:"Result"`
	Error  interface{} `json:"Error"`
}
