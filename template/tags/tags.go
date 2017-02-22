package tags

type Tags struct {
	Word     string
	Sentence string
	Custom   map[string]string
}

func Default() Tags {
	return Tags{
		Word:     word,
		Sentence: sentence,
		Custom:   make(map[string]string),
	}
}
