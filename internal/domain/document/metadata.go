package document

type Metadata struct {
	source    string
	mediaType string
}

func NewMetadata(source, mediaType string) Metadata {
	return Metadata{
		source:    source,
		mediaType: mediaType,
	}
}

func (m Metadata) Source() string {
	return m.source
}

func (m Metadata) MediaType() string {
	return m.mediaType
}
