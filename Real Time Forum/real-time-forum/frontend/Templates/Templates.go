package Templates

import (
	"fmt"
	"net/http"
	"text/template"
)

// Rendering template
func Rendertemp(w http.ResponseWriter, path string) {
	t, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println(fmt.Printf("%s Error Parsing Template", err))
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(fmt.Printf("%s Error Execute Template", err))
	}
}
