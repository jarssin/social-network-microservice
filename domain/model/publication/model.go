package publicationModel

type TypePublication string

const (
	Text  TypePublication = "text"
	Image TypePublication = "image"
	Video TypePublication = "video"
)

type Base struct {
	Owner string
	Type  TypePublication
}
