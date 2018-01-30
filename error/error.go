package error

type JsonError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
