package previewer

// Service @TODO fixme.
type Service interface {
	Fill(width, height int, imgUrl string) (img []byte, err error)
}
