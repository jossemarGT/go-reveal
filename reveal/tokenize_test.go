package reveal

import (
	"io/ioutil"
	"testing"
)

func TestSections(t *testing.T) {
	tk, err := NewTokenizer("", "", "")

	if err != nil {
		t.Fatal("Failed getting new tokenizer")
	}

	var sectionTests = []struct {
		goldenfilePath string
		childrenCount  []int
	}{
		{"testdata/00_basic.md", []int{0, 0, 0}},
		{"testdata/01_full_no_modifiers.md", []int{0, 3, 0, 0}},
		{"testdata/02_short_vlast.md", []int{0, 3}},
		{"testdata/03_links.md", []int{0, 0, 2, 0}},
	}

	for _, tt := range sectionTests {
		b, err2 := ioutil.ReadFile(tt.goldenfilePath)

		if err2 != nil {
			t.Fatal("Failed reading golden file")
		}

		s := tk.Sections(b)

		if len(s) != len(tt.childrenCount) {
			t.Fatalf("Section count missmatch, expected %d, got %d", len(tt.childrenCount), len(s))
		}

		for i, ss := range s {
			if ss.Children == nil && tt.childrenCount[i] > 0 {
				t.Errorf("Content node #%d has no children, expected %d", i, tt.childrenCount[i])
			}

			if ss.Children != nil && (len(*ss.Children) != tt.childrenCount[i]) {
				t.Errorf("Content node #%d children count missmatch, expected %d, got %d", i, tt.childrenCount[i], len(*ss.Children))
			}
		}
	}
}
