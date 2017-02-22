package mockers

import "github.com/panjiesw/apimocker/mockers/lorem"

// GenStrWord generate random word with default length range (5, 10)
func GenStrWord() string {
	return lorem.Word(5, 10)
}

// GenStrWordRange generate randowm word in range of min-max
func GenStrWordRange(min, max int) string {
	return lorem.Word(min, max)
}

// GenStrSentence generate random sentence with word count in default range (5, 10)
func GenStrSentence() string {
	return lorem.Sentence(5, 10)
}

// GenStrSentenceRange generate random sentence with word count in range of min-max
func GenStrSentenceRange(min, max int) string {
	return lorem.Sentence(min, max)
}

// GenStrParagraph generate random paragraph with sentence count in default range (5, 10)
func GenStrParagraph() string {
	return lorem.Paragraph(5, 10)
}

// GenStrParagraphRange generate random paragraph with sentence count in range of min-max
func GenStrParagraphRange(min, max int) string {
	return lorem.Paragraph(min, max)
}
