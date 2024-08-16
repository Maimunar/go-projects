package api

// Success Responses
type TwoNumbersBody struct {
	Number1 int `json:"number_1,omitempty"`
	Number2 int `json:"number_2,omitempty"`
}

type NumberArraysBody struct {
	Numbers []int `json:"numbers,omitempty"`
}

type SuccessNumberResponse struct {
	Result int `json:"result,omitempty"`
}
