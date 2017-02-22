package template

import (
	"io"

	"strings"

	"github.com/panjiesw/apimocker/mockers"
	"github.com/panjiesw/apimocker/template/tags"
	ft "github.com/valyala/fasttemplate"
)

type Template interface {
	Execute(template string) string
}

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
	}
	return w.Write([]byte("a"))
}

func (t *tpl) Execute(template string) string {
	return ft.ExecuteFuncString(template, t.startTag, t.endTag, t.tagFunc)
}
