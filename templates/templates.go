package templates

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed folder/single_file.txt
//go:embed folder/*.hash
//go:embed templateFiles/*.go.html
var folder embed.FS

func GetTemplates() {
	//println("Something that is there")
	embedTesting()
}

// based on - https://gobyexample.com/embed-directive
// Just making sure that we can read contents out of files
func embedTesting() {
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("templateFiles/userProfile.go.html")
	print(string(content2))
}

// trying to return a template file tat can be executed
func GetUsableTemplate(templatePath string) *template.Template {

	tmpl := template.Must(template.ParseFS(folder, templatePath))
	fmt.Println(tmpl)

	return tmpl
}
