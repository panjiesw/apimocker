package template_test

import (
	"testing"

	"strings"

	"fmt"

	"github.com/panjiesw/apimocker/template"
	"github.com/panjiesw/apimocker/template/tags"
)

type etests struct {
	name     string
	template string
	er       bool
	st       bool
	pg       bool
	rg       []int
}

func Test_Tempalate_Execute(t *testing.T) {
	tTags := tags.Default()
	tTags.Word = "kata"
	type newargs struct {
		start string
		end   string
		tags  tags.Tags
	}
	newT := []struct {
		name  string
		args  newargs
		tests []etests
	}{
		{
			name:  "default",
			args:  newargs{start: "{{", end: "}}", tags: tags.Default()},
			tests: generateTests("{{", "}}", tags.Default()),
		},
		{
			name:  "custom start end",
			args:  newargs{start: "<<", end: ">>", tags: tags.Default()},
			tests: generateTests("<<", ">>", tags.Default()),
		},
		{
			name:  "custom tags",
			args:  newargs{start: "<<", end: ">>", tags: tTags},
			tests: generateTests("<<", ">>", tTags),
		},
	}
	for _, nt := range newT {
		t.Run(nt.name, func(t *testing.T) {
			tn := template.New(nt.args.start, nt.args.end, nt.args.tags)
			for _, tt := range nt.tests {
				t.Run(tt.name, func(t *testing.T) {
					got := tn.Execute(tt.template)
					lg := len(got)
					doLog := true
					if strings.Contains(got, "invalid tag") {
						t.Errorf("Template.Execute() = %v, want valid", got)
					}

					if tt.st {
						lg = len(strings.Split(got, " "))
						doLog = false
					} else if tt.pg {
						lg = len(strings.Split(got, "."))
						doLog = false
					}

					if tt.er && !(lg >= tt.rg[0] && lg <= tt.rg[1]) {
						t.Errorf("len(Template.Execute()) = %v, %v<=want<=%v", lg, tt.rg[0], tt.rg[1])
					}

					if doLog {
						t.Log(got)
					}
				})
			}
		})
	}
}

func generateTests(start, end string, tags tags.Tags) []etests {
	return []etests{
		etests{
			name:     "valid word",
			template: fmt.Sprintf("%s%s%s", start, tags.Word, end),
			er:       true,
			rg:       []int{5, 10},
		},
		etests{
			name:     "valid range word",
			template: fmt.Sprintf("%s%s(2, 8)%s", start, tags.Word, end),
			er:       true,
			rg:       []int{2, 8},
		},
		etests{
			name:     "valid big range word",
			template: fmt.Sprintf("%s%s(2, 100000)%s", start, tags.Word, end),
			er:       true,
			rg:       []int{2, 100000},
		},
		etests{
			name:     "valid sentence",
			template: fmt.Sprintf("%s%s%s", start, tags.Sentence, end),
			er:       true,
			rg:       []int{5, 10},
			st:       true,
		},
		etests{
			name:     "valid range sentence",
			template: fmt.Sprintf("%s%s(2, 8)%s", start, tags.Sentence, end),
			er:       true,
			st:       true,
			rg:       []int{2, 8},
		},
		etests{
			name:     "valid big range sentence",
			template: fmt.Sprintf("%s%s(2, 100000)%s", start, tags.Sentence, end),
			er:       true,
			st:       true,
			rg:       []int{2, 100000},
		},
		etests{
			name:     "valid paragraph",
			template: fmt.Sprintf("%s%s%s", start, tags.Paragraph, end),
			er:       true,
			rg:       []int{5, 10},
			pg:       true,
		},
		etests{
			name:     "valid range paragraph",
			template: fmt.Sprintf("%s%s(2, 8)%s", start, tags.Paragraph, end),
			er:       true,
			pg:       true,
			rg:       []int{2, 8},
		},
		etests{
			name:     "valid big range paragraph",
			template: fmt.Sprintf("%s%s(2, 100000)%s", start, tags.Paragraph, end),
			er:       true,
			pg:       true,
			rg:       []int{2, 100000},
		},
	}
}
