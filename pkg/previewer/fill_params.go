package previewer

import "context"

type FillParams struct {
	ctx     context.Context
	width   int
	height  int
	url     string
	headers map[string][]string
}

func NewFillParams(ctx context.Context, width int, height int, url string, headers map[string][]string) *FillParams {
	return &FillParams{ctx: ctx, width: width, height: height, url: url, headers: headers}
}
