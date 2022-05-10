package previewer

type FillResponse struct {
	Img     []byte
	Headers map[string][]string
}

func NewFillResponse(img []byte, headers map[string][]string) *FillResponse {
	return &FillResponse{Img: img, Headers: headers}
}
