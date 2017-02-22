package tags

// Tags provide structured tag definition
type Tags struct {
	Word      string
	Sentence  string
	Paragraph string
	Custom    map[string]string
}

// Default create a default Tags structure
func Default() Tags {
	return Tags{
		Word:      word,
		Sentence:  sentence,
		Paragraph: paragraph,
		Custom:    make(map[string]string),
	}
}
