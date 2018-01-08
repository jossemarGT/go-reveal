package reveal

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/negroni"
)

//MarkdownExt is a markdown extension regexp helper
var MarkdownExt *regexp.Regexp

func init() {
	MarkdownExt = regexp.MustCompile("\\.(?i:md|markdown)$")
}

//NewRevealHandler ServerMux decorator which adds an SliderHander, a ListingHandler and some third party "middlewares"
func NewRevealHandler(m *http.ServeMux, t Tokenizer, basePath string) http.Handler {
	s := SlideHandler{
		SlidePreprocessor: t,
		BasePath:          basePath,
		TemplatePath:      path.Join("reveal", "slide.gotemplate"),
	}

	l := ListingHandler{
		BasePath:     basePath,
		TemplatePath: path.Join("reveal", "directory.gotemplate"),
	}

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		s,
		negroni.NewStatic(http.Dir(basePath)),
		l,
	)
	n.UseHandler(m)

	return n
}

type deck struct {
	Title    string
	Secret   string
	Sections []Section
}

//SlideHandler HTTP Middleware that handles Slide rendering
type SlideHandler struct {
	Title             string
	BasePath          string
	Secret            string
	TemplatePath      string
	SlidePreprocessor Tokenizer
}

func (sh SlideHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var err error
	rf := r.URL.EscapedPath()

	if !MarkdownExt.MatchString(rf) {
		next(w, r)
		return
	}

	if _, err := os.Stat(filepath.Join(sh.BasePath, rf)); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 page not found")
		return
	}

	qs := r.URL.Query().Get("secret")

	if qs != "" && sh.Secret != qs {
		qs = ""
	}

	f, _ := ioutil.ReadFile(filepath.Join(sh.BasePath, rf))
	ss := sh.SlidePreprocessor.Sections(f)

	d := &deck{
		Title:    sh.Title,
		Secret:   qs,
		Sections: ss,
	}

	t, err := template.New("slide.gotemplate").ParseFiles(sh.TemplatePath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server error:<br/>%v", err)
	}

	w.WriteHeader(http.StatusOK)
	t.Execute(w, d)
}

//ListingHandler HTTP Middleware that handles directory listing for markdown files
type ListingHandler struct {
	BasePath     string
	TemplatePath string
}

type listing struct {
	Path  string
	Files []filetuple
}

type filetuple struct {
	Name    string
	Relpath string
}

func (lh ListingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rf := filepath.Join(lh.BasePath, path.Clean(r.URL.EscapedPath()))

	if _, err := os.Stat(rf); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 page not found")
		return
	}

	ff, err := ioutil.ReadDir(rf)
	if err != nil {
		fmt.Fprintf(w, "Internal Server error:<br/>%v", err)
		return
	}

	d := listing{Path: r.URL.EscapedPath(), Files: []filetuple{}}

	for _, f := range ff {
		if isValidFile(f) {
			rel, _ := filepath.Rel(lh.BasePath, filepath.Join(rf, f.Name()))

			ft := filetuple{
				Name:    f.Name(),
				Relpath: rel,
			}

			d.Files = append(d.Files, ft)
		}
	}

	t, err := template.New("directory.gotemplate").ParseFiles(lh.TemplatePath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server error:<br/>%v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	t.Execute(w, d)
}

func isValidFile(f os.FileInfo) bool {
	return !strings.HasPrefix(f.Name(), ".") && (f.IsDir() || MarkdownExt.MatchString(f.Name()))
}
