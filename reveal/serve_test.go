package reveal

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestListingHandler(t *testing.T) {
	pwd, _ := os.Getwd()
	listingTemplate := "directory.gotemplate"

	specs := []struct {
		basePath       string
		templatePath   string
		requestMethod  string
		requestPath    string
		statusResponse int
	}{
		// List a directory that exists
		{pwd, listingTemplate, "GET", "/testdata", http.StatusOK},
		// Trying to reach ancestors redirects to /
		{pwd, listingTemplate, "GET", "/../..", http.StatusOK},
		{pwd, listingTemplate, "GET", "/..", http.StatusOK},
		// 404 when child directory doesn't exist
		{pwd, listingTemplate, "GET", "/notfound", http.StatusNotFound},
	}

	for _, s := range specs {
		req := httptest.NewRequest(s.requestMethod, s.requestPath, nil)
		wr := httptest.NewRecorder()

		l := ListingHandler{BasePath: s.basePath, TemplatePath: s.templatePath}
		l.ServeHTTP(wr, req, nil)

		res := wr.Result()

		if res.StatusCode != s.statusResponse {
			t.Log(s.basePath)
			body, _ := ioutil.ReadAll(res.Body)
			t.Logf("%s", body)
			t.Fatalf("Unexpected response status, expected %d, got %d", s.statusResponse, res.StatusCode)
		}
	}

}

type fakeFileInfo struct {
	filename string
	asDir    bool
}

func (f fakeFileInfo) Name() string {
	return f.filename
}

func (f fakeFileInfo) Size() int64 {
	return int64(0)
}

func (f fakeFileInfo) Mode() os.FileMode {
	return os.FileMode(755)
}

func (f fakeFileInfo) ModTime() time.Time {
	return time.Now()
}

func (f fakeFileInfo) IsDir() bool {
	return f.asDir
}

func (f fakeFileInfo) Sys() interface{} {
	return nil
}

func TestIsValidFile(t *testing.T) {
	var FileSpecs = []struct {
		fileSpec os.FileInfo
		expected bool
	}{
		{fakeFileInfo{filename: "..", asDir: true}, false},
		{fakeFileInfo{filename: ".hidden", asDir: false}, false},
		{fakeFileInfo{filename: ".hidden.md", asDir: false}, false},
		{fakeFileInfo{filename: ".hidden.markdown", asDir: true}, false},
		{fakeFileInfo{filename: ".hidden", asDir: true}, false},
		{fakeFileInfo{filename: "a-random-file.txt", asDir: false}, false},
		{fakeFileInfo{filename: "a-random-name-with-md", asDir: false}, false},
		{fakeFileInfo{filename: "a-random-name-with-markdown.txt", asDir: false}, false},
		{fakeFileInfo{filename: "a-random-file.md", asDir: false}, true},
		{fakeFileInfo{filename: "a-random-file.markdown", asDir: false}, true},
		{fakeFileInfo{filename: "a_directory", asDir: true}, true},
	}

	for _, tt := range FileSpecs {
		if tt.expected != isValidFile(tt.fileSpec) {
			t.Fatalf("Expect: %t for File %s as directory %t, got: %t", tt.expected, tt.fileSpec.Name(), tt.fileSpec.IsDir(), !tt.expected)
		}

	}
}
