package reveal

import (
	"fmt"
	"regexp"
)

const (
	//DefaultSeparator regexp for horizontal slide boundary token
	DefaultSeparator = "^\r?\n---\r?\n$"
	//DefaultVerticalSeparator regexp for vertical slide boundary token
	DefaultVerticalSeparator = "^\r?\n----\r?\n$"
	//DefaultNoteKeyword regexp for presenter notes paragraphs
	DefaultNoteKeyword = "note:"
)

//NewTokenizer factory function that compiles the passed regexp strings
func NewTokenizer(separator string, verticalSeparator string, noteKeyword string) (rt Tokenizer, err error) {
	bt := &defaultTokenizer{}

	bt.hsSeparator, err = regexp.Compile(fmt.Sprintf("(?m:%s)", defaultIfEmpty(separator, DefaultSeparator)))
	if err != nil {
		return
	}

	bt.vsSeparator, err = regexp.Compile(fmt.Sprintf("(?m:%s)", defaultIfEmpty(verticalSeparator, DefaultVerticalSeparator)))
	if err != nil {
		return
	}

	bt.slideSeparator, err = regexp.Compile(fmt.Sprintf("(?m:%s|%s)", defaultIfEmpty(separator, DefaultSeparator), defaultIfEmpty(verticalSeparator, DefaultVerticalSeparator)))
	if err != nil {
		return
	}

	bt.noteKeyword, err = regexp.Compile(fmt.Sprintf("(?m:%s)", fmt.Sprintf("^%s", defaultIfEmpty(noteKeyword, DefaultNoteKeyword))))
	if err != nil {
		return
	}
	return bt, nil
}

//Section holds markdown content, speaker notes and Section children
type Section struct {
	content  string
	children *[]Section
}

//Tokenizer transforms
type Tokenizer interface {
	Sections(raw []byte) (rs []Section)
}

func defaultIfEmpty(pattern string, fallback string) string {
	if len(pattern) > 0 {
		return pattern
	}

	return fallback
}

//defaultTokenizer Tokenizer implementation
type defaultTokenizer struct {
	slideSeparator *regexp.Regexp
	vsSeparator    *regexp.Regexp
	hsSeparator    *regexp.Regexp
	noteKeyword    *regexp.Regexp
}

func (t *defaultTokenizer) Sections(raw []byte) (rs []Section) {
	rs = []Section{}
	i := t.slideSeparator.FindIndex(raw)
	wasHorizontal := true
	var s Section
	var isHorizontal bool

	for len(i) > 0 {
		//TODO Extract presenter notes from content in advance
		s = Section{content: string(raw[:i[0]-1])}

		isHorizontal = t.isHorizontalSeparator(raw[i[0]:i[1]])

		if isHorizontal && wasHorizontal {
			rs = append(rs, s)
		} else {
			if wasHorizontal {
				tmpHead := Section{content: "This shouldn't be shown in the slides", children: &[]Section{s}}
				rs = append(rs, tmpHead)
			} else {
				tmpHead := rs[len(rs)-1]
				*tmpHead.children = append(*tmpHead.children, s)
			}
		}

		// Shift byte array after the separator
		raw = raw[i[1]+1:]
		i = t.slideSeparator.FindIndex(raw)
		wasHorizontal = isHorizontal
	}

	s = Section{content: string(raw)}

	if wasHorizontal {
		rs = append(rs, s)
	} else {
		tmpHead := rs[len(rs)-1]
		*tmpHead.children = append(*tmpHead.children, s)
	}

	return
}

func (t *defaultTokenizer) isHorizontalSeparator(input []byte) bool {
	return t.hsSeparator.Match(input)
}

func (t *defaultTokenizer) isVerticalSeparator(input []byte) bool {
	return t.vsSeparator.Match(input)
}
