package template

import (
	"io"

	"strings"

	"fmt"

	"github.com/panjiesw/apimocker/mockers"
	"github.com/panjiesw/apimocker/template/tags"
	ft "github.com/valyala/fasttemplate"
)

// Template is the interface that wraps the Execute method
// to replace tags with their definition.
type Template interface {
	Execute(template string) string
}

// New create a default Template implementation
func New(start string, end string, tags tags.Tags) Template {
	return &tpl{
		startTag: start,
		endTag:   end,
		tags:     tags,
	}
}

type tpl struct {
	startTag string
	endTag   string
	tags     tags.Tags
}

func (t *tpl) tagFunc(w io.Writer, tag string) (int, error) {
	tag = strings.TrimSpace(tag)
	if strings.HasPrefix(tag, t.tags.Word) {
		if tag == t.tags.Word {
			return w.Write([]byte(mockers.GenStrWord()))
		}
		args, ok := validateFn2IntArgs(t.tags.Word, tag)
		if ok {
			ww := mockers.GenStrWordRange(args[0], args[1])
			return w.Write([]byte(ww))
		}
	} else if strings.HasPrefix(tag, t.tags.Sentence) {
		if tag == t.tags.Sentence {
			return w.Write([]byte(mockers.GenStrSentence()))
		}
		args, ok := validateFn2IntArgs(t.tags.Sentence, tag)
		if ok {
			return w.Write([]byte(mockers.GenStrSentenceRange(args[0], args[1])))
		}
	} else if strings.HasPrefix(tag, t.tags.Paragraph) {
		if tag == t.tags.Paragraph {
			return w.Write([]byte(mockers.GenStrParagraph()))
		}
		args, ok := validateFn2IntArgs(t.tags.Paragraph, tag)
		if ok {
			return w.Write([]byte(mockers.GenStrParagraphRange(args[0], args[1])))
		}
	}
	return w.Write([]byte(fmt.Sprintf("<invalid tag: %s>", tag)))
}

// Execute substitute tags with their definition from the input template.
func (t *tpl) Execute(template string) string {
	return ft.ExecuteFuncString(template, t.startTag, t.endTag, t.tagFunc)
}
