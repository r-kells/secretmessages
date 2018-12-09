package views

import (
	"html/template"
	"net/http"
	"path"
	"path/filepath"
)

const (
	templateDir string = "views"
	templateExt string = ".gohtml"
	layoutDir   string = "views/layouts/"
	layoutsGlob string = layoutDir + "*" + templateExt
)

// View is a ready to be used html/template
type View struct {
	Template *template.Template
}

// Render manages the logic for rendering views so the controller doesn't have to.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.Execute(w, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// NewView helps create a new `View`. It removes the boilerplate of
// appending common template files to nested templates.
func NewView(files ...string) *View {
	buildFullPath(files)

	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{Template: t}
}

func layoutFiles() []string {
	files, err := filepath.Glob(layoutsGlob)
	if err != nil {
		panic(err)
	}
	return files
}

// buildFullPath is a helper for generating relative file paths
// from an assumed directory structure of views.
//
// Ex. input of "home" would return "views/home.gohtml".
func buildFullPath(files []string) {
	for i, f := range files {
		files[i] = path.Join(templateDir, f) + templateExt
	}
}
