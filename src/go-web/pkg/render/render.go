package render

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("/templates"),
	jet.InDevelopmentMode(),
)

func RenderTemplate(w http.ResponseWriter, tmp string, data jet.VarMap) {
	view, err := views.GetTemplate(tmp)
	if err != nil {
		log.Fatal(err)
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Fatal(err)
	}
}
